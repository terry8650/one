package common

// PageReq 公共请求参数
type PageReq struct {
	DateRange []string `p:"dateRange"`                                                        //日期范围
	PageNum   int      `p:"page" v:"required|min:1# page is required"`                        //当前页码
	PageSize  int      `p:"limit" default:"20"  v:"required|between:1,100# size is required"` //每页数
	OrderBy   string   //排序方式
}

// ListRes 列表公共返回
type ListRes struct {
	Code    int    `json:"code"` // 错误码((0:成功, 1:失败, >1:错误码))
	Message string `json:"msg"`
	//CurrentPage int    `json:"page"`
	Total int `json:"count"`
}
