package database

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/rafael-abuawad/metropolis/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

var Bun *bun.DB

func Init() {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	Bun = bun.NewDB(sqldb, sqlitedialect.New())
	Bun.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	createSchema()
	seed()
}

func createSchema() {
	ctx := context.Background()

	Bun.RegisterModel(
		(*models.CompanyToJobPost)(nil),
		(*models.JobPostToTag)(nil),
	)

	models := []interface{}{
		(*models.Tag)(nil),
		(*models.Job)(nil),
		(*models.Company)(nil),
		(*models.CompanyToJobPost)(nil),
		(*models.JobPost)(nil),
		(*models.JobPostToTag)(nil),
	}

	for _, model := range models {
		if _, err := Bun.NewCreateTable().Model(model).Exec(ctx); err != nil {
			log.Fatal(err)
		}
	}
}

func seed() {
	ctx := context.Background()

	Bun.RegisterModel(
		(*models.Tag)(nil),
		(*models.Job)(nil),
		(*models.Company)(nil),
		(*models.CompanyToJobPost)(nil),
		(*models.JobPost)(nil),
		(*models.JobPostToTag)(nil),
	)

	fixture := dbfixture.New(Bun)
	err := fixture.Load(ctx, os.DirFS("database/testdata"), "fixture.yaml")
	if err != nil {
		log.Fatal(err)
	}
}
