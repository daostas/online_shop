package models

//go:generate reform

// ParametrsLocalization represents a row in parametrs_localization table.
//
//reform:parametrs_localization
type ParametrsLocalization struct {
	ParametrID *int32 `reform:"parametr_id"`
	LangID     *int32 `reform:"lang_id"`
	Title      string `reform:"title"`
}
