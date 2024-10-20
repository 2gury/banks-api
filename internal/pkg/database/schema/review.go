package schema

type Review struct {
	ID         int64  `db:"id"`
	Content    string `db:"content"`
	IsApproved bool   `db:"is_approved"`
	UserEmail  string `db:"user_email"`
	UserPhone  string `db:"user_phone"`
	Rating     int64  `db:"rating"`
	BankID     int64  `db:"bank_id"`
}
