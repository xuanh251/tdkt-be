package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
)

func CreateCanCu(CanCu models.CanCu) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(CanCu)
	if addRc {
		err := con.Create(&CanCu).Error
		if err != nil {
			return 0, err
		}
	}
	return CanCu.MaCanCu, nil
}
func GetAllCanCuChung() ([]models.CanCu, error) {
	con := driver.Connect()
	defer con.Close()
	listCanCu := []models.CanCu{}
	err := con.Where("trang_thai=?", 0).Find(&listCanCu).Error
	if err != nil {
		return nil, err
	}
	return listCanCu, nil
}
func GetCanCuRiengByHoiDong(maHoiDong string) ([]models.CanCu, error) {
	con := driver.Connect()
	defer con.Close()
	listCanCu := []models.CanCu{}
	return listCanCu, nil
}

func GetCanCuByID(ma_can_cu int64) (models.CanCu, error) {
	con := driver.Connect()
	defer con.Close()
	canCu := models.CanCu{}
	err := con.Find(&canCu, "ma_can_cu=?", ma_can_cu).Error
	if err != nil {
		return models.CanCu{}, err
	}
	return canCu, nil
}

func UpdateCanCu(cancu models.CanCu) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	err := con.Save(&cancu).Error
	if err != nil {
		return 0, err
	}
	return con.RowsAffected, nil
}
func DeleteCanCu(ma_can_cu int64) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	can_cu, err := GetCanCuByID(ma_can_cu)
	if err != nil {
		return 0, err
	}
	err = con.Delete(can_cu).Error
	if err != nil {
		return 0, err
	}
	return con.RowsAffected, nil
}
