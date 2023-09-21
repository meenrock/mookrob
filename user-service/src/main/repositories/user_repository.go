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
