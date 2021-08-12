package handler

import (
	"encoding/json"
	"go-learn-cicd/middleware"
	"go-learn-cicd/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"gotest.tools/assert"
)

// test data
var calcPyAddHandlerTest = []models.CalcPayloadAddHandlerTest{
	// scenario 1: success validate add and success store to db
	{Url: "/add?", QueryParam: "intOne=4&intTwo=1",
		AddFunc: func(p models.CalcPayload) models.CalcPayload {
			return models.CalcPayload{
				IntOne: 4, IntTwo: 1, Result: 5,
			}
		},
		JsonRespData: `{"intOne":4,"intTwo":1,"result":5}`,
		JsonRespCode: 200,
	},
	// scenario 2 data request payload not valid contain string
	{Url: "/add?", QueryParam: "intOne=a&intTwo=1",
		JsonRespCode: 400,
	},
	// scenario 3 data request payload not complete
	{Url: "/add?", QueryParam: "intOne=&intTwo=1",
		JsonRespCode: 400,
	},
}

// create mock interface needed for calculation usecase
type calcUCMock struct{}

var addMock func(cr models.CalcPayload) models.CalcPayload

func (crm *calcUCMock) Add(p models.CalcPayload) models.CalcPayload {
	return addMock(p)
}

// init mock interface and init calculation usecase
var mock = &calcUCMock{}
var c = echo.New()
var calcHand = InitCalculationHandler(c, mock)

func TestAdd(t *testing.T) {
	middleware.InitMiddleware(c)

	for _, testData := range calcPyAddHandlerTest {
		addMock = testData.AddFunc

		req := httptest.NewRequest(http.MethodGet, testData.Url+testData.QueryParam, nil)
		rec := httptest.NewRecorder()

		resp, respAssertion := models.CalcPayload{}, models.CalcPayload{}
		err := calcHand.AddHandler(c.NewContext(req, rec))
		if err != nil {
			rec.Code = err.(*echo.HTTPError).Code
		}
		_ = json.NewDecoder(rec.Body).Decode(&resp)
		_ = json.Unmarshal([]byte(testData.JsonRespData), &respAssertion)
		assert.Equal(t, rec.Code, testData.JsonRespCode)
		assert.Equal(t, resp, respAssertion)
	}
}
