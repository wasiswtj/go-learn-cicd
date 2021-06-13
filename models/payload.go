package models

type AddPayload struct {
	IntOne int64 `json:"intOne"`
	IntTwo int64 `json:"intTwo"`
	Result int64 `json:"result"`
}
