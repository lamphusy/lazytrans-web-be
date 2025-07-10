package models

import (
    "time"
)

type UserSetting struct {
    ID                     string     `json:"id"` // UUID, dùng string cho tiện thao tác
    ApiKey                 string     `json:"api_key"`
    CurrentProRequestCount int16      `json:"current_pro_request_count"`   // short → int16
    CurrentFlashRequestCount int      `json:"current_flash_request_count"`
    EnableAutoRetryOnFailed bool      `json:"enable_auto_retry_on_failed"`
    EnableAutoSaveTranslation bool    `json:"enable_auto_save_translation"`
    LastestRequestTime     *time.Time `json:"lastest_request_time,omitempty"` // nullable
    LastResetQuotaTime     time.Time  `json:"last_reset_quota_time"`
}