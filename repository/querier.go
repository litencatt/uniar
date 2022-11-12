// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package repository

import (
	"context"
	"database/sql"
)

type Querier interface {
	GetAllScenes(ctx context.Context, db DBTX) ([]int64, error)
	GetGroup(ctx context.Context, db DBTX) ([]GetGroupRow, error)
	GetGroupNameById(ctx context.Context, db DBTX, id int64) (string, error)
	GetLiveList(ctx context.Context, db DBTX) ([]GetLiveListRow, error)
	GetMemberList(ctx context.Context, db DBTX, groupID int64) ([]GetMemberListRow, error)
	GetMembers(ctx context.Context, db DBTX) ([]GetMembersRow, error)
	GetMusicList(ctx context.Context, db DBTX) ([]GetMusicListRow, error)
	GetMusicListWithColor(ctx context.Context, db DBTX, name string) ([]GetMusicListWithColorRow, error)
	GetPhotographList(ctx context.Context, db DBTX, arg GetPhotographListParams) ([]GetPhotographListRow, error)
	GetPhotographListByPhotoType(ctx context.Context, db DBTX, photoType string) ([]GetPhotographListByPhotoTypeRow, error)
	GetProducerMember(ctx context.Context, db DBTX) ([]GetProducerMemberRow, error)
	GetProducerOffice(ctx context.Context, db DBTX) (sql.NullInt64, error)
	GetProducerScenes(ctx context.Context, db DBTX) ([]GetProducerScenesRow, error)
	GetScenesWithColor(ctx context.Context, db DBTX, arg GetScenesWithColorParams) ([]GetScenesWithColorRow, error)
	InsertOrUpdateProducerScene(ctx context.Context, db DBTX, arg InsertOrUpdateProducerSceneParams) error
	RegistLive(ctx context.Context, db DBTX, name string) error
	RegistMusic(ctx context.Context, db DBTX, arg RegistMusicParams) error
	RegistPhotograph(ctx context.Context, db DBTX, arg RegistPhotographParams) error
	RegistProducerScene(ctx context.Context, db DBTX, arg RegistProducerSceneParams) error
	RegistScene(ctx context.Context, db DBTX, arg RegistSceneParams) error
	UpdateProducerMember(ctx context.Context, db DBTX, arg UpdateProducerMemberParams) error
	UpdateProducerOffice(ctx context.Context, db DBTX, officeBonus sql.NullInt64) error
}

var _ Querier = (*Queries)(nil)
