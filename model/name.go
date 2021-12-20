package model

type Name struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (name *Name) Save() error {
	return db.Save(name).Error
}

func FindAllNames() (names []Name, err error) {
	err = db.Find(&names).Error
	return
}

func FindNameById(nameId uint) (name Name, err error) {
	err = db.First(&name, "id = ?", nameId).Error
	return
}
