package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:user"`
	ID            int64  `bun:",pk,autoincrement" json:"id"`
	Name          string `bun:",notnull" json:"name"`
	Email         string `bun:",unique,notnull" json:"email"`
	Password      string `bun:",notnull" json:"password"`
}
