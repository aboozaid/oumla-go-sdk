package types

type DefaultResponse[T any] struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type PaginationResponse struct {
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
	Pagination
}

type PaginatedResponse[T any] struct {
	DefaultResponse[T]
	Pagination PaginationResponse `json:"pagination"`
}