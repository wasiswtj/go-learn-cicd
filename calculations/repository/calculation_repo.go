package repository

import (
	"go-learn-cicd/models"

	"github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
)

type calculationRepo struct {
	pgDB *pg.DB
}

type CalculationRepo interface {
	Store(calcRes models.CalculationResult) error
	Get(id int64) models.CalculationResult
}

func InitCalculationPsqlRepo(pgDB *pg.DB) CalculationRepo {
	return &calculationRepo{pgDB: pgDB}
}

func (cpr *calculationRepo) Store(calcRes models.CalculationResult) error {
	_, err := cpr.pgDB.Model(&calcRes).Insert()

	if err != nil {
		log.Debug(err)
		return err
	}

	return nil
}

func (cpr *calculationRepo) Get(id int64) models.CalculationResult {
	var calcRes models.CalculationResult
	calcRes.ID = id

	err := cpr.pgDB.Model(&calcRes).Select()

	if err != nil {
		log.Debug(err)
		return models.CalculationResult{}
	}

	return calcRes
}
