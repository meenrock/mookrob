package repositories

import (
	"database/sql"
	"time"

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
		") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id", enums.ACTIVE, user.FirstName, user.LastName, user.NickName,
		user.PhoneNumber, user.Email, user.Gender, user.Age, user.Height, user.Weight, user.ExpectedBmi, time.Now(), time.Now()).Scan(&id)

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

func (r *UserRepository) UpdateUser(user models.User) (*models.User, error) {
	err := r.DB.QueryRow("UPDATE users "+
		"SET "+
		"status = $1, "+
		"first_name = $2, "+
		"last_name = $3, "+
		"nick_name = $4, "+
		"phone_number = $5, "+
		"email = $6,"+
		"gender = $7, "+
		"age = $8, "+
		"height = $9, "+
		"weight = $10, "+
		"expected_bmi = $11, "+
		"created_at = $12, "+
		"updated_at = $13"+
		" WHERE users.id = $14", enums.ACTIVE, user.FirstName, user.LastName, user.NickName,
		user.PhoneNumber, user.Email, user.Gender, user.Age, user.Height, user.Weight, user.ExpectedBmi, time.Now(), time.Now(), user.Id).Scan(user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) ExistByEmail(email string) (bool, error) {
	row := r.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email=$1)", email)

	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *UserRepository) ExistByID(uuid string) (bool, error) {
	row := r.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE id=$1)", uuid)

	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}
