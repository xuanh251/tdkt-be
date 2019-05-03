package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
	"errors"
)

var (
	ErrChucDanhNotFound = errors.New("Không tìm thấy chức danh này!")
)

func CreateChucDanh(ChucDanh models.ChucDanh) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(ChucDanh)
	if addRc {
		err := con.Create(&ChucDanh).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
func GetAllChucDanh() ([]models.ChucDanh, error) {
	con := driver.Connect()
	defer con.Close()
	listChucDanh := []models.ChucDanh{}
	err := con.Find(&listChucDanh).Error
	if err != nil {
		return nil, err
	}
	return listChucDanh, nil
}

func GetChucDanhByID(ma_chuc_danh uint32) (models.ChucDanh, error) {
	con := driver.Connect()
	defer con.Close()
	chucVu := models.ChucDanh{}
	err := con.Find(&chucVu, ma_chuc_danh).Error
	if err != nil {
		return models.ChucDanh{}, err
	}
	return chucVu, nil
}

func UpdateChucDanh(chucdanh models.ChucDanh) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	err := con.Save(&chucdanh).Error
	if err != nil {
		return 0, err
	}
	return con.RowsAffected, nil
}
func DeleteChucDanh(ma_chuc_danh uint32) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	chuc_danh, err := GetChucDanhByID(ma_chuc_danh)
	if err != nil {
		return 0, err
	}
	err = con.Delete(chuc_danh).Error
	if err != nil {
		return 0, err
	}
	return con.RowsAffected, nil
}
