package models

type Dataweb struct {
	Id          int    `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description" gorm:"type:text"`
	Favico      string `json:"favico"`
	Logo        string `json:"logo"`
	Footer      string `json:"footer" gorm:"type:text"`
}
