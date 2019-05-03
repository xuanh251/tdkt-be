package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
	"errors"
)

var (
	ErrDonViNotFound = errors.New("Không tìm thấy đơn vị này!")
)

func CreateDonVi(DonVi models.DonVi) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(DonVi)
	if addRc {
		err := con.Create(&DonVi).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
func GetAllDonVi() ([]models.DonVi, error) {
	con := driver.Connect()
	defer con.Close()
	listDonVi := []models.DonVi{}
	err := con.Find(&listDonVi).Error
	if err != nil {
		return nil, err
	}
	return listDonVi, nil
}

func GetDonViByID(ma_don_vi uint32) (models.DonVi, error) {
	con := driver.Connect()
	defer con.Close()
	donVi := models.DonVi{}
	err := con.Find(&donVi, ma_don_vi).Error
	if err != nil {
		return models.DonVi{}, err
	}
	return donVi, nil
}

func UpdateDonVi(donvi models.DonVi) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	err := con.Save(&donvi).Error
	if err != nil {
		return 0, err
	}
	return con.RowsAffected, nil
}
func DeleteDonVi(ma_don_vi uint32) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	donVi, err := GetDonViByID(ma_don_vi)
	if err != nil {
		return 0, err
	}
	err = con.Delete(donVi).Error
	if err != nil {
		return 0, err
	}
	return con.RowsAffected, nil
}
