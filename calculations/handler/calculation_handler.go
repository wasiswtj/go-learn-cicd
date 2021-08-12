package handler

import (
	"go-learn-cicd/calculations/usecase"
	"go-learn-cicd/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type calculationHandler struct {
	CalculationUseCase usecase.CalculationUseCase
	Echo               *echo.Echo
}

type CalculationHandler interface {
	AddHandler(c echo.Context) error
}

func InitCalculationHandler(c *echo.Echo, calcUC usecase.CalculationUseCase) CalculationHandler {
	calcHandler := &calculationHandler{CalculationUseCase: calcUC, Echo: c}

	// path endpoint
	calcHandler.Echo.GET("/add", calcHandler.AddHandler)

	return calcHandler
}

func (cH *calculationHandler) AddHandler(c echo.Context) error {
	p := new(models.CalcPayload)

	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp := cH.CalculationUseCase.Add(*p)
	return c.JSON(http.StatusOK, resp)
}
