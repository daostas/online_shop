package models

//go:generate reform

// Languages represents a row in languages table.
//
//reform:languages
type Languages struct {
	LangID    int32   `reform:"lang_id,pk"`
	Code      string  `reform:"code"`
	Image     *string `reform:"image"`
	Locale    string  `reform:"locale"`
	LangName  string  `reform:"lang_name"`
	SortOrder int32   `reform:"sort_order"`
	Status    bool    `reform:"status"`
}
