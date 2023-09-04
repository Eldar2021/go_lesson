package models

import (
	"time"

	"gorm.io/gorm"
)

/// gorm.Model
/*
GORM defined a gorm.Model struct, which includes fields
`ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
*/

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

// User1 = User2
type User1 struct {
	gorm.Model
	Name string
}

// equals
type User2 struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

/*
You can embed it into your struct to include those fields, refer Embedded Struct
*/
