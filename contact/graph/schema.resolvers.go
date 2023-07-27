package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"strings"

	"github.com/danny-yamamoto/go-graphql-federation-example/contact/graph/model"
	"github.com/danny-yamamoto/go-graphql-federation-example/contact/internal"
	"github.com/vektah/gqlparser/v2/formatter"
)

// CreateContact is the resolver for the createContact field.
func (r *mutationResolver) CreateContact(ctx context.Context, input model.NewContact) (*model.Contact, error) {
	//panic(fmt.Errorf("not implemented: CreateContact - createContact"))
	return &model.Contact{Firstname: "daisuke"}, nil
}

// Service is the resolver for the service field.
func (r *queryResolver) Service(ctx context.Context) (*model.Service, error) {
	//panic(fmt.Errorf("not implemented: Service - service"))
	s := new(strings.Builder)
	f := formatter.NewFormatter(s)
	f.FormatSchema(r.Schema)
	service := model.Service{
		Name:    "user-service",
		Version: "0.1.0",
		Schema:  s.String(),
	}
	return &service, nil
}

// Contacts is the resolver for the contacts field.
func (r *queryResolver) Contacts(ctx context.Context) ([]*model.Contact, error) {
	//panic(fmt.Errorf("not implemented: Contacts - contacts"))
	return []*model.Contact{{Firstname: "mirai"}}, nil
}

// Mutation returns internal.MutationResolver implementation.
func (r *Resolver) Mutation() internal.MutationResolver { return &mutationResolver{r} }

// Query returns internal.QueryResolver implementation.
func (r *Resolver) Query() internal.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }