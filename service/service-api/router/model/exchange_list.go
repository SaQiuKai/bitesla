package model

type ExchangeList struct {
	page int32 `json:"page" example:"1"`
	size int32 `json:"size" example:"10"`
}
