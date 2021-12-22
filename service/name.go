package service

import "go_web/model"

func DeleteName(nameId uint) (err error) {
	if err = model.DeleteNameById(uint(nameId)); err != nil {
		return
	}
	return
}
