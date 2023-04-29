package handler

import (
	"context"

	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/service"
)

type SceneService interface {
	ListScene(context.Context, *service.ListSceneRequest) ([]entity.Scene, error)
	ListSceneAll(context.Context, *service.ListSceneAllRequest) ([]entity.ProducerScene, error)
}
type ProducerSceneService interface {
	ListScene(context.Context, *service.ListProducerSceneRequest) ([]entity.ProducerScene, error)
	RegistScene(context.Context, *service.RegistProducerSceneRequest) error
	InitAllScene(context.Context, *service.InitProducerSceneRequest) error
}
type MemberService interface {
	ListMember(context.Context) ([]entity.Member, error)
	GetMemberByGroup(context.Context, int64) ([]entity.Member, error)
	ListProducerMember(context.Context) ([]entity.ProducerMember, error)
	UpdateProducerMember(context.Context, entity.ProducerMember) error
}
type PhotographService interface {
	ListPhotograph(context.Context) ([]entity.Photograph, error)
	GetPhotographByGroup(context.Context, int64) ([]entity.Photograph, error)
	GetSsrPlusReleasedPhotographList(context.Context) ([]entity.Photograph, error)
}
