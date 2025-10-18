package models

type PagedResult[T any] struct {
	Result []T `json:"result"`
	Count  int `json:"count"`
}
