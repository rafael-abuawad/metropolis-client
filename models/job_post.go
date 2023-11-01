package models

import (
	"time"

	"github.com/uptrace/bun"
)

type JobPost struct {
	bun.BaseModel `bun:"table:jobs_posts,alias:jp"`

	ID              int64     `bun:"id,pk,autoincrement" json:"id"`
	Title           string    `bun:",notnull" json:"title"`
	ImageURL        string    `json:"image_url"`
	DescriptionAbbr string    `bun:",notnull" json:"description_abbr"`
	Description     string    `bun:",notnull" json:"description"`
	Location        string    `json:"location"`
	IsRemote        bool      `json:"is_remote"`
	SalaryMin       string    `json:"salary_min"`
	SalaryMax       string    `json:"salary_max"`
	JobID           int64     `json:"job_id"`
	Tags            []Tag     `bun:"m2m:job_post_to_tags,join:JobPost=Tag" json:"tags"`
	CreatedAt       time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt       time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

type JobPostToTag struct {
	JobPostID int64    `bun:",pk"`
	JobPost   *JobPost `bun:"rel:belongs-to,join:job_post_id=id"`
	TagID     int64    `bun:",pk"`
	Tag       *Tag     `bun:"rel:belongs-to,join:tag_id=id"`
}
