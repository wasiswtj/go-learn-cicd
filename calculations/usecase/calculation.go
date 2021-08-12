package usecase

import (
	"go-learn-cicd/calculations/repository"
	"go-learn-cicd/models"

	log "github.com/sirupsen/logrus"
)

type calc struct {
	CalculationRepo repository.CalculationRepo
}

type CalculationUseCase interface {
	Add(p models.CalcPayload) models.CalcPayload
}

func InitCalculations(cRepo repository.CalculationRepo) CalculationUseCase {
	return &calc{CalculationRepo: cRepo}
}

func (auc *calc) Add(p models.CalcPayload) models.CalcPayload {
	p.Result = p.IntOne + p.IntTwo

	crData := new(models.CalculationResult)
	crData.MapStoreData(p, "add")

	err := auc.CalculationRepo.Store(*crData)
	if err != nil {
		log.Debug(err)
	}

	return p
}
