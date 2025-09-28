package service

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/repository"
)

func TestListMember(t *testing.T) {
	db, _, _ := sqlmock.New()
	testCases := []struct {
		name     string
		expected []entity.Member
	}{
		{
			name: "BasicMemberList",
			expected: []entity.Member{
				{
					Name: "加藤史帆",
				},
				{
					Name: "山﨑天",
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

			q.EXPECT().GetAllMembers(ctx, gomock.Any()).Return([]repository.Member{
				{
					Name: "加藤史帆",
				},
				{
					Name: "山﨑天",
				},
			}, nil)

			svc := &Member{
				DB:      db,
				Querier: q,
			}

			res, err := svc.ListMember(ctx)
			if err != nil {
				t.Fatal(err)
			}

			if len(res) != len(tc.expected) {
				t.Errorf("Expected %d members, got %d", len(tc.expected), len(res))
			}

			for i, member := range res {
				if member.Name != tc.expected[i].Name {
					t.Errorf("Expected member name %s, got %s", tc.expected[i].Name, member.Name)
				}
			}
		})
	}
}