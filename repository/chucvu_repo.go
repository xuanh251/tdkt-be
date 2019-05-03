package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
	"errors"
)

var (
	ErrChucVuNotFound = errors.New("Không tìm thấy chức vụ này!")
)

func CreateChucVu(ChucVu models.ChucVu) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(ChucVu)
	if addRc {
		err := con.Create(&ChucVu).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
func GetAllChucVu() ([]models.ChucVu, error) {
	con := driver.Connect()
	defer con.Close()
	listChucVu := []models.ChucVu{}
	err := con.Find(&listChucVu).Error
	if err != nil {
		return nil, err
	}
	return listChucVu, nil
}

func GetChucVuByID(ma_chuc_vu uint32) (models.ChucVu, error) {
	con := driver.Connect()
	defer con.Close()
	chucVu := models.ChucVu{}
	err := con.Find(&chucVu, ma_chuc_vu).Error
	if err != nil {
		return models.ChucVu{}, err
	}
	return chucVu, nil
}

func UpdateChucVu(chucvu models.ChucVu) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	err := con.Save(&chucvu).Error
	if err != nil {
		return 0, err
	}
	return con.RowsAffected, nil
}
func DeleteChucVu(ma_chuc_vu uint32) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	chuc_vu, err := GetChucVuByID(ma_chuc_vu)
	if err != nil {
		return 0, err
	}
	err = con.Delete(chuc_vu).Error
	if err != nil {
		return 0, err
	}
	return con.RowsAffected, nil
}
