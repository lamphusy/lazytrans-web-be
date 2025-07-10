package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lamphusy/lazytrans-web-be/models"
)

type TranslationHistoryRepository struct {
	DB *sql.DB
}

// Lấy tất cả tên file đã dịch (không trùng lặp)
func (r *TranslationHistoryRepository) GetAllFileNames() ([]string, error) {
	rows, err := r.DB.Query("SELECT DISTINCT file_path FROM translation_histories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var files []string
	for rows.Next() {
		var file string
		if err := rows.Scan(&file); err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

// Lấy lịch sử dịch theo tên file (tìm kiếm gần đúng)
func (r *TranslationHistoryRepository) GetTranslationHistoryItemsByName(name string) ([]models.TranslationHistoryItem, error) {
	rows, err := r.DB.Query(`
        SELECT id, file_path, from_page_index, to_page_index, source_language, target_language, last_updated_time
        FROM translation_histories
        WHERE LOWER(file_path) = LOWER(?) OR LOWER(file_path) LIKE LOWER(?) OR LOWER(file_path) LIKE LOWER(?) 
        ORDER BY last_updated_time DESC
    `, name, name+"%", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.TranslationHistoryItem
	for rows.Next() {
		var id, filePath, sourceLang, targetLang string
		var fromPage, toPage int
		var lastUpdated time.Time
		if err := rows.Scan(&id, &filePath, &fromPage, &toPage, &sourceLang, &targetLang, &lastUpdated); err != nil {
			return nil, err
		}
		items = append(items, models.TranslationHistoryItem{
			ID:                id,
			FileName:          filePath,
			PageRangeLabel:    "Page " + itoa(fromPage+1) + " → Page " + itoa(toPage+1),
			LanguagePairLabel: sourceLang + " → " + targetLang,
			LastUpdatedTime:   lastUpdated,
		})
	}
	return items, nil
}

// Lấy 10 lịch sử dịch mới nhất
func (r *TranslationHistoryRepository) GetLatestTranslationHistoryItems() ([]models.TranslationHistoryItem, error) {
	rows, err := r.DB.Query(`
        SELECT id, file_path, from_page_index, to_page_index, source_language, target_language, last_updated_time
        FROM translation_histories
        ORDER BY last_updated_time DESC
        LIMIT 10
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.TranslationHistoryItem
	for rows.Next() {
		var id, filePath, sourceLang, targetLang string
		var fromPage, toPage int
		var lastUpdated time.Time
		if err := rows.Scan(&id, &filePath, &fromPage, &toPage, &sourceLang, &targetLang, &lastUpdated); err != nil {
			return nil, err
		}
		items = append(items, models.TranslationHistoryItem{
			ID:                id,
			FileName:          filePath,
			PageRangeLabel:    "Page " + itoa(fromPage+1) + " → Page " + itoa(toPage+1),
			LanguagePairLabel: sourceLang + " → " + targetLang,
			LastUpdatedTime:   lastUpdated,
		})
	}
	return items, nil
}

// Lấy lịch sử dịch theo ID
func (r *TranslationHistoryRepository) GetTranslationHistoryById(id string) (*models.TranslationHistory, error) {
	row := r.DB.QueryRow(`
        SELECT id, file_path, from_page_index, to_page_index, source_language, target_language, last_updated_time, sha256, deep_analysis_result, deep_research_result
        FROM translation_histories WHERE id = ?
    `, id)
	var h models.TranslationHistory
	var lastUpdated time.Time
	err := row.Scan(&h.ID, &h.FilePath, &h.FromPageIndex, &h.ToPageIndex, &h.SourceLanguage, &h.TargetLanguage, &lastUpdated, &h.SHA256, &h.DeepAnalysisResult, &h.DeepResearchResult)
	if err != nil {
		return nil, err
	}
	h.LastUpdatedTime = lastUpdated
	return &h, nil
}

// Thống kê số tài liệu đã dịch (theo SHA256)
func (r *TranslationHistoryRepository) GetTranslatedDocumentsCount() (int, error) {
	row := r.DB.QueryRow("SELECT COUNT(DISTINCT sha256) FROM translation_histories")
	var count int
	err := row.Scan(&count)
	return count, err
}

// Thống kê số trang đã dịch
func (r *TranslationHistoryRepository) GetTranslatedPagesCount() (int, error) {
	row := r.DB.QueryRow("SELECT SUM(to_page_index - from_page_index + 1) FROM translation_histories")
	var count int
	err := row.Scan(&count)
	return count, err
}

// Thống kê số trang dịch hôm nay
func (r *TranslationHistoryRepository) GetTranslatedPageCountToday() (int, error) {
	middayYesterday := time.Now().Truncate(24 * time.Hour).Add(-12 * time.Hour)
	row := r.DB.QueryRow("SELECT SUM(to_page_index - from_page_index + 1) FROM translation_histories WHERE last_updated_time >= ?", middayYesterday)
	var count int
	err := row.Scan(&count)
	return count, err
}

// Thống kê số reasoning request hôm nay
func (r *TranslationHistoryRepository) GetReasoningRequestCountToday() (int, error) {
	middayYesterday := time.Now().Truncate(24 * time.Hour).Add(-12 * time.Hour)
	row := r.DB.QueryRow(`
        SELECT COUNT(*) FROM translation_histories
        WHERE last_updated_time >= ?
        AND (deep_analysis_result IS NOT NULL OR deep_research_result IS NOT NULL)
    `, middayYesterday)
	var count int
	err := row.Scan(&count)
	return count, err
}

// Helper chuyển int sang string
func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}
