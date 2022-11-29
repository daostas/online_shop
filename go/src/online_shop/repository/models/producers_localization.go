package models

//go:generate reform

// ProducersLocalization represents a row in producers_localization table.
//
//reform:producers_localization
type ProducersLocalization struct {
	ProducerID  *int32  `reform:"producer_id"`
	LangID      *int32  `reform:"lang_id"`
	Title       string  `reform:"title"`
	Description *string `reform:"description"`
}
