package models

type Pokemon struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Weight int `json:"weight"`
	Base_experience int `json:"base_experience"`
	Sprites Sprites `json:"sprites"`
}

type Sprites struct {
	Front_default string `json:"front_default"`
	Back_default string `json:"back_default"`
}