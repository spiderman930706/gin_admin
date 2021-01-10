package models

type PageResult struct {
	Items    interface{} `json:"items"`
	Dict     interface{} `json:"dict"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type Data struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	SourceName string `json:"source_name"`
}
