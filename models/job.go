package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Job struct {
	bun.BaseModel `bun:"table:jobs,alias:j"`

	ID         int64     `bun:"id,pk,autoincrement" json:"id"`
	Content    string    `bun:",notnull" json:"content"`
	ApplyURL   string    `json:"apply_url"`
	ApplyEmail string    `json:"apply_email"`
	JobPost    *JobPost  `bun:"rel:has-one,join:id=job_id" json:"job_post"`
	CreatedAt  time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}
