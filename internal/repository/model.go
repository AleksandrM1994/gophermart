package repository

import "time"

type User struct {
	ID       string  `db:"id;primary_key"`
	Login    string  `db:"login;not null;unique"`
	Password string  `db:"password;not null"`
	Balance  float32 `db:"withdrawal"`
	Session  Session `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}

type Session struct {
	Cookie       string     `db:"cookie;primary_key"`
	CookieFinish *time.Time `db:"cookie_finish;not null"`
	UserID       string     `db:"user_id;not null"`
}

func (Session) TableName() string {
	return "sessions"
}

type Order struct {
	ID         string      `db:"id;primary_key"`
	Accrual    float32     `db:"accrual"`
	Status     OrderStatus `gorm:"type:enum('UNKNOWN', 'REGISTERED', 'INVALID', 'PROCESSING', 'PROCESSED')"`
	UploadedAt *time.Time  `db:"uploaded_at;not null"`
	UserID     string      `db:"user_id;not null"`
}

func (Order) TableName() string {
	return "orders"
}

type Withdrawal struct {
	OrderID     string     `db:"order_id;primary_key"`
	Sum         float32    `db:"sum"`
	ProcessedAt *time.Time `db:"processed_at;not null"`
	UserID      string     `db:"user_id;not null"`
}

func (User) Withdrawal() string {
	return "withdrawals"
}
