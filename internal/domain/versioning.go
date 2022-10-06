package domain

import (
	"database/sql"
	"time"
)

type Versioning struct {
	Version   int32          `db:"version"`
	CreatedAt time.Time      `db:"created_at"`
	CreatedBy string         `db:"created_by"`
	UpdatedAt time.Time      `db:"updated_at"`
	UpdatedBy sql.NullString `db:"updated_by"`
	DeletedAt sql.NullTime   `db:"deleted_at"`
	DeletedBy sql.NullString `db:"deleted_by"`
}