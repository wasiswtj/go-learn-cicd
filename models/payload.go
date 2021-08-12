package models

type CalcPayload struct {
	IntOne int64 `json:"intOne" param:"intOne" query:"intOne" validate:"required"`
	IntTwo int64 `json:"intTwo" param:"intTwo" query:"intTwo" validate:"required"`
	Result int64 `json:"result"`
}

type CalcPayloadAddUsecaseTest struct {
	IntOne    int64 `json:"intOne" param:"intOne" query:"intOne" validate:"required"`
	IntTwo    int64 `json:"intTwo" param:"intTwo" query:"intTwo" validate:"required"`
	Result    int64 `json:"result"`
	StoreFunc func(crData CalculationResult) error
}

type CalcPayloadAddHandlerTest struct {
	Url          string
	QueryParam   string
	JsonRespData string
	JsonRespCode int
	AddFunc      func(p CalcPayload) CalcPayload
}
