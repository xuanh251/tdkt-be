package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
)

func GetListQuyenLimit() ([]models.Quyen, error) {
	con := driver.Connect()
	defer con.Close()
	var listQuyen []models.Quyen
	err := con.Where("ma_quyen<?", "3").Find(&listQuyen).Error
	if err != nil {
		return nil, err
	}
	return listQuyen, nil

}
func GetQuyenByID(ma_quyen uint32) (models.Quyen, error) {
	con := driver.Connect()
	defer con.Close()
	quyen := models.Quyen{}
	err := con.Find(&quyen, ma_quyen).Error
	if err != nil {
		return models.Quyen{}, err
	}
	return quyen, nil
}
