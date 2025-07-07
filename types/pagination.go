package types

type Pagination struct {
	Skip int `json:"skip"`
	Take int `json:"take"`
}