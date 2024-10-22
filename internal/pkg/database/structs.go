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

const (
	reviewsTableName = "reviews"

	reviewIDColumn         = "id"
	reviewContentColumn    = "content"
	reviewIsApprovedColumn = "is_approved"
	reviewUserEmailColumn  = "user_email"
	reviewUserPhoneColumn  = "user_phone"
	reviewRatingColumn     = "rating"
	reviewBankIDColumn     = "bank_id"
	reviewUserNameColumn   = "user_name"
	reviewDateColumn       = "date"
)

const (
	settingsTableName = "settings"

	settingsIDColumn                     = "id"
	settingsAutomoderationStrategyColumn = "automoderation_strategy"
)
