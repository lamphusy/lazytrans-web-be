package repository

import (
	"database/sql"
	"github.com/lamphusy/lazytrans-web-be/models"
)

type UserRepository struct {
	DB *sql.DB
}

// Tạo user mới
func (r *UserRepository) CreateUser(apiKey string) (*models.UserSetting, error) {
	res, err := r.DB.Exec("INSERT INTO user_settings (api_key, current_flash_request_count) VALUES (?, 0)", apiKey)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	return &models.UserSetting{ID: string(id), ApiKey: apiKey, CurrentFlashRequestCount: 0}, nil
}

// Cập nhật API key
func (r *UserRepository) UpdateApiKey(apiKey string) (*models.UserSetting, error) {
	_, err := r.DB.Exec("UPDATE user_settings SET api_key = ?, current_flash_request_count = current_flash_request_count + 1 WHERE id = 1")
	if err != nil {
		return nil, err
	}
	return r.GetUser()
}

// Lấy user hiện tại
func (r *UserRepository) GetUser() (*models.UserSetting, error) {
	row := r.DB.QueryRow("SELECT id, api_key, current_flash_request_count FROM user_settings WHERE id = 1")
	var user models.UserSetting
	err := row.Scan(&user.ID, &user.ApiKey, &user.CurrentFlashRequestCount)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
