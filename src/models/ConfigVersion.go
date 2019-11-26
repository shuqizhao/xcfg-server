package models

type RemoteConfigSectionCollection struct {
	Machine     string                 `xml:"machine,attr"`
	Application string                 `xml:"application,attr"`
	Environment string                 `xml:"env,attr"`
	Sections    []*RemoteConfigSection `xml:"section"`
}
type RemoteConfigSection struct {
	SectionName  string `xml:"name,attr"`
	MajorVersion int    `xml:"majorVerion,attr"`
	MinorVersion int    `xml:"minorVerion,attr"`
	DownloadUrl  string `xml:"downloadUrl,attr"`
	TemplateUrl  string `xml:"templateUrl,attr"`
}
