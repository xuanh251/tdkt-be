package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
	"errors"
)

var (
	ErrHinhThucNotFound = errors.New("Không tìm thấy hình thức khen thưởng này!")
)

func CreateHinhThucKhenThuong(HinhThucKhenThuong models.HinhThucKhenThuong) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(HinhThucKhenThuong)
	if addRc {
		err := con.Create(&HinhThucKhenThuong).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
func GetAllHinhThucKhenThuong() ([]models.HinhThucKhenThuong, error) {
	con := driver.Connect()
	defer con.Close()
	listHinhThuc := []models.HinhThucKhenThuong{}
	err := con.Find(&listHinhThuc).Error
	if err != nil {
		return nil, err
	}
	return listHinhThuc, nil
}

func GetHinhThucKhenThuongByID(ma_hinh_thuc uint32) (models.HinhThucKhenThuong, error) {
	con := driver.Connect()
	defer con.Close()
	hinhThuc := models.HinhThucKhenThuong{}
	err := con.Find(&hinhThuc, ma_hinh_thuc).Error
	if err != nil {
		return models.HinhThucKhenThuong{}, err
	}
	return hinhThuc, nil
}

func UpdateHinhThucKhenThuong(hinhthuc models.HinhThucKhenThuong) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	err := con.Save(&hinhthuc).Error
	if err != nil {
		return 0, err
	}
	return con.RowsAffected, nil
}
func DeleteHinhThucKhenThuong(ma_hinh_thuc uint32) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	hinhThuc, err := GetHinhThucKhenThuongByID(ma_hinh_thuc)
	if err != nil {
		return 0, err
	}
	err = con.Delete(hinhThuc).Error
	if err != nil {
		return 0, err
	}
	return con.RowsAffected, nil
}
