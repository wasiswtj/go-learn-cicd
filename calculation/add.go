package calculation

import "go-learn-cicd/models"

type add struct{}

type UseCase interface {
	Add(p models.AddPayload) models.AddPayload
}

func InitAdd() UseCase {
	return &add{}
}

func (auc *add) Add(p models.AddPayload) models.AddPayload {
	p.Result = p.IntOne * p.IntTwo
	return p
}
