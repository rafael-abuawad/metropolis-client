package api_controllers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rafael-abuawad/metropolis/models"
	"github.com/uptrace/bun"
)

type JobPostApiController struct {
	db *bun.DB
}

func NewJobPostApiController(db *bun.DB) *JobPostApiController {
	return &JobPostApiController{
		db: db,
	}
}

func (jc *JobPostApiController) GetAllJobPosts(c *fiber.Ctx) error {
	var jobPosts []models.JobPost
	if err := jc.db.NewSelect().Model(&jobPosts).Relation("Tags").Scan(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusNotFound, "job posts not found")
	}

	return c.JSON(jobPosts)
}

func (jc *JobPostApiController) CreateJobPost(c *fiber.Ctx) error {
	var jobPost models.JobPost

	if err := c.BodyParser(&jobPost); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not create job post")
	}

	if _, err := jc.db.NewInsert().Model(&jobPost).Exec(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not create job post")
	}

	c.Status(http.StatusCreated)
	return c.JSON(jobPost)
}

func (jc *JobPostApiController) GetJobPostByID(c *fiber.Ctx) error {
	jobPostID, err := c.ParamsInt("id")
	if err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusNotFound, "job post not foudn")
	}

	var jobPost models.JobPost
	jobPost.ID = int64(jobPostID)

	if err := jc.db.NewSelect().Model(&jobPost).WherePK().Relation("Tags").Scan(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusNotFound, "job post not foudn")
	}

	return c.JSON(jobPost)
}

func (jc *JobPostApiController) UpdateJobPost(c *fiber.Ctx) error {
	jobPostID, err := c.ParamsInt("id")
	if err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not update job post")
	}

	var jobPost models.JobPost
	jobPost.ID = int64(jobPostID)
	jobPost.UpdatedAt = time.Now()

	if err := c.BodyParser(&jobPost); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not update job post")
	}

	if _, err := jc.db.NewUpdate().Model(&jobPost).Column("title", "image_url", "description_abbr", "description", "location", "is_remote", "salary_min", "salary_max", "job_id", "updated_at").WherePK().Exec(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not update job post")
	}

	return c.JSON(jobPost)
}

func (jc *JobPostApiController) DeleteJobPost(c *fiber.Ctx) error {
	jobPostID, err := c.ParamsInt("id")
	if err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not delete job post")
	}

	var jobPost models.JobPost
	jobPost.ID = int64(jobPostID)

	if _, err := jc.db.NewDelete().Model(&jobPost).WherePK().Exec(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not delete job post")
	}

	return c.JSON(nil)
}
