package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/shuymn-sandbox/go-gqlgen-and-ent/ent"
	"github.com/shuymn-sandbox/go-gqlgen-and-ent/graphql/alpha/graph/generated"
	"github.com/shuymn-sandbox/go-gqlgen-and-ent/graphql/alpha/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*ent.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
