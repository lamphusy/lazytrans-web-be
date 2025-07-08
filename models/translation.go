package models

import "time"

type TranslationResult struct {
    ID            int       `json:"id"`
    FileName      string    `json:"file_name"`
    PageIndex     int       `json:"page_index"`
    TranslatedText string   `json:"translated_text"`
    CreatedAt     time.Time `json:"created_at"`
}