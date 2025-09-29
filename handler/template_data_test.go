package handler

import (
	"context"
	"encoding/gob"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/service"
)

// MockSceneService は SceneService のモック実装
type MockSceneService struct {
	listSceneFunc    func(ctx context.Context, req *service.ListSceneRequest) ([]entity.Scene, error)
	listSceneAllFunc func(ctx context.Context, req *service.ListSceneAllRequest) ([]entity.Scene, []entity.ProducerScene, error)
}

func (m *MockSceneService) ListScene(ctx context.Context, req *service.ListSceneRequest) ([]entity.Scene, error) {
	if m.listSceneFunc != nil {
		return m.listSceneFunc(ctx, req)
	}
	return []entity.Scene{}, nil
}

func (m *MockSceneService) ListSceneAll(ctx context.Context, req *service.ListSceneAllRequest) ([]entity.Scene, []entity.ProducerScene, error) {
	if m.listSceneAllFunc != nil {
		return m.listSceneAllFunc(ctx, req)
	}
	return []entity.Scene{}, []entity.ProducerScene{}, nil
}

// MockMemberService は MemberService のモック実装
type MockMemberService struct {
	listMemberFunc         func(ctx context.Context) ([]entity.Member, error)
	listProducerMemberFunc func(ctx context.Context, producerId int64) ([]entity.ProducerMember, error)
}

func (m *MockMemberService) ListMember(ctx context.Context) ([]entity.Member, error) {
	if m.listMemberFunc != nil {
		return m.listMemberFunc(ctx)
	}
	return []entity.Member{}, nil
}

func (m *MockMemberService) ListProducerMember(ctx context.Context, producerId int64) ([]entity.ProducerMember, error) {
	if m.listProducerMemberFunc != nil {
		return m.listProducerMemberFunc(ctx, producerId)
	}
	return []entity.ProducerMember{}, nil
}

func (m *MockMemberService) GetMemberByGroup(ctx context.Context, groupId int64) ([]entity.Member, error) {
	return []entity.Member{}, nil
}

func (m *MockMemberService) UpdateProducerMember(ctx context.Context, pm entity.ProducerMember) error {
	return nil
}

// MockPhotographService は PhotographService のモック実装
type MockPhotographService struct {
	listPhotographFunc                  func(ctx context.Context) ([]entity.Photograph, error)
	getSsrPlusReleasedPhotographListFunc func(ctx context.Context) ([]entity.Photograph, error)
}

func (m *MockPhotographService) ListPhotograph(ctx context.Context) ([]entity.Photograph, error) {
	if m.listPhotographFunc != nil {
		return m.listPhotographFunc(ctx)
	}
	return []entity.Photograph{}, nil
}

func (m *MockPhotographService) GetPhotographByGroup(ctx context.Context, groupId int64) ([]entity.Photograph, error) {
	return []entity.Photograph{}, nil
}

func (m *MockPhotographService) GetSsrPlusReleasedPhotographList(ctx context.Context) ([]entity.Photograph, error) {
	if m.getSsrPlusReleasedPhotographListFunc != nil {
		return m.getSsrPlusReleasedPhotographListFunc(ctx)
	}
	return []entity.Photograph{}, nil
}

// setupRouterWithSession はセッション付きテストルーターをセットアップ
func setupRouterWithSession(userSession *UserSession) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// セッション設定
	gob.Register(&UserSession{})
	store := cookie.NewStore([]byte("test-secret"))
	router.Use(sessions.Sessions("test-session", store))

	// セッションにユーザー情報を設定するミドルウェア
	router.Use(func(c *gin.Context) {
		if userSession != nil {
			session := sessions.Default(c)
			session.Set("uniar_session", userSession)
			session.Save()
		}
		c.Next()
	})

	return router
}

func TestListScene_IsAdminInTemplateData(t *testing.T) {
	tests := []struct {
		name          string
		adminDebug    string
		userIsAdmin   bool
		expectedAdmin bool
	}{
		{
			name:          "Regular user without ADMIN_DEBUG",
			adminDebug:    "",
			userIsAdmin:   false,
			expectedAdmin: false,
		},
		{
			name:          "Admin user without ADMIN_DEBUG",
			adminDebug:    "",
			userIsAdmin:   true,
			expectedAdmin: true,
		},
		{
			name:          "Regular user with ADMIN_DEBUG=true",
			adminDebug:    "true",
			userIsAdmin:   false,
			expectedAdmin: true,
		},
		{
			name:          "Admin user with ADMIN_DEBUG=true",
			adminDebug:    "true",
			userIsAdmin:   true,
			expectedAdmin: true,
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

			// テストユーザーセッション
			userSession := &UserSession{
				ProducerId: 1,
				IdentityId: "test-user",
				EMail:      "test@example.com",
				LoggedIn:   true,
				IsAdmin:    tt.userIsAdmin,
			}

			router := setupRouterWithSession(userSession)

			// モックサービス
			mockSceneService := &MockSceneService{
				listSceneFunc: func(ctx context.Context, req *service.ListSceneRequest) ([]entity.Scene, error) {
					return []entity.Scene{}, nil
				},
			}

			mockMemberService := &MockMemberService{
				listMemberFunc: func(ctx context.Context) ([]entity.Member, error) {
					return []entity.Member{}, nil
				},
			}

			mockPhotographService := &MockPhotographService{
				listPhotographFunc: func(ctx context.Context) ([]entity.Photograph, error) {
					return []entity.Photograph{}, nil
				},
			}

			listScene := &ListScene{
				SceneService:      mockSceneService,
				MemberService:     mockMemberService,
				PhotographService: mockPhotographService,
			}

			// レスポンスを検証するカスタムレンダラー
			var templateData gin.H
			router.HTMLRender = &mockHTMLRender{
				onRender: func(code int, name string, data interface{}) {
					if h, ok := data.(gin.H); ok {
						templateData = h
					}
				},
			}

			router.GET("/auth/scenes", listScene.ListScene)

			// テスト実行
			req, _ := http.NewRequest("GET", "/auth/scenes", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// テンプレートデータのIsAdminフィールドを確認
			if isAdmin, exists := templateData["IsAdmin"]; !exists {
				t.Errorf("IsAdmin field not found in template data")
			} else if isAdmin != tt.expectedAdmin {
				t.Errorf("Expected IsAdmin %v, got %v", tt.expectedAdmin, isAdmin)
			}

			// 他の必須フィールドも確認
			expectedFields := []string{"title", "LoggedIn", "EMail", "IsAdmin"}
			for _, field := range expectedFields {
				if _, exists := templateData[field]; !exists {
					t.Errorf("Required field %s not found in template data", field)
				}
			}
		})
	}
}

func TestListMember_IsAdminInTemplateData(t *testing.T) {
	// テストユーザーセッション
	userSession := &UserSession{
		ProducerId: 1,
		IdentityId: "test-user",
		EMail:      "test@example.com",
		LoggedIn:   true,
		IsAdmin:    false,
	}

	// ADMIN_DEBUG=trueを設定
	originalValue := os.Getenv("ADMIN_DEBUG")
	defer func() {
		if originalValue == "" {
			os.Unsetenv("ADMIN_DEBUG")
		} else {
			os.Setenv("ADMIN_DEBUG", originalValue)
		}
	}()
	os.Setenv("ADMIN_DEBUG", "true")

	router := setupRouterWithSession(userSession)

	// モックサービス
	mockMemberService := &MockMemberService{
		listProducerMemberFunc: func(ctx context.Context, producerId int64) ([]entity.ProducerMember, error) {
			return []entity.ProducerMember{}, nil
		},
	}

	listMember := &ListMember{
		MemberService: mockMemberService,
	}

	// レスポンスを検証するカスタムレンダラー
	var templateData gin.H
	router.HTMLRender = &mockHTMLRender{
		onRender: func(code int, name string, data interface{}) {
			if h, ok := data.(gin.H); ok {
				templateData = h
			}
		},
	}

	router.GET("/auth/members", listMember.ListMember)

	// テスト実行
	req, _ := http.NewRequest("GET", "/auth/members", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// テンプレートデータのIsAdminフィールドを確認（ADMIN_DEBUG=trueなのでtrueになるはず）
	if isAdmin, exists := templateData["IsAdmin"]; !exists {
		t.Errorf("IsAdmin field not found in template data")
	} else if isAdmin != true {
		t.Errorf("Expected IsAdmin true (due to ADMIN_DEBUG=true), got %v", isAdmin)
	}
}

// mockHTMLRender はテスト用のHTMLレンダラー
type mockHTMLRender struct {
	onRender func(code int, name string, data interface{})
}

func (m *mockHTMLRender) Instance(name string, data interface{}) render.Render {
	return &mockRender{
		name:     name,
		data:     data,
		onRender: m.onRender,
	}
}

type mockRender struct {
	name     string
	data     interface{}
	onRender func(code int, name string, data interface{})
}

func (m *mockRender) Render(w http.ResponseWriter) error {
	if m.onRender != nil {
		m.onRender(http.StatusOK, m.name, m.data)
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func (m *mockRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}