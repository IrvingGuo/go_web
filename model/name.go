package model

import "go_web/config"

var db = config.Db

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (name *Album) Save() error {
	return db.Save(name).Error
}

func FindAllNames() (names []Album, err error) {
	err = db.Find(&names).Error
	return
}

func FindNameById(nameId uint) (name Album, err error) {
	err = db.First(&name, "id = ?", nameId).Error
	return
}

func DeleteNameById(id uint) error {
	return db.Where("id = ?", id).Delete(&Album{}).Error
}
