package service

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/litencatt/uniar/repository"
)

func TestSearchMusic(t *testing.T) {
	testCases := []struct {
		name   string
		params MusicSearchParams
	}{
		{
			name: "SearchByName",
			params: MusicSearchParams{
				Name:        "テスト楽曲",
				LiveID:      0,
				ColorTypeID: 0,
			},
		},
		{
			name: "SearchByLiveID",
			params: MusicSearchParams{
				Name:        "",
				LiveID:      1,
				ColorTypeID: 0,
			},
		},
		{
			name: "SearchByColorTypeID",
			params: MusicSearchParams{
				Name:        "",
				LiveID:      0,
				ColorTypeID: 2,
			},
		},
		{
			name: "SearchByAllParams",
			params: MusicSearchParams{
				Name:        "テスト楽曲",
				LiveID:      1,
				ColorTypeID: 2,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, _, _ := sqlmock.New()
			defer db.Close()

			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			expectedRows := []repository.SearchMusicListRow{
				{
					ID:        1,
					Name:      "テスト楽曲",
					LiveName:  "1st Live",
					ColorName: "Red",
				},
			}

			q.EXPECT().SearchMusicList(ctx, db, repository.SearchMusicListParams{
				Column1: tc.params.Name,
				Column2: tc.params.LiveID,
				Column3: tc.params.ColorTypeID,
			}).Return(expectedRows, nil)

			svc := &SearchService{
				DB:      db,
				Querier: q,
			}

			result, err := svc.SearchMusic(ctx, tc.params)

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if len(result) != 1 {
				t.Errorf("Expected 1 result, got %d", len(result))
			}

			if result[0].Name != "テスト楽曲" {
				t.Errorf("Expected name 'テスト楽曲', got '%s'", result[0].Name)
			}
		})
	}
}

func TestSearchPhotograph(t *testing.T) {
	testCases := []struct {
		name   string
		params PhotographSearchParams
	}{
		{
			name: "SearchByName",
			params: PhotographSearchParams{
				Name:      "テスト撮影",
				GroupID:   0,
				PhotoType: "",
			},
		},
		{
			name: "SearchByGroupID",
			params: PhotographSearchParams{
				Name:      "",
				GroupID:   1,
				PhotoType: "",
			},
		},
		{
			name: "SearchByPhotoType",
			params: PhotographSearchParams{
				Name:      "",
				GroupID:   0,
				PhotoType: "Live",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, _, _ := sqlmock.New()
			defer db.Close()

			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			expectedRows := []repository.SearchPhotographListRow{
				{
					ID:        1,
					Name:      "テスト撮影",
					GroupName: "Group A",
				},
			}

			q.EXPECT().SearchPhotographList(ctx, db, repository.SearchPhotographListParams{
				Column1: tc.params.Name,
				Column2: tc.params.GroupID,
				Column3: tc.params.PhotoType,
			}).Return(expectedRows, nil)

			svc := &SearchService{
				DB:      db,
				Querier: q,
			}

			result, err := svc.SearchPhotograph(ctx, tc.params)

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if len(result) != 1 {
				t.Errorf("Expected 1 result, got %d", len(result))
			}

			if result[0].Name != "テスト撮影" {
				t.Errorf("Expected name 'テスト撮影', got '%s'", result[0].Name)
			}
		})
	}
}

func TestSearchScene(t *testing.T) {
	testCases := []struct {
		name   string
		params SceneSearchParams
	}{
		{
			name: "SearchByMemberID",
			params: SceneSearchParams{
				MemberID:     1,
				PhotographID: 0,
				ColorTypeID:  0,
				SSRPlus:      -1,
			},
		},
		{
			name: "SearchBySSRPlus",
			params: SceneSearchParams{
				MemberID:     0,
				PhotographID: 0,
				ColorTypeID:  0,
				SSRPlus:      1,
			},
		},
		{
			name: "SearchByAllParams",
			params: SceneSearchParams{
				MemberID:     1,
				PhotographID: 2,
				ColorTypeID:  3,
				SSRPlus:      1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, _, _ := sqlmock.New()
			defer db.Close()

			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			expectedRows := []repository.SearchSceneListRow{
				{
					ID:             1,
					MemberName:     "Member A",
					PhotographName: "Photo A",
					ColorName:      "Red",
				},
			}

			q.EXPECT().SearchSceneList(ctx, db, repository.SearchSceneListParams{
				Column1: tc.params.MemberID,
				Column2: tc.params.PhotographID,
				Column3: tc.params.ColorTypeID,
				Column4: int64(tc.params.SSRPlus),
			}).Return(expectedRows, nil)

			svc := &SearchService{
				DB:      db,
				Querier: q,
			}

			result, err := svc.SearchScene(ctx, tc.params)

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if len(result) != 1 {
				t.Errorf("Expected 1 result, got %d", len(result))
			}

			if result[0].MemberName != "Member A" {
				t.Errorf("Expected member name 'Member A', got '%s'", result[0].MemberName)
			}
		})
	}
}