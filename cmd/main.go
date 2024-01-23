package main

import (
	"TOPO/appctr"
	"TOPO/internal/fixtures"
	"TOPO/internal/migrations"
	"TOPO/pkg/CLI"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	appctr.Start()
	appctr.Log().Info("Enviroment is: " + appctr.Env())

	// Do Migrations and Fixtures in Local and Dev if needed
	if appctr.Env() == appctr.EnvLocal || appctr.Env() == appctr.EnvDev {
		if appctr.Cfg().GetString("migrations") == "true" {
			migrations.MigrateSchema()
		}
		if appctr.Cfg().GetString("fixtures") == "true" {
			fixtures.MakeFixtures()
		}
	}
	// Do Migrations in Prod if needed
	if appctr.Env() == appctr.EnvProd && appctr.Cfg().GetString("migrations") == "true" {
		perm := CLI.AskForAuthorize("Do you want to do migrations in Prod? (y/n)")
		if perm == "y" {
			migrations.MigrateSchema()
		} else {
			appctr.Log().Info("Migrations in Prod skipped")
		}
	}

	appctr.UseMiddlewares(r)
	prepareIoC(r)
}
