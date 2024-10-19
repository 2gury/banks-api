package schema

// Review struct
type Review struct {
	ID         int64  `db:"id"`
	ProductID  int64  `db:"product_id"`
	Content    string `db:"content"`
	Rating     int    `db:"rating"`
	Approved   bool   `db:"approved"`
	ReviewData string `db:"review_data"`
}
