package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
	"errors"
)

var (
	ErrDanhHieuNotFound = errors.New("Không tìm thấy danh hiệu thi đua này!")
)

func CreateDanhHieu(DanhHieu models.DanhHieuThiDua) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(DanhHieu)
	if addRc {
		err := con.Create(&DanhHieu).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
func GetAllDanhHieu() ([]models.DanhHieuThiDua, error) {
	con := driver.Connect()
	defer con.Close()
	listDanhHieu := []models.DanhHieuThiDua{}
	err := con.Find(&listDanhHieu).Error
	if err != nil {
		return nil, err
	}
	return listDanhHieu, nil
}

func GetDanhHieuByID(ma_danh_hieu uint32) (models.DanhHieuThiDua, error) {
	con := driver.Connect()
	defer con.Close()
	danhHieu := models.DanhHieuThiDua{}
	err := con.Find(&danhHieu, ma_danh_hieu).Error
	if err != nil {
		return models.DanhHieuThiDua{}, err
	}
	return danhHieu, nil
}

func UpdateDanhHieu(danhhieu models.DanhHieuThiDua) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	err := con.Save(&danhhieu).Error
	if err != nil {
		return 0, err
	}
	return con.RowsAffected, nil
}
func DeleteDanhHieu(ma_danh_hieu uint32) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	danhHieu, err := GetDanhHieuByID(ma_danh_hieu)
	if err != nil {
		return 0, err
	}
	err = con.Delete(danhHieu).Error
	if err != nil {
		return 0, err
	}
	return con.RowsAffected, nil
}
