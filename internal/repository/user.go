// TODO: error handling on not found
package repository

// will be unused because of sqlc generated db.Queries already acts as repository/dao
// just used it straight on the service layer.

import (
	"context"

	db "github.com/rslbn/blog/postgres"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*db.User, error)
	GetAll(ctx context.Context) ([]db.User, error)
	Insert(ctx context.Context, args db.InsertUserParams) (*db.User, error)
	// UpdateByUsername(ctx context.Context, user *db.User) (db.User, error)
	// Delete(ctx context.Context, id uint32) (db.User, error)
}

type userRepository struct {
	queries *db.Queries
}

func NewUserRepository(queries *db.Queries) UserRepository {
	return &userRepository{
		queries: queries,
	}
}

func (ur *userRepository) GetAll(ctx context.Context) ([]db.User, error) {
	users, err := ur.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetByUsername(ctx context.Context, username string) (*db.User, error) {
	user, err := ur.queries.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) GetByEmail(ctx context.Context, email string) (*db.User, error) {
	user, err := ur.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) Insert(ctx context.Context, args db.InsertUserParams) (*db.User, error) {
	user, err := ur.queries.InsertUser(ctx, args)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// import (
// 	"context"
// 	"fmt"

// 	"github.com/jackc/pgx/v5"
// 	"github.com/rslbn/blog/internal/model"
// 	db "github.com/rslbn/blog/postgres"
// )

// type UserRepository interface {
// 	GetByID(ctx context.Context, id uint32) (*model.User, error)
// 	GetByUsername(ctx context.Context, username string) (*model.User, error)
// 	GetByEmail(ctx context.Context, email string) (*model.User, error)
// 	Save(ctx context.Context, user *model.User) (*model.User, error)
// }

// type userRepository struct {
// 	db *pgx.Conn
// }

// func NewUserRepository(db *pgx.Conn) UserRepository {
// 	return &userRepository{db}
// }

// func (r *userRepository) GetByID(ctx context.Context, id uint32) (*model.User, error) {
// 	query := `SELECT
// 		user_id, username,
// 		email, password,
// 		created_at, updated_at
// 		FROM users WHERE user_id = $1`
// 	row := r.db.QueryRowContext(ctx, query, id)
// 	var user model.User
// 	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
// 	query := `SELECT
// 		user_id, username,
// 		email, password,
// 		created_at, updated_at
// 		FROM users WHERE username = $1`
// 	row := r.db.QueryRowContext(ctx, query, username)
// 	var user model.User
// 	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
// 	query := `SELECT
// 		user_id, username,
// 		email, password,
// 		created_at, updated_at
// 		FROM users WHERE email = $1`
// 	row := r.db.QueryRowContext(ctx, query, email)
// 	var user model.User
// 	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *userRepository) Save(ctx context.Context, user *model.User) (savedUser *model.User, err error) {
// 	// Begin transaction
// 	tx, err := r.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to begin transaction: %w", err)
// 	}

// 	// defer a rollback
// 	defer func() {
// 		if p := recover(); p != nil {
// 			tx.Rollback()
// 			panic(p)
// 		} else if err != nil {
// 			tx.Rollback()
// 		} else {
// 			err = tx.Commit()
// 		}
// 	}()

// 	// Prepare the statement
// 	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)
// 		RETURNING user_id, username, email, password, created_at, updated_at`
// 	stmt, err := tx.PrepareContext(ctx, query)
// 	// Check if there is an error
// 	if err != nil {
// 		return
// 	}
// 	defer stmt.Close()
// 	savedUser = &model.User{}
// 	// will handle tx.Commit
// 	err = stmt.QueryRowContext(ctx, user.Username, user.Email, user.Password).Scan(
// 		&savedUser.UserID,
// 		&savedUser.Username,
// 		&savedUser.Email,
// 		&savedUser.Password,
// 		&savedUser.CreatedAt,
// 		&savedUser.UpdatedAt,
// 	)
// 	if err != nil {
// 		savedUser = nil
// 		return
// 	}

// 	return
// }
