package models

import "reflect"

//jquery datatable 输入参数
type DataTableParameter struct {
	/// DataTable请求服务器端次数
	Echo int64 `json:sEcho`
	/// 过滤文本
	Search string `json:sSearch`
	/// 每页显示的数量
	DisplayLength int64 `json:iDisplayLength`
	/// 分页时每页跨度数量
	DisplayStart int64 `json:iDisplayStart`
	/// 列数
	Columns_i int `json:iColumns`
	/// 排序列的数量
	SortingCols int `json:iSortingCols`
	/// 逗号分割所有的列
	Columns_s int `json:sColumns`
	/// 排序的列
	SortCol_0 int `json:iSortCol_0`
	/// 排序的状态
	SortDir_0 string `json:sSortDir_0`
}

func (base *DataTableParameter) Unmarshal(input interface{}) {
	fv := reflect.ValueOf(input)

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf("sEcho")
	values := fv.MethodByName("GetInt").Call(params)
	base.Echo = values[0].Int()

	params = make([]reflect.Value, 1)
	params[0] = reflect.ValueOf("iDisplayLength")
	values = fv.MethodByName("GetInt").Call(params)
	base.DisplayLength = values[0].Int()

	params = make([]reflect.Value, 1)
	params[0] = reflect.ValueOf("iDisplayStart")
	values = fv.MethodByName("GetInt").Call(params)
	base.DisplayStart = values[0].Int()

	params = make([]reflect.Value, 1)
	params[0] = reflect.ValueOf("sSearch")
	values = fv.MethodByName("GetString").Call(params)
	base.Search = values[0].String()

}
