package calculation

import (
	"go-learn-cicd/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var AddTestData = []models.AddPayload{
	{IntOne: 1, IntTwo: 1, Result: 2},
	{IntOne: 1, IntTwo: 2, Result: 3},
	{IntOne: 4, IntTwo: 1, Result: 5},
	{IntOne: 3, IntTwo: 3, Result: 6},
}

func TestAdd(t *testing.T) {
	for _, testData := range AddTestData {
		addFunc := add{}
		proccessedData := testData

		proccessedData = addFunc.Add(proccessedData)

		assert.Equal(t, testData.Result, proccessedData.Result)
	}
}
