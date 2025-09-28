package service

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
)

func TestPhotographGetByID(t *testing.T) {
	db, _, _ := sqlmock.New()
	testCases := []struct {
		name     string
		id       int64
		expected *entity.Photograph
	}{
		{
			name: "ValidPhotographID",
			id:   1,
			expected: &entity.Photograph{
				ID:           1,
				Name:         "Test Photograph",
				Abbreviation: "TP",
				PhotoType:    "SSR",
				GroupID:      1,
				ReleasedAt:   time.Now(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			q.EXPECT().GetPhotographById(ctx, gomock.Any(), tc.id).Return(repository.Photograph{
				ID:           tc.expected.ID,
				Name:         tc.expected.Name,
				Abbreviation: tc.expected.Abbreviation,
				PhotoType:    tc.expected.PhotoType,
				GroupID:      tc.expected.GroupID,
				ReleasedAt:   tc.expected.ReleasedAt,
				CreatedAt:    time.Now(),
			}, nil)

			svc := &Photgraph{
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
			if res.Name != tc.expected.Name {
				t.Errorf("Expected Name %s, got %s", tc.expected.Name, res.Name)
			}
			if res.PhotoType != tc.expected.PhotoType {
				t.Errorf("Expected PhotoType %s, got %s", tc.expected.PhotoType, res.PhotoType)
			}
		})
	}
}

func TestPhotographListAllForAdmin(t *testing.T) {
	db, _, _ := sqlmock.New()
	testCases := []struct {
		name     string
		expected []entity.PhotographWithDetails
	}{
		{
			name: "BasicPhotographList",
			expected: []entity.PhotographWithDetails{
				{
					ID:           1,
					Name:         "Test Photograph 1",
					Abbreviation: "TP1",
					PhotoType:    "SSR",
					GroupID:      1,
					ReleasedAt:   time.Now(),
					NameForOrder: "test_photograph_1",
					GroupName:    "Test Group",
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

			q.EXPECT().GetPhotographListForAdmin(ctx, gomock.Any()).Return([]repository.GetPhotographListForAdminRow{
				{
					ID:           tc.expected[0].ID,
					Name:         tc.expected[0].Name,
					Abbreviation: tc.expected[0].Abbreviation,
					PhotoType:    tc.expected[0].PhotoType,
					GroupID:      tc.expected[0].GroupID,
					ReleasedAt:   tc.expected[0].ReleasedAt,
					NameForOrder: tc.expected[0].NameForOrder,
					GroupName:    tc.expected[0].GroupName,
					CreatedAt:    time.Now(),
				},
			}, nil)

			svc := &Photgraph{
				DB:      db,
				Querier: q,
			}

			res, err := svc.ListAllForAdmin(ctx)
			if err != nil {
				t.Fatal(err)
			}

			if len(res) != len(tc.expected) {
				t.Errorf("Expected %d photograph items, got %d", len(tc.expected), len(res))
			}

			if len(res) > 0 {
				if res[0].Name != tc.expected[0].Name {
					t.Errorf("Expected Name %s, got %s", tc.expected[0].Name, res[0].Name)
				}
				if res[0].GroupName != tc.expected[0].GroupName {
					t.Errorf("Expected GroupName %s, got %s", tc.expected[0].GroupName, res[0].GroupName)
				}
			}
		})
	}
}

func TestPhotographUpdate(t *testing.T) {
	db, _, _ := sqlmock.New()
	testCases := []struct {
		name   string
		id     int64
		params UpdatePhotographParams
	}{
		{
			name: "ValidUpdate",
			id:   1,
			params: UpdatePhotographParams{
				Name:         "Updated Photograph",
				GroupID:      2,
				PhotoType:    "SR",
				Abbreviation: "UP",
				NameForOrder: "updated_photograph",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			expectedParams := repository.UpdatePhotographParams{
				ID:           tc.id,
				Name:         tc.params.Name,
				GroupID:      tc.params.GroupID,
				PhotoType:    tc.params.PhotoType,
				Abbreviation: tc.params.Abbreviation,
				NameForOrder: tc.params.NameForOrder,
			}

			q.EXPECT().UpdatePhotograph(ctx, gomock.Any(), expectedParams).Return(nil)

			svc := &Photgraph{
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

func TestPhotographDelete(t *testing.T) {
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

			q.EXPECT().DeletePhotograph(ctx, gomock.Any(), tc.id).Return(nil)

			svc := &Photgraph{
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