package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/rafael-abuawad/metropolis/api_controllers"
	"github.com/rafael-abuawad/metropolis/database"
)

func init() {
	database.Init()
}

func main() {
	engine := html.New("./views", ".html")
	engine.Reload(true)

	var (
		companyApiController = api_controllers.NewCompanyApiController(database.Bun)
		jobPostApiController = api_controllers.NewJobPostApiController(database.Bun)
		jobApiController     = api_controllers.NewJobApiController(database.Bun)
		tagApiController     = api_controllers.NewTagApiController(database.Bun)
	)

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "_layouts/base",
	})

	api := app.Group("/api/v1")

	api.Get("/companies", companyApiController.GetAllCompanies)
	api.Get("/companies/:id<int>", companyApiController.GetCompanyByID)
	api.Post("/companies", companyApiController.CreateCompany)
	api.Put("/companies/:id<int>", companyApiController.UpdateCompany)
	api.Delete("/companies/:id<int>", companyApiController.DeleteCompany)

	api.Get("/job-posts", jobPostApiController.GetAllJobPosts)
	api.Get("/job-posts/:id<int>", jobPostApiController.GetJobPostByID)
	api.Post("/job-posts", jobPostApiController.CreateJobPost)
	api.Put("/job-posts/:id<int>", jobPostApiController.UpdateJobPost)
	api.Delete("/job-posts/:id<int>", jobPostApiController.DeleteJobPost)

	api.Get("/jobs", jobApiController.GetAllJobs)
	api.Get("/jobs/:id<int>", jobApiController.GetJobByID)
	api.Post("/jobs", jobApiController.CreateJob)
	api.Put("/jobs/:id<int>", jobApiController.UpdateJob)
	api.Delete("/jobs/:id<int>", jobApiController.DeleteJob)

	api.Get("/tags", tagApiController.GetAllTags)
	api.Get("/tags/:id<int>", tagApiController.GetTagByID)
	api.Post("/tags", tagApiController.CreateTag)
	api.Put("/tags/:id<int>", tagApiController.UpdateTag)
	api.Delete("/tags/:id<int>", tagApiController.DeleteTag)

	app.Get("/", func(c *fiber.Ctx) error {
		tags := []string{
			"Tecnología (72)",
			"Salud (45)",
			"Educación (89)",
			"Finanzas (61)",
			"Marketing (34)",
			"Ventas (27)",
			"Diseño (50)",
			"Ingeniería (93)",
			"Gastronomía (14)",
			"Deporte (67)",
			"Arte (18)",
			"Turismo (80)",
			"Transporte (58)",
			"Medioambiente (39)",
			"Ciencia (99)",
			"Construcción (31)",
			"Moda (15)",
			"Música (88)",
			"Servicios (77)",
			"Investigación (20)",
			"Cine (54)",
			"Literatura (22)",
			"Telecomunicaciones (63)",
			"Agricultura (37)",
			"Energía (79)",
			"Derecho (41)",
			"Recursos Humanos (70)",
			"Seguridad (11)",
			"Logística (76)",
			"Comunicación (84)",
			"Social (33)",
			"Arquitectura (48)",
			"Televisión (94)",
		}

		return c.Render("home", fiber.Map{
			"Tags": tags,
		})
	})

	app.Get("/jobs/:id", func(c *fiber.Ctx) error {
		tags := []string{
			"Tecnología",
			"Salud",
			"Educación",
			"Finanzas",
		}

		return c.Render("job/job-details", fiber.Map{
			"Tags": tags,
		})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("login", fiber.Map{})
	})

	app.Static("/public", "./public")

	log.Fatal(app.Listen(":3000"))
}
