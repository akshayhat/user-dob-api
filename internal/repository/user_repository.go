package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/akshayhat/user-dob-api/db/sqlc"
)

type UserRepository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

// Constructor
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db:      db,
		queries: sqlc.New(db),
	}
}

// Create User
func (r *UserRepository) CreateUser(
	ctx context.Context,
	name string,
	dob time.Time,
) (sqlc.User, error) {
	return r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

// Get User by ID
func (r *UserRepository) GetUserByID(
	ctx context.Context,
	id int32,
) (sqlc.User, error) {
	return r.queries.GetUserByID(ctx, id)
}

// List All Users
func (r *UserRepository) ListUsers(
	ctx context.Context,
) ([]sqlc.User, error) {
	return r.queries.ListUsers(ctx)
}

// Update User
func (r *UserRepository) UpdateUser(
	ctx context.Context,
	id int32,
	name string,
	dob time.Time,
) (sqlc.User, error) {
	return r.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}

// Delete User
func (r *UserRepository) DeleteUser(
	ctx context.Context,
	id int32,
) error {
	return r.queries.DeleteUser(ctx, id)
}
