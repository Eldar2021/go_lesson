package models

import (
	"time"

	"gorm.io/gorm"
)

/// Advanced
/*
Field-Level Permission
Exported fields have all permissions when doing CRUD with GORM,
and GORM allows you to change the field-level permission with tag,
so you can make a field to be read-only, write-only, create-only,
update-only or ignored
*/
// Field-Level Permission
type User3 struct {
	// Name        string `gorm: "<-:create"` // allow read and create
	// Age         int    `gorm: "<-:update"` // allow read and update
	Email       string `gorm:"<-"`       // allow read and write (create and update)
	Description string `gorm:"<-:false"` // allow read, disable write permission
	// ----------------------------------------------------------------
	LastName   string `gorm:"->"`                 // read only (disable write permission unless it configured)
	DadName    string `gorm:"->;<-:create"`       // allow read and create
	MomName    string `gorm:"->:false;<-:create"` // create only (disabled read from db)
	ChildName  string `gorm:"-"`                  // ignore this field when write and read with struct
	WifeName   string `gorm:"-:all"`              // ignore this field when write, read and migrate with struct
	FamilyName string `gorm:"-:migration"`        // ignore this field when migrate with struct
}

/*
Creating/Updating Time/Unix (Milli/Nano) Seconds Tracking
GORM use CreatedAt, UpdatedAt to track creating/updating time by convention,
and GORM will set the current time when creating/updating if the fields are defined

To use fields with a different name, you can configure those fields
with tag autoCreateTime, autoUpdateTime

If you prefer to save UNIX (milli/nano) seconds instead of time,
you can simply change the field’s data type from time.Time to int
*/
// Creating/Updating Time/Unix (Milli/Nano) Seconds Tracking
type User4 struct {
	CreatedAt    time.Time // Set to current time if it is zero on creating
	UpdatedAt    int       // Set to current unix seconds on updating or if it is zero on creating
	CreatedNano  int64     `gorm:"autoUpdateTime:nano"`  // Use unix nano seconds as updating time
	UpdatedMilli int64     `gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
	Create       int64     `gorm:"autoCreateTime"`       // Use unix seconds as creating time
}

/*
Embedded Struct
For anonymous fields, GORM will include its fields into
its parent struct, for example:
*/
// Embedded Struct
type User5 struct {
	gorm.Model
	Name string
}

// equals
type User6 struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

// For a normal struct field, you can embed it with the tag embedded, for example:
type Author struct {
	Name  string
	Email string
}

type Post1 struct {
	ID     int
	Author Author `gorm:"embedded"`
}

// equals
type Post2 struct {
	ID    int
	Name  string
	Email string
}
