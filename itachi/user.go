package itachi

import (
	"context"
	"time"
)

// User represents a user in the system
type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        string
	Email     string
	Wallet    string
}

// UserService represents the data layer to manage users
type UserService interface {
	FindUserByID(ctx context.Context, id string) (*User, error)
	FindUsers(ctx context.Context, filter UserFilter) ([]*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, id string, update UserUpdate) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}

type UserFilter struct {
	ID     *string
	Email  *string
	Wallet *string

	Offset int
	Limit  int
}

type UserUpdate struct {
	Email  *string
	Wallet *string
}
