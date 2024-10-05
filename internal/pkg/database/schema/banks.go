package schema

type Bank struct {
	ID       int64  `db:"id"`
	BankName string `db:"bank_name"`
}
