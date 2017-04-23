package models

type CfgHistory struct {
	Id           int
	CfgId        int
	AppName      string
	CfgName      string
	MajorVersion int
	MinorVersion int
	CreateTime   string
	CreateBy     int
	UpdateTime   int
	UpdateBy     int
	CfgFile      string
	Environment  int
}
