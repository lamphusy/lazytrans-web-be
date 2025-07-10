package models

import "time"

type TranslationHistoryItem struct {
	ID                string    `json:"id"`
	FileName          string    `json:"file_name"`
	PageRangeLabel    string    `json:"page_range_label"`
	LanguagePairLabel string    `json:"language_pair_label"`
	LastUpdatedTime   time.Time `json:"last_updated_time"`
}

type TranslationHistory struct {
	ID                 string    `json:"id"`
	FilePath           string    `json:"file_path"`
	FromPageIndex      int       `json:"from_page_index"`
	ToPageIndex        int       `json:"to_page_index"`
	SourceLanguage     string    `json:"source_language"`
	TargetLanguage     string    `json:"target_language"`
	LastUpdatedTime    time.Time `json:"last_updated_time"`
	SHA256             string    `json:"sha256"`
	DeepAnalysisResult string    `json:"deep_analysis_result"`
	DeepResearchResult string    `json:"deep_research_result"`
}
