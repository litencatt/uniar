package handler

import (
	"context"
	"database/sql"
	"encoding/gob"
	"fmt"
	"net/http"
	"net/http/httptest"
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