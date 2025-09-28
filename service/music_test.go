package service

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
)

func TestMusicGetByID(t *testing.T) {
	db, _, _ := sqlmock.New()
	testCases := []struct {
		name     string
		id       int64
		expected *entity.Music
	}{
		{
			name: "ValidMusicID",
			id:   1,
			expected: &entity.Music{
				ID:          1,
				Name:        "Test Music",
				Normal:      100,
				Pro:         200,
				Master:      300,
				Length:      180,
				ColorTypeID: 1,
				LiveID:      1,
				MusicBonus:  50,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			q.EXPECT().GetMusicById(ctx, gomock.Any(), tc.id).Return(repository.Music{
				ID:          tc.expected.ID,
				Name:        tc.expected.Name,
				Normal:      tc.expected.Normal,
				Pro:         tc.expected.Pro,
				Master:      tc.expected.Master,
				Length:      tc.expected.Length,
				ColorTypeID: tc.expected.ColorTypeID,
				LiveID:      tc.expected.LiveID,
				MusicBonus:  sql.NullInt64{Int64: tc.expected.MusicBonus, Valid: true},
				CreatedAt:   time.Now(),
			}, nil)

			svc := &Music{
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
			if res.MusicBonus != tc.expected.MusicBonus {
				t.Errorf("Expected MusicBonus %d, got %d", tc.expected.MusicBonus, res.MusicBonus)
			}
		})
	}
}

func TestMusicListAll(t *testing.T) {
	db, _, _ := sqlmock.New()
	testCases := []struct {
		name     string
		expected []entity.MusicWithDetails
	}{
		{
			name: "BasicMusicList",
			expected: []entity.MusicWithDetails{
				{
					ID:          1,
					Name:        "Test Music 1",
					Normal:      100,
					Pro:         200,
					Master:      300,
					Length:      180,
					ColorTypeID: 1,
					LiveID:      1,
					LiveName:    "Test Live",
					ColorName:   "Test Color",
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

			q.EXPECT().GetMusicListAll(ctx, gomock.Any()).Return([]repository.GetMusicListAllRow{
				{
					ID:          tc.expected[0].ID,
					Name:        tc.expected[0].Name,
					Normal:      tc.expected[0].Normal,
					Pro:         tc.expected[0].Pro,
					Master:      tc.expected[0].Master,
					Length:      tc.expected[0].Length,
					ColorTypeID: tc.expected[0].ColorTypeID,
					LiveID:      tc.expected[0].LiveID,
					LiveName:    tc.expected[0].LiveName,
					ColorName:   tc.expected[0].ColorName,
					CreatedAt:   time.Now(),
				},
			}, nil)

			svc := &Music{
				DB:      db,
				Querier: q,
			}

			res, err := svc.ListAll(ctx)
			if err != nil {
				t.Fatal(err)
			}

			if len(res) != len(tc.expected) {
				t.Errorf("Expected %d music items, got %d", len(tc.expected), len(res))
			}

			if len(res) > 0 {
				if res[0].Name != tc.expected[0].Name {
					t.Errorf("Expected Name %s, got %s", tc.expected[0].Name, res[0].Name)
				}
				if res[0].LiveName != tc.expected[0].LiveName {
					t.Errorf("Expected LiveName %s, got %s", tc.expected[0].LiveName, res[0].LiveName)
				}
			}
		})
	}
}

func TestMusicUpdate(t *testing.T) {
	db, _, _ := sqlmock.New()
	testCases := []struct {
		name   string
		id     int64
		params UpdateMusicParams
	}{
		{
			name: "ValidUpdate",
			id:   1,
			params: UpdateMusicParams{
				Name:        "Updated Music",
				Normal:      150,
				Pro:         250,
				Master:      350,
				Length:      200,
				ColorTypeID: 2,
				LiveID:      2,
				MusicBonus:  75,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			expectedParams := repository.UpdateMusicParams{
				ID:          tc.id,
				Name:        tc.params.Name,
				Normal:      tc.params.Normal,
				Pro:         tc.params.Pro,
				Master:      tc.params.Master,
				Length:      tc.params.Length,
				ColorTypeID: tc.params.ColorTypeID,
				LiveID:      tc.params.LiveID,
				MusicBonus:  sql.NullInt64{Int64: tc.params.MusicBonus, Valid: true},
			}

			q.EXPECT().UpdateMusic(ctx, gomock.Any(), expectedParams).Return(nil)

			svc := &Music{
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

func TestMusicDelete(t *testing.T) {
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

			q.EXPECT().DeleteMusic(ctx, gomock.Any(), tc.id).Return(nil)

			svc := &Music{
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