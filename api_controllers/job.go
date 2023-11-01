package api_controllers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rafael-abuawad/metropolis/models"
	"github.com/uptrace/bun"
)

type JobApiController struct {
	db *bun.DB
}

func NewJobApiController(db *bun.DB) *JobApiController {
	return &JobApiController{
		db: db,
	}
}

func (jc *JobApiController) GetAllJobs(c *fiber.Ctx) error {
	var jobs []models.Job

	if err := jc.db.NewSelect().Model(&jobs).Relation("JobPost.Tags").Scan(c.Context()); err != nil {
		return fiber.NewError(http.StatusBadRequest, "jobs not found")
	}

	return c.JSON(jobs)
}

func (jc *JobApiController) CreateJob(c *fiber.Ctx) error {
	var job models.Job

	if err := c.BodyParser(&job); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not create job")
	}

	if _, err := jc.db.NewInsert().Model(&job).Exec(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not create job")
	}

	c.Status(http.StatusCreated)
	return c.JSON(job)
}

func (jc *JobApiController) GetJobByID(c *fiber.Ctx) error {
	jobID, err := c.ParamsInt("id")
	if err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusNotFound, "job not found")
	}

	var job models.Job
	job.ID = int64(jobID)

	if err := jc.db.NewSelect().Model(&job).Relation("JobPost.Tags").WherePK().Scan(c.Context()); err != nil {
		return fiber.NewError(http.StatusNotFound, "job not found")
	}

	return c.JSON(job)
}

func (jc *JobApiController) UpdateJob(c *fiber.Ctx) error {
	jobID, err := c.ParamsInt("id")
	if err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not update job")
	}

	var job models.Job
	job.ID = int64(jobID)
	job.UpdatedAt = time.Now()

	if err := c.BodyParser(&job); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not update job")
	}

	if _, err := jc.db.NewUpdate().Model(&job).Column("content", "apply_url", "apply_email", "updated_at").WherePK().Exec(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not update job")
	}

	return c.JSON(job)
}

func (jc *JobApiController) DeleteJob(c *fiber.Ctx) error {
	jobID, err := c.ParamsInt("id")
	if err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not delete job")
	}

	var job models.Job
	job.ID = int64(jobID)

	if _, err := jc.db.NewDelete().Model(&job).WherePK().Exec(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not delete job")
	}

	return c.JSON(nil)
}
