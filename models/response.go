package models

type PageResult struct {
	Items     interface{} `json:"items"`
	Dict      interface{} `json:"dict"`
	Total     int64       `json:"total"`
	Page      int         `json:"page"`
	PageSize  int         `json:"pageSize"`
	CanModify bool        `json:"can_modify"`
	CanDelete bool        `json:"can_delete"`
}

type DataResult struct {
	Item      interface{} `json:"item"`
	Dict      interface{} `json:"dict"`
	CanModify bool        `json:"can_modify"`
	CanDelete bool        `json:"can_delete"`
}
