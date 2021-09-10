package model

import (
	"time"

	_ "gorm.io/driver/sqlite"
)

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}
