package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/mookrob/serviceauth/main/enums"
	"github.com/mookrob/serviceauth/main/models"
	"github.com/mookrob/shared/constants"
)

type AuthenticationRepository struct {
	DB *sql.DB
}

func NewAuthenticationRepository(db *sql.DB) *AuthenticationRepository {
	return &AuthenticationRepository{DB: db}
}

func (r *AuthenticationRepository) CreateAuthenticationUser(authData models.AuthenticationData) (*uuid.UUID, error) {
	var id uuid.UUID
	err := r.DB.QueryRow("INSERT INTO authentication_data ("+
		"status, "+
		"username, "+
		"password, "+
		"user_id, "+
		"role, "+
		"created_at, "+
		"updated_at "+
		") VALUES ($1, $2, $3, $4, $5, now(), now()) RETURNING id", enums.ACTIVE, authData.Username, authData.Password, authData.UserId,
		constants.GENERAL_USER).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *AuthenticationRepository) GetAuthenticationByUsernameAndStatusActive(username string) (models.AuthenticationData, error) {
	var authenticationData models.AuthenticationData
	err := r.DB.QueryRow("SELECT "+
		"ad.id, "+
		"ad.username, "+
		"ad.password, "+
		"ad.user_id, "+
		"ad.role, "+
		"ad.status, "+
		"ad.created_at, "+
		"ad.updated_at "+
		"FROM authentication_data ad "+
		"WHERE ad.username = $1 and ad.status = 'ACTIVE'", username).Scan(&authenticationData.Id, &authenticationData.Username, &authenticationData.Password,
		&authenticationData.UserId, &authenticationData.Role, &authenticationData.Status, &authenticationData.CreatedAt, &authenticationData.UpdatedAt)

	if err != nil {
		return models.AuthenticationData{}, err
	}

	return authenticationData, nil
}
