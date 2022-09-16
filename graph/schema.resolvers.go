package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"gqlgen-todo/graph/generated"
	"gqlgen-todo/internal/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	var user = model.User{}
	model.DB.Find(&model.User{}).Scan(&user)

	if user.ID <= 0 {
		return nil, errors.New("user not found")
	}
	return nil, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user model.User
	model.DB.Find(&model.User{}).Where("name = ?", input.Name).Scan(&user)

	if user.ID > 0 {
		return nil, errors.New("user exists")
	}

	var newUser = model.User{
		Name: input.Name,
	}
	model.DB.Create(&newUser)

	return &newUser, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos = []*model.Todo{}
	model.DB.Find(&model.Todo{}).Scan(&todos)

	return todos, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	var user *model.User
	model.DB.Find(&model.User{}).Scan(&user)
	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
