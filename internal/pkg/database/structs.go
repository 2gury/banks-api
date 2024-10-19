package database

const (
	banksTableName = "banks"

	bankIDColumn               = "id"
	bankExternalIDColumn       = "external_id"
	bankExternalLegacyIDColumn = "external_legacy_id"
	bankNameColumn             = "name"
	bankLogoColumn             = "logo"
	bankURLColumn              = "url"
	bankDataColumn             = "bank_data"
	bankCreatedAtColumn        = "created_at"
	bankUpdatedAtColumn        = "updated_at"
)

const (
	translationsTableName = "translations"

	translationIDColumn               = "id"
	translationLexemeColumn           = "lexeme"
	translationTranslatedLexemeColumn = "translated_lexeme"
	translationSourceLanguageColumn   = "source_language"
	translationTargetLanguageColumn   = "target_language"
)
