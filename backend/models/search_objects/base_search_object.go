package searchobjects

type BaseSearchObject struct {
	Page     int `form:"page" json:"page" binding:"gte=0"`
	PageSize int `form:"page_size" json:"page_size" binding:"gte=0"`
}
