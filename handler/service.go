package handler

import (
	"context"

	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/service"
)

type SceneService interface {
	ListScene(context.Context, *service.ListSceneRequest) ([]entity.Scene, error)
	RegistScene(context.Context, *service.RegistSceneRequest) error
}
type MemberService interface {
	ListMember(context.Context) ([]entity.Member, error)
	GetMemberByGroup(context.Context, int) ([]entity.Member, error)
}
type PhotographService interface {
	ListPhotograph(context.Context) ([]entity.Photograph, error)
	GetPhotographByGroup(context.Context, int) ([]entity.Photograph, error)
}
