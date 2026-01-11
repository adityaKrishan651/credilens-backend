package models

type AnalyzeRequest struct {
	Content string `json:"content"`
}

type AnalyzeResponse struct {
	RiskLevel string   `json:"risk_level"`
	Summary   string   `json:"summary"`
	Signals   []string `json:"signals"`
}
