package request

type PageInfo struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	KeyWord  string `json:"keyWord"`
}
type Empty struct {
}
