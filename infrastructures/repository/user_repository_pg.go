package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/generator"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
	"strings"
)

type UserRepositoryPG struct /* implements UserRepository */ {
	db          *sql.DB
	idGenerator generator.IdGenerator
}

func NewUserRepositoryPG(db *sql.DB, idGenerator generator.IdGenerator) repository.UserRepository {
	return &UserRepositoryPG{
		db:          db,
		idGenerator: idGenerator,
	}
}

func (r *UserRepositoryPG) AddUser(payload *entity.RegisterUserPayload) string {
	// Create ID
	id := r.idGenerator.Generate()

	// Query
	query := `INSERT INTO 
    			users(id, username, password, email) 
			  VALUES
			      ($1, $2, $3, $4)
			  RETURNING id`

	var returnedId string
	err := r.db.QueryRow(
		query,
		id,
		payload.Username,
		payload.Password,
		payload.Email,
	).Scan(&returnedId)

	// Evaluate
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			panic(fiber.NewError(fiber.StatusConflict, "Username or Email already exists!"))
		}

		panic(fmt.Errorf("user_repo_pg_error: add user: %v", err))
	}

	return returnedId
}

func (r *UserRepositoryPG) GetUserForLogin(identity string) (*entity.User, string) {
	var userToken entity.User
	var encryptedPassword string

	// Query
	query := `
		SELECT 
		    id, 
		    username,
		    email,
		    avatar_link,
		    is_verified,
		    role,
		    password 
		FROM users 
		WHERE email = $1 OR username = $1`
	err := r.db.QueryRow(query, identity).Scan(
		&userToken.Id,
		&userToken.Username,
		&userToken.Email,
		&userToken.AvatarLink,
		&userToken.IsVerified,
		&userToken.Role,
		&encryptedPassword,
	)

	// Evaluate
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(fiber.NewError(fiber.StatusNotFound, "User not found!"))
		} else {
			panic(fmt.Errorf("user_repo_pg_error: get userId by identity %v", err))
		}
	}

	return &userToken, encryptedPassword
}

func (r *UserRepositoryPG) GetUserById(id string) *entity.User {
	var result entity.User

	// Query
	query := `SELECT id, username, email, avatar_link FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&result.Id,
		&result.Username,
		&result.Email,
		&result.AvatarLink,
	)

	// Evaluate
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(fiber.NewError(fiber.StatusNotFound, "User not found!"))
		}

		panic(fmt.Errorf("user_repo_pg_error: get userId by id %v", err))
	}

	return &result
}

func (r *UserRepositoryPG) UpdateUserById(id string, payload *entity.UpdateUserPayload, newAvatarLink string) string {
	// Base query and arguments (Only updating the password if its not empty)
	query := `
		WITH old_data AS (
			SELECT avatar_link
			FROM users
			WHERE id = $1
		)
		UPDATE users 
		SET username = $2, email = $3, avatar_link = $4`

	args := []interface{}{id, payload.Username, payload.Email, newAvatarLink}

	// Conditionally add password update
	if payload.Password != "" {
		query += `, password = $5`
		args = append(args, payload.Password)
	}

	query += `
		FROM old_data
		WHERE users.id = $1
		RETURNING old_data.avatar_link`

	// Execute the query
	var oldAvatarLink string
	err := r.db.QueryRow(query, args...).Scan(&oldAvatarLink)

	// Evaluate
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(fiber.NewError(fiber.StatusNotFound, "User not found!"))
		}
		if strings.Contains(err.Error(), "unique constraint") {
			panic(fiber.NewError(fiber.StatusConflict, "Username or Email already exists!"))
		}
		panic(fmt.Errorf("user_repo_pg_error: update user: %v", err))
	}

	return oldAvatarLink
}

func (r *UserRepositoryPG) ActivateAccount(id string) {
	// Query
	query := `UPDATE users SET is_verified = TRUE WHERE id = $1`
	result, err := r.db.Exec(query, id)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("user_repo_pg_error: update user: %v", err))
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		panic(fiber.NewError(fiber.StatusNotFound, "Unable to activate user!"))
	}
}
