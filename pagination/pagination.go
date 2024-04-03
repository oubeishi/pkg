package pagination

type Pagination struct {
	Total    int `json:"total"`    // Total count
	Current  int `json:"current"`  // Current Page
	PageSize int `json:"pageSize"` // Page Size
}

type Param struct {
	Pagination bool         `form:"pagination"`                                   // Pagination 是否分页
	OnlyCount  bool         `form:"only_count"`                                   // Only count 是否只返回总数
	Current    int          `form:"current,default=1" binding:"min=1"`            // Current page
	PageSize   int          `form:"page_size,default=10" binding:"max=100,min=1"` // Page size
	Orders     []OrderField `form:"orders"`                                       // Order fields
}

func (a Param) GetCurrent() int {
	return a.Current
}

func (a Param) GetPageSize() int {
	pageSize := a.PageSize
	if a.PageSize == 0 {
		pageSize = 100
	}
	return pageSize
}

type OrderDirection int

type OrderField struct {
	Key  string `form:"key"`
	Desc bool   `form:"desc"`
}

//func NewOrderFields(orderFields ...*OrderField) []*OrderField {
//	return orderFields
//}
//
//func NewOrderField(key string, d OrderDirection) *OrderField {
//	return &OrderField{
//		Key:       key,
//		Direction: d,
//	}
//}

type OrderFields []*OrderField

//
//func (a OrderFields) GetString() string {
//	return ""
//}
//func (a OrderFields) AddIdSortField() OrderFields {
//	return append(a, NewOrderField("id", OrderByDESC))
//}
