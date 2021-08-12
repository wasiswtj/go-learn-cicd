package usecase

import (
	"errors"
	"go-learn-cicd/models"
	"testing"

	"gotest.tools/assert"
)

var calcPyAddUCTest = []models.CalcPayloadAddUsecaseTest{
	// scenario 1: success validate add and success store to db
	{IntOne: 4, IntTwo: 1, Result: 5, StoreFunc: func(crData models.CalculationResult) error { return nil }},
	// scenario 2: success validate add and error store to db
	{IntOne: 3, IntTwo: 3, Result: 6, StoreFunc: func(crData models.CalculationResult) error { return errors.New("error insert") }},
}

// create mock interface needed for calculation usecase
type calcRepoMock struct{}

var storeMock func(cr models.CalculationResult) error
var getMock func(id int64) models.CalculationResult

func (crm *calcRepoMock) Store(cr models.CalculationResult) error {
	return storeMock(cr)
}

func (crm *calcRepoMock) Get(id int64) models.CalculationResult {
	return getMock(id)
}

// init mock interface and init calculation usecase
var mock = &calcRepoMock{}
var calcUC = InitCalculations(mock)

func TestAdd(t *testing.T) {
	for _, testData := range calcPyAddUCTest {
		proccessedData := testData

		var addTestData models.CalcPayload
		mapTestAdd(&addTestData, &proccessedData)
		storeMock = proccessedData.StoreFunc

		addTestData = calcUC.Add(addTestData)

		assert.Equal(t, testData.Result, addTestData.Result)
	}
}

func mapTestAdd(data *models.CalcPayload, testData *models.CalcPayloadAddUsecaseTest) {
	data.IntOne = testData.IntOne
	data.IntTwo = testData.IntTwo
	data.Result = testData.Result
}
