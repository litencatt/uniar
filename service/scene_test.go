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

func TestListScene(t *testing.T) {
	db, _, _ := sqlmock.New()
	testCases := []struct {
		name string
		req  ListSceneRequest
		exp  []entity.Scene
	}{
		{
			name: "DefaultRequest",
			req: ListSceneRequest{
				Color:      "%",
				Member:     "%",
				Photograph: "%",
				ProducerID: 1,
			},
			exp: []entity.Scene{
				{
					Photograph: "キュン",
					Member:     "加藤史帆",
					Color:      "Blue",
					Total:      0,
					All35:      1,
					VoDa50:     1,
					DaPe50:     1,
					VoPe50:     1,
					Vo85:       1,
					Da85:       1,
					Pe85:       1,
					Expect:     float32(0),
					SsrPlus:    true,
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
			q.EXPECT().GetScenesWithColor(ctx, gomock.Any(), gomock.Any()).Return([]repository.GetScenesWithColorRow{
				{
					Photograph:     "キュン",
					Member:         "加藤史帆",
					Color:          "Blue",
					Total:          0,
					VocalMax:       0,
					DanceMax:       0,
					PerformanceMax: 0,
					ExpectedValue:  sql.NullString{String: "0", Valid: true},
					SsrPlus:        1,
					Bonds:          sql.NullInt64{Int64: 0, Valid: false},
					Discography:    sql.NullInt64{Int64: 0, Valid: false},
				},
			}, nil)
			q.EXPECT().GetProducerScenesWithProducerId(ctx, gomock.Any(), int64(1)).Return([]repository.GetProducerScenesWithProducerIdRow{}, nil)
			svc := &Scene{
				DB:      db,
				Querier: q,
			}

			res, err := svc.ListScene(ctx, &tc.req)
			if err != nil {
				t.Fatal(err)
			}
			if res[0] != tc.exp[0] {
				t.Errorf("Expected %+v, got %+v instead", tc.exp[0], res[0])
			}
		})
	}
}
