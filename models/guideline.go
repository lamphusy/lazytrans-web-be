package models

type TranslationGuideline struct {
    Language                        string                        `json:"language"`
    DocumentAnalysisSynthesis       DocumentAnalysisSynthesis     `json:"document_analysis_synthesis"`
    RecommendedTranslationStrategy  RecommendedTranslationStrategy `json:"recommended_translation_strategy"`
    DetailedImplementationGuideline string                        `json:"detailed_implementation_guideline"`
    PotentialChallengesAndRiskAssessment []string                 `json:"potential_challenges_and_risk_assessment"`
}

type DocumentAnalysisSynthesis struct {
    InferredPurposeAndAudience      string `json:"inferred_purpose_and_audience"`
    CoreToneAndStyleCharacterization string `json:"core_tone_and_style_characterization"`
    ContentComplexityAndNature      string `json:"content_complexity_and_nature"`
    CulturalEmbeddingLevel          string `json:"cultural_embedding_level"`
    StructuralSignificance          string `json:"structural_significance"`
}

type RecommendedTranslationStrategy struct {
    PrimaryTheoreticalOrientation string   `json:"primary_theoretical_orientation"`
    StrategicJustification        string   `json:"strategic_justification"`
    KeyStrategicPillars           []string `json:"key_strategic_pillars"`
    AdaptationNotes               string   `json:"adaptation_notes"`
}