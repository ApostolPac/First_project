package models

type Human struct {
	Id   int    `json:"-"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}
