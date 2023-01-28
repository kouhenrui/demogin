package pojo

type Examp struct {
	Base Base   `gorm:"embedded"`
	Name string `json:"name" gorm:"default:admin;unique:true"`
}