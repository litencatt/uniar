package handler

import (
	"context"
	"encoding/gob"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/entity"
)

// MockMusicService は MusicService のモック実装
type MockMusicService struct {
	listAllFunc func(ctx context.Context) ([]entity.MusicWithDetails, error)
	getByIDFunc func(ctx context.Context, id int64) (*entity.Music, error)
	updateFunc  func(ctx context.Context, id int64, params interface{}) error
	deleteFunc  func(ctx context.Context, id int64) error
}

func (m *MockMusicService) ListAll(ctx context.Context) ([]entity.MusicWithDetails, error) {
	if m.listAllFunc != nil {
		return m.listAllFunc(ctx)
	}
	return []entity.MusicWithDetails{}, nil
}

func (m *MockMusicService) GetByID(ctx context.Context, id int64) (*entity.Music, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(ctx, id)
	}
	return &entity.Music{}, nil
}

func (m *MockMusicService) Update(ctx context.Context, id int64, params interface{}) error {
	if m.updateFunc != nil {
		return m.updateFunc(ctx, id, params)
	}
	return nil
}

func (m *MockMusicService) Delete(ctx context.Context, id int64) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(ctx, id)
	}
	return nil
}

// MockPhotographService は PhotographService のモック実装
type MockPhotographService struct {
	listAllForAdminFunc func(ctx context.Context) ([]entity.PhotographWithDetails, error)
	getByIDFunc         func(ctx context.Context, id int64) (*entity.Photograph, error)
	updateFunc          func(ctx context.Context, id int64, params interface{}) error
	deleteFunc          func(ctx context.Context, id int64) error
}

func (m *MockPhotographService) ListAllForAdmin(ctx context.Context) ([]entity.PhotographWithDetails, error) {
	if m.listAllForAdminFunc != nil {
		return m.listAllForAdminFunc(ctx)
	}
	return []entity.PhotographWithDetails{}, nil
}

func (m *MockPhotographService) GetByID(ctx context.Context, id int64) (*entity.Photograph, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(ctx, id)
	}
	return &entity.Photograph{}, nil
}

func (m *MockPhotographService) Update(ctx context.Context, id int64, params interface{}) error {
	if m.updateFunc != nil {
		return m.updateFunc(ctx, id, params)
	}
	return nil
}

func (m *MockPhotographService) Delete(ctx context.Context, id int64) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(ctx, id)
	}
	return nil
}

// MockSceneService は SceneService のモック実装
type MockSceneService struct {
	listForAdminFunc func(ctx context.Context, limit, offset int64) ([]entity.SceneWithDetails, error)
	countForAdminFunc func(ctx context.Context) (int64, error)
	getByIDFunc      func(ctx context.Context, id int64) (*entity.SceneWithDetails, error)
	updateFunc       func(ctx context.Context, id int64, params interface{}) error
	deleteFunc       func(ctx context.Context, id int64) error
}

func (m *MockSceneService) ListForAdmin(ctx context.Context, limit, offset int64) ([]entity.SceneWithDetails, error) {
	if m.listForAdminFunc != nil {
		return m.listForAdminFunc(ctx, limit, offset)
	}
	return []entity.SceneWithDetails{}, nil
}

func (m *MockSceneService) CountForAdmin(ctx context.Context) (int64, error) {
	if m.countForAdminFunc != nil {
		return m.countForAdminFunc(ctx)
	}
	return 0, nil
}

func (m *MockSceneService) GetByID(ctx context.Context, id int64) (*entity.SceneWithDetails, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(ctx, id)
	}
	return &entity.SceneWithDetails{}, nil
}

func (m *MockSceneService) Update(ctx context.Context, id int64, params interface{}) error {
	if m.updateFunc != nil {
		return m.updateFunc(ctx, id, params)
	}
	return nil
}

func (m *MockSceneService) Delete(ctx context.Context, id int64) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(ctx, id)
	}
	return nil
}

// setupAdminTestRouter はAdmin用のテストルーターをセットアップ
func setupAdminTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// セッション設定
	gob.Register(&UserSession{})
	store := cookie.NewStore([]byte("test-secret"))
	router.Use(sessions.Sessions("test-session", store))

	// 管理者ミドルウェアのテスト版
	router.Use(func(c *gin.Context) {
		c.Set("isAdmin", true)
		c.Next()
	})

	return router
}

func TestAdminDashboard(t *testing.T) {
	router := setupAdminTestRouter()

	adminDashboard := &AdminDashboardHandler{}
	router.GET("/admin/", adminDashboard.Dashboard)

	req, _ := http.NewRequest("GET", "/admin/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestAdminMusicList(t *testing.T) {
	router := setupAdminTestRouter()

	mockService := &MockMusicService{
		listAllFunc: func(ctx context.Context) ([]entity.MusicWithDetails, error) {
			return []entity.MusicWithDetails{
				{
					ID:        1,
					Name:      "Test Music",
					LiveName:  "Test Live",
					ColorName: "Test Color",
				},
			}, nil
		},
	}

	adminMusic := &AdminMusic{MusicService: mockService}
	router.GET("/admin/music", adminMusic.ListMusic)

	req, _ := http.NewRequest("GET", "/admin/music", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestAdminMusicEdit(t *testing.T) {
	router := setupAdminTestRouter()

	mockService := &MockMusicService{
		getByIDFunc: func(ctx context.Context, id int64) (*entity.Music, error) {
			if id == 1 {
				return &entity.Music{
					ID:   1,
					Name: "Test Music",
				}, nil
			}
			return nil, nil
		},
	}

	adminMusic := &AdminMusic{MusicService: mockService}
	router.GET("/admin/music/:id/edit", adminMusic.EditMusic)

	req, _ := http.NewRequest("GET", "/admin/music/1/edit", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestAdminMusicUpdate(t *testing.T) {
	router := setupAdminTestRouter()

	mockService := &MockMusicService{
		updateFunc: func(ctx context.Context, id int64, params interface{}) error {
			return nil
		},
	}

	adminMusic := &AdminMusic{MusicService: mockService}
	router.POST("/admin/music/:id", adminMusic.UpdateMusic)

	form := url.Values{}
	form.Add("name", "Updated Music")
	form.Add("normal", "100")
	form.Add("pro", "200")
	form.Add("master", "300")
	form.Add("length", "180")
	form.Add("color_type_id", "1")
	form.Add("live_id", "1")
	form.Add("music_bonus", "50")

	req, _ := http.NewRequest("POST", "/admin/music/1", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Errorf("Expected status code %d, got %d", http.StatusFound, w.Code)
	}
}

func TestAdminPhotographList(t *testing.T) {
	router := setupAdminTestRouter()

	mockService := &MockPhotographService{
		listAllForAdminFunc: func(ctx context.Context) ([]entity.PhotographWithDetails, error) {
			return []entity.PhotographWithDetails{
				{
					ID:        1,
					Name:      "Test Photograph",
					GroupName: "Test Group",
				},
			}, nil
		},
	}

	adminPhotograph := &AdminPhotograph{PhotographService: mockService}
	router.GET("/admin/photograph", adminPhotograph.ListPhotograph)

	req, _ := http.NewRequest("GET", "/admin/photograph", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestAdminSceneList(t *testing.T) {
	router := setupAdminTestRouter()

	mockService := &MockSceneService{
		listForAdminFunc: func(ctx context.Context, limit, offset int64) ([]entity.SceneWithDetails, error) {
			return []entity.SceneWithDetails{
				{
					ID:             1,
					PhotographName: "Test Photograph",
					MemberName:     "Test Member",
				},
			}, nil
		},
		countForAdminFunc: func(ctx context.Context) (int64, error) {
			return 1, nil
		},
	}

	adminScene := &AdminScene{SceneService: mockService}
	router.GET("/admin/scene", adminScene.ListScene)

	req, _ := http.NewRequest("GET", "/admin/scene", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestAdminMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// セッション設定
	gob.Register(&UserSession{})
	store := cookie.NewStore([]byte("test-secret"))
	router.Use(sessions.Sessions("test-session", store))

	testCases := []struct {
		name         string
		isAdmin      bool
		expectedCode int
	}{
		{
			name:         "AdminUser",
			isAdmin:      true,
			expectedCode: http.StatusOK,
		},
		{
			name:         "NonAdminUser",
			isAdmin:      false,
			expectedCode: http.StatusForbidden,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			router.GET("/admin/test", func(c *gin.Context) {
				// セッションに管理者フラグを設定
				session := sessions.Default(c)
				session.Set("user", &UserSession{IsAdmin: tc.isAdmin})
				session.Save()
				c.Next()
			}, AdminCheck(), func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
			})

			req, _ := http.NewRequest("GET", "/admin/test", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tc.expectedCode {
				t.Errorf("Expected status code %d, got %d", tc.expectedCode, w.Code)
			}
		})
	}
}

func TestAdminDeleteOperations(t *testing.T) {
	router := setupAdminTestRouter()

	// Music削除テスト
	musicService := &MockMusicService{
		deleteFunc: func(ctx context.Context, id int64) error {
			return nil
		},
	}
	adminMusic := &AdminMusic{MusicService: musicService}
	router.DELETE("/admin/music/:id", adminMusic.DeleteMusic)

	req, _ := http.NewRequest("DELETE", "/admin/music/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Photograph削除テスト
	photographService := &MockPhotographService{
		deleteFunc: func(ctx context.Context, id int64) error {
			return nil
		},
	}
	adminPhotograph := &AdminPhotograph{PhotographService: photographService}
	router.DELETE("/admin/photograph/:id", adminPhotograph.DeletePhotograph)

	req, _ = http.NewRequest("DELETE", "/admin/photograph/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Scene削除テスト
	sceneService := &MockSceneService{
		deleteFunc: func(ctx context.Context, id int64) error {
			return nil
		},
	}
	adminScene := &AdminScene{SceneService: sceneService}
	router.DELETE("/admin/scene/:id", adminScene.DeleteScene)

	req, _ = http.NewRequest("DELETE", "/admin/scene/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestAdminPagination(t *testing.T) {
	router := setupAdminTestRouter()

	mockService := &MockSceneService{
		listForAdminFunc: func(ctx context.Context, limit, offset int64) ([]entity.SceneWithDetails, error) {
			// ページネーションパラメータの検証
			if limit != 20 || offset != 20 {
				t.Errorf("Expected limit=20, offset=20, got limit=%d, offset=%d", limit, offset)
			}
			return []entity.SceneWithDetails{}, nil
		},
		countForAdminFunc: func(ctx context.Context) (int64, error) {
			return 100, nil
		},
	}

	adminScene := &AdminScene{SceneService: mockService}
	router.GET("/admin/scene", adminScene.ListScene)

	req, _ := http.NewRequest("GET", "/admin/scene?page=2", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}