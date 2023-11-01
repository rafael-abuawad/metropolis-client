package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Company struct {
	bun.BaseModel `bun:"table:companies,alias:c"`

	ID           int64     `bun:"id,pk,autoincrement" json:"id"`
	Name         string    `bun:",notnull" json:"name"`
	TwitterURL   string    `bun:"twitter_url" json:"twitter_url"`
	FacebookURL  string    `bun:"facebook_url" json:"facebook_url"`
	InstagramURL string    `bun:"instagram_url" json:"instagram_url"`
	WebsiteURL   string    `bun:"website_url" json:"website_url"`
	JobPosts     []JobPost `bun:"m2m:company_to_job_posts,join:Company=JobPost" json:"job_posts"`
	CreatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

type CompanyToJobPost struct {
	CompanyID int64    `bun:",pk"`
	Company   *Company `bun:"rel:belongs-to,join:company_id=id"`
	JobPostID int64    `bun:",pk"`
	JobPost   *JobPost `bun:"rel:belongs-to,join:job_post_id=id"`
}
