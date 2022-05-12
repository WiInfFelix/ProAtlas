package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/WiInfFelix/proatlas/internal/exercises"
	"strconv"

	"github.com/WiInfFelix/proatlas/graph/generated"
	"github.com/WiInfFelix/proatlas/graph/model"
)

func (r *mutationResolver) CreateExercise(ctx context.Context, input model.NewExercise) (*model.Exercise, error) {
	var exer exercises.Exercise
	exer.Name = input.Name
	exer.Description = input.Description
	exer.TargetMuscle = input.TargetMuscle
	exerID := exer.Save()
	return &model.Exercise{
		ID:           strconv.FormatInt(exerID, 10),
		Name:         exer.Name,
		Description:  exer.Description,
		TargetMuscle: exer.TargetMuscle,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Exercises(ctx context.Context) ([]*model.Exercise, error) {
	var resExecs []*model.Exercise
	var dbExec []exercises.Exercise

	dbExec = exercises.GetAll()
	for _, exec := range dbExec {
		resExecs = append(resExecs, &model.Exercise{
			ID:           exec.ID,
			Name:         exec.Name,
			Description:  exec.Description,
			TargetMuscle: exec.TargetMuscle,
			Creator:      nil,
		})
	}

	return resExecs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
