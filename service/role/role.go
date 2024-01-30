package role

import (
	"context"
)

type IRoleService interface {
	Get(ctx context.Context, id string) (*Role, error)
	List(ctx context.Context) ([]*Role, error)
	Save(ctx context.Context, id string, name string) error
	Create(ctx context.Context, id string, name string) error
	Delete(ctx context.Context, id string) error
}
