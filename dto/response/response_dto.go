package response

type PaginationResponse struct {
	Message        string                 `json:"message"`
	Status         bool                   `json:"status"`
	Data           interface{}            `json:"data"`
	PaginationItem PaginationItemResponse `json:"pagination_item"`
}

type PaginationItemResponse struct {
	TotalItems  int `json:"total_items"`
	TotalPages  int `json:"total_pages"`
	PageSize    int `json:"page_size"`
	CurrentPage int `json:"current_page"`
}

type Response struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
