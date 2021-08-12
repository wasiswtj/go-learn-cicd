package main

import (
	"context"
	"fmt"
	calculationHand "go-learn-cicd/calculations/handler"
	calculationRepo "go-learn-cicd/calculations/repository"
	calculationUC "go-learn-cicd/calculations/usecase"
	"go-learn-cicd/middleware"

	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// load env
	loadEnv()

	// load middleware
	c := echo.New()
	middleware.InitMiddleware(c)

	// connection pool for database
	dbPG := dbInitAndMigration()

	// repository package
	calcRepo := calculationRepo.InitCalculationPsqlRepo(dbPG)

	// load function package
	calcUC := calculationUC.InitCalculations(calcRepo)

	// append all handler
	_ = calculationHand.InitCalculationHandler(c, calcUC)

	c.Logger.Fatal(c.Start(os.Getenv(`PORT`)))
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func dbInitAndMigration() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv(`DB_HOST`) + os.Getenv(`DB_PORT`),
		User:     os.Getenv(`DB_USER`),
		Password: os.Getenv(`DB_PASS`),
		Database: os.Getenv(`DB_NAME`),
	})

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		panic(err)
	}

	// for migrations
	dbConf := db.Options()
	dbMigrationUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbConf.User, dbConf.Password, dbConf.Addr, dbConf.Database)

	m, err := migrate.New(
		"file://db/migrations",
		dbMigrationUrl)

	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()

	if err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	return db
}
