package models

//go:generate reform

// GroupsLocalization represents a row in groups_localization table.
//
//reform:groups_localization
type GroupsLocalization struct {
	GroupID     *int32  `reform:"group_id"`
	LangID      *int32  `reform:"lang_id"`
	Title       string  `reform:"title"`
	Description *string `reform:"description"`
}
