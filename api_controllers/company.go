package api_controllers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rafael-abuawad/metropolis/models"
	"github.com/uptrace/bun"
)

type CompanyApiController struct {
	db *bun.DB
}

func NewCompanyApiController(db *bun.DB) *CompanyApiController {
	return &CompanyApiController{
		db: db,
	}
}

func (cc *CompanyApiController) GetAllCompanies(c *fiber.Ctx) error {
	var companies []models.Company

	if err := cc.db.NewSelect().Model(&companies).Scan(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusNotFound, "companies not found")
	}

	return c.JSON(companies)
}

func (cc *CompanyApiController) CreateCompany(c *fiber.Ctx) error {
	var company models.Company

	if err := c.BodyParser(&company); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not create company")
	}

	if _, err := cc.db.NewInsert().Model(&company).Exec(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not create company")
	}

	c.Status(http.StatusCreated)
	return c.JSON(company)
}

func (cc *CompanyApiController) GetCompanyByID(c *fiber.Ctx) error {
	companyID, err := c.ParamsInt("id")
	if err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusNotFound, "company not found")
	}

	var company models.Company
	company.ID = int64(companyID)

	if err := cc.db.NewSelect().Model(&company).Relation("JobPosts").WherePK().Scan(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not return company")
	}

	return c.JSON(company)
}

func (cc *CompanyApiController) UpdateCompany(c *fiber.Ctx) error {
	companyID, err := c.ParamsInt("id")
	if err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not update company")
	}

	var company models.Company
	company.ID = int64(companyID)
	company.UpdatedAt = time.Now()

	if err := c.BodyParser(&company); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not update company")
	}

	if _, err := cc.db.NewUpdate().Model(&company).Column("name", "twitter_url", "facebook_url", "instagram_url", "website_url", "updated_at").WherePK().Exec(c.Context()); err != nil {
		log.Error(err)
		return fiber.NewError(http.StatusBadRequest, "could not update company")
	}

	return c.JSON(company)
}

func (cc *CompanyApiController) DeleteCompany(c *fiber.Ctx) error {
	companyID, err := c.ParamsInt("id")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	var company models.Company
	company.ID = int64(companyID)

	if _, err := cc.db.NewDelete().Model(&company).WherePK().Exec(c.Context()); err != nil {
		c.Status(http.StatusNotFound)
		return fiber.NewError(http.StatusBadRequest, "could not delete company")
	}

	return c.JSON(nil)
}
