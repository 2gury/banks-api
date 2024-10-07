package schema

import "time"

// Review struct
type Review struct {
	ID         int64     `db:"id"`
	ProductID  int64     `db:"product_id"`
	UserID     int64     `db:"user_id"`
	Content    string    `db:"content"`
	Rating     int       `db:"rating"`
	CreatedAt  time.Time `db:"created_at`
	Approved   bool      `db:"approved"`
	ReviewData string    `db:"review_data"`
}
