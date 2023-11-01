package api_controllers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rafael-abuawad/metropolis/models"
	"github.com/uptrace/bun"
)

type TagApiController struct {
	db *bun.DB
}

func NewTagApiController(db *bun.DB) *TagApiController {
	return &TagApiController{
		db: db,
	}
}

func (tc *TagApiController) GetAllTags(c *fiber.Ctx) error {
	var tags []models.Tag

	if err := tc.db.NewSelect().Model(&tags).Scan(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusNotFound, "tags not found")
	}

	return c.JSON(tags)
}

func (tc *TagApiController) CreateTag(c *fiber.Ctx) error {
	var tag models.Tag

	if err := c.BodyParser(&tag); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not create tag")
	}

	if _, err := tc.db.NewInsert().Model(&tag).Exec(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not create tag")
	}

	c.Status(http.StatusCreated)
	return c.JSON(tag)
}

func (tc *TagApiController) GetTagByID(c *fiber.Ctx) error {
	tagID, err := c.ParamsInt("id")
	if err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusNotFound, "tag not found")
	}

	var tag models.Tag
	tag.ID = int64(tagID)

	if err := tc.db.NewSelect().Model(&tag).WherePK().Scan(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusNotFound, "tag not found")
	}

	return c.JSON(tag)
}

func (tc *TagApiController) UpdateTag(c *fiber.Ctx) error {
	tagID, err := c.ParamsInt("id")
	if err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not update tag")
	}

	var tag models.Tag
	tag.ID = int64(tagID)
	tag.UpdatedAt = time.Now()

	if err := c.BodyParser(&tag); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not update tag")
	}

	if _, err := tc.db.NewUpdate().Model(&tag).Column("name", "updated_at").WherePK().Exec(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not update tag")
	}

	return c.JSON(tag)
}

func (tc *TagApiController) DeleteTag(c *fiber.Ctx) error {
	tagID, err := c.ParamsInt("id")
	if err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not delete tag")
	}

	var tag models.Tag
	tag.ID = int64(tagID)

	if _, err := tc.db.NewDelete().Model(&tag).WherePK().Exec(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not delete tag")
	}

	return c.JSON(nil)
}
