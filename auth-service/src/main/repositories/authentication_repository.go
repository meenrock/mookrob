package repositories

import (
	"database/sql"

	"github.com/mookrob/serviceauth/main/models"
)

type AuthenticationRepository struct {
	DB *sql.DB
}

func NewAuthenticationRepository(db *sql.DB) *AuthenticationRepository {
	return &AuthenticationRepository{DB: db}
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
