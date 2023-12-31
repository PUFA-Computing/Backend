package app

import (
	"Backend/internal/database"
	"Backend/internal/models"
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"log"
)

func TableHasRows(tableName string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM " + tableName + ")"
	var exists bool
	err := database.DB.QueryRow(context.Background(), query).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (id, username, password, first_name, middle_name, last_name, email, student_id, major, role_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	log.Printf("SQL Query: %v", query)

	_, err := database.DB.Exec(context.Background(), query, user.ID, user.Username, user.Password, user.FirstName, user.MiddleName, user.LastName, user.Email, user.StudentID, user.Major, user.RoleID)

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return err
}

func GetUserByUsernameOrEmail(username, email string) (*models.User, error) {
	query := "SELECT * FROM users WHERE username = $1 OR email = $2"
	var user models.User
	err := database.DB.QueryRow(context.Background(), query, username, email).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.StudentID, &user.Major, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found, return nil
		}
		return nil, err
	}
	return &user, nil
}

func IsUsernameExists(username string) (bool, error) {
	log.Printf("username: %v", username)
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)"
	log.Printf("query: %v", query)
	var exists bool
	err := database.DB.QueryRow(context.Background(), query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	log.Printf("exists: %v", exists)

	return exists, nil
}

func IsEmailExists(email string) (bool, error) {
	log.Printf("email: %v", email)
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
	log.Printf("query: %v", query)
	var exists bool
	err := database.DB.QueryRow(context.Background(), query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	log.Printf("exists: %v", exists)

	return exists, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT * FROM users WHERE username = $1"
	var user models.User
	var userID string
	err := database.DB.QueryRow(context.Background(), query, username).Scan(&userID, &user.Username, &user.Password, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.StudentID, &user.Major, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found, return nil
		}
		return nil, err // Return the actual error here
	}
	user.ID, err = uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	var user models.User
	var userID string
	err := database.DB.QueryRow(context.Background(), query, email).Scan(&userID, &user.Username, &user.Password, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.StudentID, &user.Major, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found, return nil
		}
		return nil, err
	}
	user.ID, err = uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	err := database.DB.QueryRow(context.Background(), "SELECT * FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.StudentID, &user.Major, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(UserID uuid.UUID, updatedUser *models.User) error {
	log.Printf("updatedUser: %v", updatedUser)
	_, err := database.DB.Exec(context.Background(), "UPDATE users SET username = $1, password = $2, first_name = $3, middle_name = $4, last_name = $5, email = $6, student_id = $7, major = $8, role_id = $9 WHERE id = $10",
		updatedUser.Username, updatedUser.Password, updatedUser.FirstName, updatedUser.MiddleName, updatedUser.LastName, updatedUser.Email, updatedUser.StudentID, updatedUser.Major, updatedUser.RoleID, UserID)
	log.Printf("err: %v", err)
	return err
}

func DeleteUser(userID uuid.UUID) error {
	_, err := database.DB.Exec(context.Background(), "DELETE FROM users WHERE id = $1", userID)
	return err
}

func ListUsers() ([]*models.User, error) {
	var users []*models.User
	rows, err := database.DB.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.StudentID, &user.Major, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
