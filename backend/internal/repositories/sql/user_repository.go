package sqlrepo

import (
	"database/sql"

	"forum-backend/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	const query = `
		INSERT INTO users (username)
		VALUES (?)
	`

	result, err := r.db.Exec(query, user.Username)
	if err != nil {
		return err
	}

	// Get auto-generated ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)

	const selectQuery = `
		SELECT created_at
		FROM users
		WHERE id = ?
	`

	err = r.db.QueryRow(selectQuery, user.ID).Scan(&user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	const query = `
		SELECT id, username, created_at
		FROM users
		WHERE id = ?
	`

	row := r.db.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	const query = `
		SELECT id, username, created_at
		FROM users
		WHERE username = ?
	`

	row := r.db.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
