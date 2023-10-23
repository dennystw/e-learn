package domain

import (
	"context"
	"time"
)

// Role is representing the role data struct
type Role struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// RoleUsecase represent the role's usecases
type RoleUsecase interface {
	// Fetch(ctx context.Context, cursor string, num int64) ([]Role, string, error)
	// GetByID(ctx context.Context, id int64) (Role, error)
	// Update(ctx context.Context, ar *Role) error
	// GetByTitle(ctx context.Context, title string) (Role, error)
	Store(context.Context, *Role) error
	// Delete(ctx context.Context, id int64) error
}

// ArticleRepository represent the article's repository contract
type RoleRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Role, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (Role, error)
	Update(ctx context.Context, ar *Role) error
	Store(ctx context.Context, a *Role) error
	Delete(ctx context.Context, id int64) error
}
