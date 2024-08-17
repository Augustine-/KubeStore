package repository

import (
	"database/sql"
	"user/internal/models"
)

// interfaces are contracts
// they promise functionality for any type that implements them
// or in other words, they are the way you 'interface' with the contents.
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id int64) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type PostgresUserRepository struct {
	DB *sql.DB
}

func (r *PostgresUserRepository) CreateUser(user *models.User) error {
	_, err := r.DB.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	return err
}

func (r *PostgresUserRepository) GetUserByID(id int64) (*models.User, error) {
	user := &models.User{}
	err := r.DB.QueryRow("SELECT id, username, email, password FROM users WHERE id=$1", id).Scan(&user.ID, &user.Username, &user.Email)
	return user, err
}

func (r *PostgresUserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := r.DB.QueryRow("SELECT id, username, email, password FROM users WHERE email=$1", email).Scan(&user.ID, &user.Username, &user.Email)
	return user, err
}
