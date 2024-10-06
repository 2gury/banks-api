package schema

import "time"

type Bank struct {
	ID               int64     `db:"id"`
	ExternalID       int64     `db:"external_id"`
	ExternalLegacyID int64     `db:"external_legacy_id"`
	Name             string    `db:"name"`
	Logo             string    `db:"logo"`
	URL              string    `db:"url"`
	BankData         string    `db:"bank_data"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}
