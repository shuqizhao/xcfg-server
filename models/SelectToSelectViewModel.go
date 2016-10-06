package models

type SelectToSelectViewModel struct {
	Left  []IdNameViewModel `json:"left"`
	Right []IdNameViewModel `json:"right"`
}

type IdNameViewModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
