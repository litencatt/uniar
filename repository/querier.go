// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package repository

import (
	"context"
)

type Querier interface {
	GetAllMembers(ctx context.Context, db DBTX) ([]Member, error)
	GetAllScenes(ctx context.Context, db DBTX) ([]GetAllScenesRow, error)
	GetAllScenesWithGroupId(ctx context.Context, db DBTX, groupID int64) ([]GetAllScenesWithGroupIdRow, error)
	GetGroup(ctx context.Context, db DBTX) ([]GetGroupRow, error)
	GetGroupNameById(ctx context.Context, db DBTX, id int64) (string, error)
	GetLiveList(ctx context.Context, db DBTX) ([]GetLiveListRow, error)
	GetMembers(ctx context.Context, db DBTX) ([]GetMembersRow, error)
	GetMembersByGroup(ctx context.Context, db DBTX, groupID int64) ([]GetMembersByGroupRow, error)
	GetMusicList(ctx context.Context, db DBTX) ([]GetMusicListRow, error)
	GetMusicListWithColor(ctx context.Context, db DBTX, name string) ([]GetMusicListWithColorRow, error)
	GetPhotographByGroupId(ctx context.Context, db DBTX, groupID int64) ([]GetPhotographByGroupIdRow, error)
	GetPhotographList(ctx context.Context, db DBTX, arg GetPhotographListParams) ([]GetPhotographListRow, error)
	GetPhotographListAll(ctx context.Context, db DBTX) ([]GetPhotographListAllRow, error)
	GetPhotographListByPhotoType(ctx context.Context, db DBTX, photoType string) ([]GetPhotographListByPhotoTypeRow, error)
	GetProducer(ctx context.Context, db DBTX, identityID string) (Producer, error)
	GetProducerMember(ctx context.Context, db DBTX, producerID int64) ([]GetProducerMemberRow, error)
	GetProducerOffice(ctx context.Context, db DBTX, producerID int64) (ProducerOffice, error)
	GetProducerScenes(ctx context.Context, db DBTX, arg GetProducerScenesParams) ([]GetProducerScenesRow, error)
	GetProducerScenesByGroupId(ctx context.Context, db DBTX, groupID int64) ([]GetProducerScenesByGroupIdRow, error)
	GetProducerScenesWithProducerId(ctx context.Context, db DBTX, producerID int64) ([]GetProducerScenesWithProducerIdRow, error)
	GetScenesWithColor(ctx context.Context, db DBTX, arg GetScenesWithColorParams) ([]GetScenesWithColorRow, error)
	GetScenesWithGroupId(ctx context.Context, db DBTX, arg GetScenesWithGroupIdParams) ([]GetScenesWithGroupIdRow, error)
	GetSsrPlusReleasedPhotographList(ctx context.Context, db DBTX) ([]int64, error)
	InitProducerSceneAll(ctx context.Context, db DBTX, arg InitProducerSceneAllParams) error
	RegistLive(ctx context.Context, db DBTX, name string) error
	RegistMusic(ctx context.Context, db DBTX, arg RegistMusicParams) error
	RegistPhotograph(ctx context.Context, db DBTX, arg RegistPhotographParams) error
	RegistProducer(ctx context.Context, db DBTX, arg RegistProducerParams) error
	RegistProducerMember(ctx context.Context, db DBTX, arg RegistProducerMemberParams) error
	RegistProducerOffice(ctx context.Context, db DBTX, producerID int64) error
	RegistProducerScene(ctx context.Context, db DBTX, arg RegistProducerSceneParams) error
	RegistScene(ctx context.Context, db DBTX, arg RegistSceneParams) error
	UpdateProducerMember(ctx context.Context, db DBTX, arg UpdateProducerMemberParams) error
	UpdateProducerOffice(ctx context.Context, db DBTX, arg UpdateProducerOfficeParams) error
	UpdateProducerScene(ctx context.Context, db DBTX, arg UpdateProducerSceneParams) error
}

var _ Querier = (*Queries)(nil)
