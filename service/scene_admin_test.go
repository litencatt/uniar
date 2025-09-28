package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
)

func TestSceneGetByID(t *testing.T) {
	db, _, _ := sqlmock.New()
	testCases := []struct {
		name     string
		id       int64
		expected *entity.SceneWithDetails
	}{
		{
			name: "ValidSceneID",
			id:   1,
			expected: &entity.SceneWithDetails{
				ID:             1,
				PhotographID:   1,
				MemberID:       1,
				ColorTypeID:    1,
				VocalMax:       100,
				DanceMax:       200,
				PerformanceMax: 300,
				CenterSkill:    "Test Skill",
				ExpectedValue:  "50.5",
				SsrPlus:        1,
				PhotographName: "Test Photograph",
				MemberName:     "Test Member",
				ColorName:      "Test Color",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			q.EXPECT().GetSceneById(ctx, gomock.Any(), tc.id).Return(repository.GetSceneByIdRow{
				ID:             tc.expected.ID,
				PhotographID:   tc.expected.PhotographID,
				MemberID:       tc.expected.MemberID,
				ColorTypeID:    tc.expected.ColorTypeID,
				VocalMax:       tc.expected.VocalMax,
				DanceMax:       tc.expected.DanceMax,
				PerformanceMax: tc.expected.PerformanceMax,
				CenterSkill:    sql.NullString{String: tc.expected.CenterSkill, Valid: true},
				ExpectedValue:  sql.NullString{String: tc.expected.ExpectedValue, Valid: true},
				SsrPlus:        tc.expected.SsrPlus,
				PhotographName: tc.expected.PhotographName,
				MemberName:     tc.expected.MemberName,
				ColorName:      tc.expected.ColorName,
			}, nil)

			svc := &Scene{
				DB:      db,
				Querier: q,
			}

			res, err := svc.GetByID(ctx, tc.id)
			if err != nil {
				t.Fatal(err)
			}

			if res.ID != tc.expected.ID {
				t.Errorf("Expected ID %d, got %d", tc.expected.ID, res.ID)
			}
			if res.PhotographName != tc.expected.PhotographName {
				t.Errorf("Expected PhotographName %s, got %s", tc.expected.PhotographName, res.PhotographName)
			}
			if res.CenterSkill != tc.expected.CenterSkill {
				t.Errorf("Expected CenterSkill %s, got %s", tc.expected.CenterSkill, res.CenterSkill)
			}
		})
	}
}

func TestSceneListForAdmin(t *testing.T) {
	db, _, _ := sqlmock.New()
	testCases := []struct {
		name     string
		limit    int64
		offset   int64
		expected []entity.SceneWithDetails
	}{
		{
			name:   "BasicSceneList",
			limit:  10,
			offset: 0,
			expected: []entity.SceneWithDetails{
				{
					ID:             1,
					PhotographID:   1,
					MemberID:       1,
					ColorTypeID:    1,
					VocalMax:       100,
					DanceMax:       200,
					PerformanceMax: 300,
					CenterSkill:    "Test Skill",
					ExpectedValue:  "50.5",
					SsrPlus:        1,
					PhotographName: "Test Photograph",
					MemberName:     "Test Member",
					ColorName:      "Test Color",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			expectedParams := repository.GetSceneListForAdminParams{
				Limit:  tc.limit,
				Offset: tc.offset,
			}

			q.EXPECT().GetSceneListForAdmin(ctx, gomock.Any(), expectedParams).Return([]repository.GetSceneListForAdminRow{
				{
					ID:             tc.expected[0].ID,
					PhotographID:   tc.expected[0].PhotographID,
					MemberID:       tc.expected[0].MemberID,
					ColorTypeID:    tc.expected[0].ColorTypeID,
					VocalMax:       tc.expected[0].VocalMax,
					DanceMax:       tc.expected[0].DanceMax,
					PerformanceMax: tc.expected[0].PerformanceMax,
					CenterSkill:    sql.NullString{String: tc.expected[0].CenterSkill, Valid: true},
					ExpectedValue:  sql.NullString{String: tc.expected[0].ExpectedValue, Valid: true},
					SsrPlus:        tc.expected[0].SsrPlus,
					PhotographName: tc.expected[0].PhotographName,
					MemberName:     tc.expected[0].MemberName,
					ColorName:      tc.expected[0].ColorName,
				},
			}, nil)

			svc := &Scene{
				DB:      db,
				Querier: q,
			}

			res, err := svc.ListForAdmin(ctx, tc.limit, tc.offset)
			if err != nil {
				t.Fatal(err)
			}

			if len(res) != len(tc.expected) {
				t.Errorf("Expected %d scene items, got %d", len(tc.expected), len(res))
			}

			if len(res) > 0 {
				if res[0].PhotographName != tc.expected[0].PhotographName {
					t.Errorf("Expected PhotographName %s, got %s", tc.expected[0].PhotographName, res[0].PhotographName)
				}
				if res[0].MemberName != tc.expected[0].MemberName {
					t.Errorf("Expected MemberName %s, got %s", tc.expected[0].MemberName, res[0].MemberName)
				}
			}
		})
	}
}

func TestSceneCountForAdmin(t *testing.T) {
	db, _, _ := sqlmock.New()
	testCases := []struct {
		name     string
		expected int64
	}{
		{
			name:     "BasicCount",
			expected: 100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			q.EXPECT().CountScenesForAdmin(ctx, gomock.Any()).Return(tc.expected, nil)

			svc := &Scene{
				DB:      db,
				Querier: q,
			}

			res, err := svc.CountForAdmin(ctx)
			if err != nil {
				t.Fatal(err)
			}

			if res != tc.expected {
				t.Errorf("Expected count %d, got %d", tc.expected, res)
			}
		})
	}
}

func TestSceneUpdate(t *testing.T) {
	db, _, _ := sqlmock.New()
	testCases := []struct {
		name   string
		id     int64
		params UpdateSceneParams
	}{
		{
			name: "ValidUpdate",
			id:   1,
			params: UpdateSceneParams{
				PhotographID:   2,
				MemberID:       2,
				ColorTypeID:    2,
				VocalMax:       150,
				DanceMax:       250,
				PerformanceMax: 350,
				CenterSkill:    "Updated Skill",
				ExpectedValue:  "75.5",
				SsrPlus:        1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			expectedParams := repository.UpdateSceneParams{
				ID:             tc.id,
				PhotographID:   tc.params.PhotographID,
				MemberID:       tc.params.MemberID,
				ColorTypeID:    tc.params.ColorTypeID,
				VocalMax:       tc.params.VocalMax,
				DanceMax:       tc.params.DanceMax,
				PerformanceMax: tc.params.PerformanceMax,
				CenterSkill:    sql.NullString{String: tc.params.CenterSkill, Valid: true},
				ExpectedValue:  sql.NullString{String: tc.params.ExpectedValue, Valid: true},
				SsrPlus:        tc.params.SsrPlus,
			}

			q.EXPECT().UpdateScene(ctx, gomock.Any(), expectedParams).Return(nil)

			svc := &Scene{
				DB:      db,
				Querier: q,
			}

			err := svc.Update(ctx, tc.id, tc.params)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestSceneDelete(t *testing.T) {
	db, _, _ := sqlmock.New()
	testCases := []struct {
		name string
		id   int64
	}{
		{
			name: "ValidDelete",
			id:   1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			q.EXPECT().DeleteScene(ctx, gomock.Any(), tc.id).Return(nil)

			svc := &Scene{
				DB:      db,
				Querier: q,
			}

			err := svc.Delete(ctx, tc.id)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}