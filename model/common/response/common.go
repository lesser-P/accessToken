package response

type PageResult struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Records  any `json:"records"`
	Total    int `json:"total"`
}
