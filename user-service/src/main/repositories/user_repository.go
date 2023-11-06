package repositories

import (
	"database/sql"

	enums "github.com/mookrob/serviceuser/main/enums"
	"github.com/mookrob/serviceuser/main/models"

	"github.com/google/uuid"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user models.User) (*uuid.UUID, error) {
	var id uuid.UUID
	err := r.DB.QueryRow("INSERT INTO users ("+
		"status, "+
		"first_name, "+
		"last_name, "+
		"nick_name, "+
		"phone_number, "+
		"email,"+
		"gender, "+
		"age, "+
		"height, "+
		"weight, "+
		"expected_bmi, "+
		"created_at, "+
		"updated_at "+
		") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, now(), now()) RETURNING id", enums.ACTIVE, user.FirstName, user.LastName, user.NickName,
		user.PhoneNumber, user.Email, user.Gender, user.Age, user.Height, user.Weight, user.ExpectedBmi).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &id, nil
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

func (r *UserRepository) ExistByEmail(email string) (bool, error) {
	row := r.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email=$1)", email)

	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}
