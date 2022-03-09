package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/ssabrinadias/goGraphql/graph/generated"
	"github.com/ssabrinadias/goGraphql/graph/model"
)

func (r *mutationResolver) CreateList(ctx context.Context, input model.NewList) (*model.List, error) {
	list := &model.List{
		Title:  input.Title,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
		Items:  input.Items,
	}
	r.list = append(r.list, list)
	return list, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		Name: input.Name,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		City: input.City,
	}
	r.user = append(r.user, user)
	return user, nil
}

func (r *queryResolver) List(ctx context.Context) ([]*model.List, error) {
	var links []*model.List
	var x *string
	y := "cerveja"
	x = &y
	arrayItems := make([]*string, 0, 10)
	arrayItems = append(arrayItems, x)
	dummyLink := model.List{
		Title:  "Lista de Bebidas",
		UserID: "123",
		Items:  arrayItems,
	}
	links = append(links, &dummyLink)
	return links, nil
}

func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	return []*model.User{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
