package schema

type Translation struct {
	ID               int64  `db:"id"`
	Lexeme           string `db:"lexeme"`
	TranslatedLexeme string `db:"translated_lexeme"`
	SourceLanguage   string `db:"source_language"`
	TargetLanguage   string `db:"target_language"`
}
