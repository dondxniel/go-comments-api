package comment

import (
	"context"
	"fmt"
)

// * Representation of our comment structure
type Comment struct {
	ID string 
	Slug string 
	Body string 
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// * The struct on which all our logic will be built
type Service struct {
	Store Store
}

// * Returns a pointer to a new service;
// * Serves as a constructor for the Service struct
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return cmt, nil
}