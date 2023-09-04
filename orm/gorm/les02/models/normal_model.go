package models

import "time"

/*
Models are normal structs with basic Go types, pointers/alias of
them or custom types implementing Scanner and Valuer interfaces
*/

type User struct {
	ID       int
	name     string
	Email    *string
	Username string
	Password string
	Birthday *time.Time
}

/*
GORM prefers convention over configuration. By default, GORM uses
ID as primary key, pluralizes struct name to snake_cases as table name,
snake_case as column name, and uses CreatedAt, UpdatedAt to track
`creating/updating` time

If you follow the conventions adopted by GORM, you’ll need to write very
little configuration/code. If convention doesn’t match your requirements,
GORM allows you to configure them
*/
