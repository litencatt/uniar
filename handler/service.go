package handler

import (
	"context"

	"github.com/litencatt/uniar/entity"
	"github.com/litencatt/uniar/service"
)

type ListSceneService interface {
	ListScene(context.Context, *service.ListSceneRequest) ([]entity.Scene, error)
}
type ListMemberService interface {
	ListMember(context.Context) ([]entity.Member, error)
}
type ListPhotographService interface {
	ListPhotograph(context.Context) ([]entity.Photograph, error)
}
