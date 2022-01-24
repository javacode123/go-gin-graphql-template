package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/javacode_123/go-gin-graphql-template/graphql/model"
)

type groupResolver struct {
	group *model.Group
}

func (r *groupResolver) GID() *graphql.ID {
	gid := graphql.ID(r.group.GID)
	return &gid
}

func (r *groupResolver) Name() *string {
	return &r.group.Name
}
