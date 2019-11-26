package models

//jquery datatable 返回参数
type DataTableResult struct {
	/// DataTable请求服务器端次数
	Echo                int64       `json:"sEcho"`
	TotalRecords        int64       `json:"iTotalRecords"`
	TotalDisplayRecords int64       `json:"iTotalDisplayRecords"`
	Data                interface{} `json:"aaData"`
}
