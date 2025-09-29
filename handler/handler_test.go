package handler

import (
	"context"
	"database/sql"
	"encoding/gob"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/entity"
	goauth "google.golang.org/api/oauth2/v2"
)

// MockProducerService は ProducerService のモック実装
type MockProducerService struct {
	findProducerFunc   func(ctx context.Context, identityId string) (entity.Producer, error)
	registProducerFunc func(ctx context.Context, identityId string) error
}

func (m *MockProducerService) FindProducer(ctx context.Context, identityId string) (entity.Producer, error) {
	if m.findProducerFunc != nil {
		return m.findProducerFunc(ctx, identityId)
	}
	return entity.Producer{}, nil
}

func (m *MockProducerService) RegistProducer(ctx context.Context, identityId string) error {
	if m.registProducerFunc != nil {
		return m.registProducerFunc(ctx, identityId)
	}
	return nil
}

// setupTestRouter はテスト用のルーターをセットアップ
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// セッション設定
	gob.Register(&UserSession{})
	store := cookie.NewStore([]byte("test-secret"))
	router.Use(sessions.Sessions("test-session", store))

	return router
}

func TestAuthHandler_NewUser(t *testing.T) {
	router := setupTestRouter()

	// モックサービス設定（新規ユーザー）
	var findCallCount int
	mockService := &MockProducerService{
		findProducerFunc: func(ctx context.Context, identityId string) (entity.Producer, error) {
			findCallCount++
			if identityId == "new-user-123" {
				if findCallCount == 1 {
					return entity.Producer{}, sql.ErrNoRows // 最初の呼び出しではユーザーが見つからない
				}
				// 2回目の呼び出しでは登録後のユーザーを返す
				return entity.Producer{ID: 1, IdentityId: identityId}, nil
			}
			return entity.Producer{}, sql.ErrNoRows
		},
		registProducerFunc: func(ctx context.Context, identityId string) error {
			if identityId == "new-user-123" {
				return nil // 登録成功
			}
			return fmt.Errorf("unexpected identity id: %s", identityId)
		},
	}

	loginHandler := &LoginProducer{
		ProducerService: mockService,
	}

	router.GET("/auth/", func(c *gin.Context) {
		// Google認証情報をモック
		c.Set("user", goauth.Userinfo{
			Id:    "new-user-123",
			Email: "test@example.com",
		})
		loginHandler.AuthHandler(c)
	})

	// テスト実行
	req, _ := http.NewRequest("GET", "/auth/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 結果検証
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

	// リダイレクト先の確認
	location := w.Header().Get("Location")
	if location != "/auth/members" {
		t.Errorf("Expected redirect to /auth/members, got %s", location)
	}
}

func TestAuthHandler_ExistingUser(t *testing.T) {
	router := setupTestRouter()

	// モックサービス設定（既存ユーザー）
	mockService := &MockProducerService{
		findProducerFunc: func(ctx context.Context, identityId string) (entity.Producer, error) {
			if identityId == "existing-user-456" {
				return entity.Producer{ID: 2, IdentityId: identityId}, nil
			}
			return entity.Producer{}, sql.ErrNoRows
		},
	}

	loginHandler := &LoginProducer{
		ProducerService: mockService,
	}

	router.GET("/auth/", func(c *gin.Context) {
		// Google認証情報をモック
		c.Set("user", goauth.Userinfo{
			Id:    "existing-user-456",
			Email: "existing@example.com",
		})
		loginHandler.AuthHandler(c)
	})

	// テスト実行
	req, _ := http.NewRequest("GET", "/auth/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 結果検証
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

	location := w.Header().Get("Location")
	if location != "/auth/members" {
		t.Errorf("Expected redirect to /auth/members, got %s", location)
	}
}

func TestAuthHandler_RegistrationError(t *testing.T) {
	router := setupTestRouter()

	// モックサービス設定（登録エラー）
	mockService := &MockProducerService{
		findProducerFunc: func(ctx context.Context, identityId string) (entity.Producer, error) {
			return entity.Producer{}, sql.ErrNoRows // ユーザーが見つからない
		},
		registProducerFunc: func(ctx context.Context, identityId string) error {
			return fmt.Errorf("registration failed") // 登録エラー
		},
	}

	loginHandler := &LoginProducer{
		ProducerService: mockService,
	}

	router.GET("/auth/", func(c *gin.Context) {
		c.Set("user", goauth.Userinfo{
			Id:    "error-user-789",
			Email: "error@example.com",
		})
		loginHandler.AuthHandler(c)
	})

	// テスト実行
	req, _ := http.NewRequest("GET", "/auth/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 結果検証（エラーの場合はInternal Server Error）
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expected := `{"status":"ok"}`
	if w.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, w.Body.String())
	}
}

func TestBasicRouting(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := []struct {
		name         string
		method       string
		path         string
		expectedCode int
	}{
		{
			name:         "NotFoundRoute",
			method:       "GET",
			path:         "/nonexistent",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "GetToNonExistentPath",
			method:       "GET",
			path:         "/api/test",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			router := gin.New()
			router.GET("/health", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
			})

			req, _ := http.NewRequest(tc.method, tc.path, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tc.expectedCode {
				t.Errorf("Expected status code %d, got %d", tc.expectedCode, w.Code)
			}
		})
	}
}

func TestGetUserSession_AdminDebug(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		adminDebug     string
		originalAdmin  bool
		expectedAdmin  bool
	}{
		{
			name:          "ADMIN_DEBUG=true sets IsAdmin to true",
			adminDebug:    "true",
			originalAdmin: false,
			expectedAdmin: true,
		},
		{
			name:          "ADMIN_DEBUG=true keeps IsAdmin true",
			adminDebug:    "true",
			originalAdmin: true,
			expectedAdmin: true,
		},
		{
			name:          "ADMIN_DEBUG=false preserves original IsAdmin false",
			adminDebug:    "false",
			originalAdmin: false,
			expectedAdmin: false,
		},
		{
			name:          "ADMIN_DEBUG=false preserves original IsAdmin true",
			adminDebug:    "false",
			originalAdmin: true,
			expectedAdmin: true,
		},
		{
			name:          "No ADMIN_DEBUG preserves original IsAdmin",
			adminDebug:    "",
			originalAdmin: false,
			expectedAdmin: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 環境変数を設定
			originalValue := os.Getenv("ADMIN_DEBUG")
			defer func() {
				if originalValue == "" {
					os.Unsetenv("ADMIN_DEBUG")
				} else {
					os.Setenv("ADMIN_DEBUG", originalValue)
				}
			}()

			if tt.adminDebug != "" {
				os.Setenv("ADMIN_DEBUG", tt.adminDebug)
			} else {
				os.Unsetenv("ADMIN_DEBUG")
			}

			router := setupTestRouter()

			// セッションにユーザー情報を設定
			router.GET("/test", func(c *gin.Context) {
				session := sessions.Default(c)
				userSession := &UserSession{
					ProducerId: 1,
					IdentityId: "test-user",
					EMail:      "test@example.com",
					LoggedIn:   true,
					IsAdmin:    tt.originalAdmin,
				}
				session.Set("uniar_session", userSession)
				session.Save()

				// getUserSessionを呼び出してテスト
				us, err := getUserSession(c)
				if err != nil {
					t.Errorf("getUserSession returned error: %v", err)
					return
				}

				if us.IsAdmin != tt.expectedAdmin {
					t.Errorf("Expected IsAdmin %v, got %v", tt.expectedAdmin, us.IsAdmin)
				}

				c.JSON(http.StatusOK, gin.H{"isAdmin": us.IsAdmin})
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
			}
		})
	}
}

func TestAuthHandler_IsAdminField(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupTestRouter()

	// モックサービス設定（管理者ユーザー）
	mockService := &MockProducerService{
		findProducerFunc: func(ctx context.Context, identityId string) (entity.Producer, error) {
			if identityId == "admin-user-123" {
				return entity.Producer{
					ID:         1,
					IdentityId: identityId,
					IsAdmin:    true,
				}, nil
			}
			return entity.Producer{}, sql.ErrNoRows
		},
	}

	loginHandler := &LoginProducer{
		ProducerService: mockService,
	}

	router.GET("/auth/", func(c *gin.Context) {
		// Google認証情報をモック
		c.Set("user", goauth.Userinfo{
			Id:    "admin-user-123",
			Email: "admin@example.com",
		})
		loginHandler.AuthHandler(c)
	})

	// テスト実行
	req, _ := http.NewRequest("GET", "/auth/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// セッションからユーザー情報を取得して検証
	router.GET("/verify", func(c *gin.Context) {
		us, err := getUserSession(c)
		if err != nil {
			t.Errorf("getUserSession returned error: %v", err)
			return
		}

		if !us.IsAdmin {
			t.Errorf("Expected IsAdmin to be true for admin user")
		}

		c.JSON(http.StatusOK, gin.H{"isAdmin": us.IsAdmin})
	})

	// 検証のための2回目のリクエスト
	req2, _ := http.NewRequest("GET", "/verify", nil)
	// セッションCookieを引き継ぐ
	if cookies := w.Result().Cookies(); len(cookies) > 0 {
		for _, cookie := range cookies {
			req2.AddCookie(cookie)
		}
	}

	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w2.Code)
	}
}