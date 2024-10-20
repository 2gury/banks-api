package schema

type Settings struct {
	ID                     int64 `db:"id"`
	AutomoderationStrategy bool  `db:"automoderation_strategy"`
}
