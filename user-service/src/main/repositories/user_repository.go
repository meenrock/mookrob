package repositories

import (
	"database/sql"

	"github.com/mookrob/serviceuser/main/models"

	"github.com/google/uuid"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserById(id uuid.UUID) (models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT "+
		"id, "+
		"status, "+
		"first_name, "+
		"last_name, "+
		"nick_name, "+
		"phone_number, "+
		"email, "+
		"gender, "+
		"age, "+
		"height, "+
		"weight, "+
		"expected_bmi, "+
		"created_at, "+
		"updated_at "+
		"FROM users "+
		"WHERE id = $1", id).Scan(&user.Id, &user.Status, &user.FirstName, &user.LastName, &user.NickName,
		&user.PhoneNumber, &user.Email, &user.Gender, &user.Age, &user.Height, &user.Weight, &user.ExpectedBmi,
		&user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) EditUserById(id uuid.UUID) (models.User, error) {
	var user models.User
	query := "UPDATE users SET " +
		"status = $2, " +
		"first_name = $3, " +
		"last_name = $4, " +
		"nick_name = $5, " +
		"phone_number = $6, " +
		"email = $7, " +
		"gender = $8, " +
		"age = $9, " +
		"height = $10, " +
		"weight = $11, " +
		"expected_bmi = $12, " +
		"updated_at = NOW() " +
		"WHERE id = $1"

	// Execute the update query.
	_, err := r.DB.Exec(query, id, user.Status, user.FirstName, user.LastName, user.NickName,
		user.PhoneNumber, user.Email, user.Gender, user.Age, user.Height, user.Weight,
		user.ExpectedBmi)

	if err != nil {
		return models.User{}, err
	}

	errSelect := r.DB.QueryRow(query, id, user.Status, user.FirstName, user.LastName, user.NickName,
		user.PhoneNumber, user.Email, user.Gender, user.Age, user.Height, user.Weight,
		user.ExpectedBmi).Scan(&user.Id, &user.Status, &user.FirstName, &user.LastName, &user.NickName,
		&user.PhoneNumber, &user.Email, &user.Gender, &user.Age, &user.Height,
		&user.Weight, &user.ExpectedBmi, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return models.User{}, errSelect
	}

	return user, nil
}
