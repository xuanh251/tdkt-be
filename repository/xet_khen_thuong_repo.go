package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
)

func GetDanhSach_KTTT(maHoiDong string) ([]models.XetKhenThuongTapThe, error) {
	con := driver.Connect()
	defer con.Close()
	data := []models.XetKhenThuongTapThe{}
	list := []models.XetKhenThuongTapThe{}
	err := con.Where("ma_hoi_dong=?", maHoiDong).Find(&data).Error
	if err != nil {
		return nil, err
	}
	for _, obj := range data {
		obj.DonVi, err = GetDonViByID(obj.MaDonVi)
		if err != nil {
			return nil, err
		}
		obj.HinhThucKhenThuong, err = GetHinhThucKhenThuongByID(obj.MaHinhThuc)
		if err != nil {
			return nil, err
		}
		list = append(list, obj)
	}
	return list, nil
}
func GetObjById_KTTT(id int8) (models.XetKhenThuongTapThe, error) {
	con := driver.Connect()
	defer con.Close()
	data := models.XetKhenThuongTapThe{}
	err := con.Find(&data, id).Error
	if err != nil {
		return models.XetKhenThuongTapThe{}, err
	}
	data.DonVi, err = GetDonViByID(data.MaDonVi)
	if err != nil {
		return models.XetKhenThuongTapThe{}, err
	}
	data.HinhThucKhenThuong, err = GetHinhThucKhenThuongByID(data.MaHinhThuc)
	if err != nil {
		return models.XetKhenThuongTapThe{}, err
	}
	return data, nil
}
func CapNhatXetKhenThuongTapThe(data models.XetKhenThuongTapThe) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(data)
	if addRc {
		err := con.Create(&data).Error
		if err != nil {
			return false, err
		}
	} else {
		err := con.Save(&data).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
func DeleteObj_KTTT(id int8) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	obj, err := GetObjById_KTTT(id)
	if err != nil {
		return false, err
	}
	err = con.Delete(obj).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

//-------------------------------------------------
//cá nhân

func GetDanhSach_KTCN(maHoiDong string) ([]models.XetKhenThuongCaNhan, error) {
	con := driver.Connect()
	defer con.Close()
	data := []models.XetKhenThuongCaNhan{}
	list := []models.XetKhenThuongCaNhan{}
	err := con.Where("ma_hoi_dong=?", maHoiDong).Find(&data).Error
	if err != nil {
		return nil, err
	}
	for _, obj := range data {
		obj.CanBo, err = GetCanBoByID(obj.MaCanBo)
		if err != nil {
			return nil, err
		}
		obj.HinhThucKhenThuong, err = GetHinhThucKhenThuongByID(obj.MaHinhThuc)
		if err != nil {
			return nil, err
		}
		list = append(list, obj)
	}
	return list, nil
}
func GetObjById_KTCN(id int8) (models.XetKhenThuongCaNhan, error) {
	con := driver.Connect()
	defer con.Close()
	data := models.XetKhenThuongCaNhan{}
	err := con.Find(&data, id).Error
	if err != nil {
		return models.XetKhenThuongCaNhan{}, err
	}
	data.CanBo, err = GetCanBoByID(data.MaCanBo)
	if err != nil {
		return models.XetKhenThuongCaNhan{}, err
	}
	data.HinhThucKhenThuong, err = GetHinhThucKhenThuongByID(data.MaHinhThuc)
	if err != nil {
		return models.XetKhenThuongCaNhan{}, err
	}
	return data, nil
}
func CapNhatXetKhenThuongCaNhan(data models.XetKhenThuongCaNhan) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(data)
	if addRc {
		err := con.Create(&data).Error
		if err != nil {
			return false, err
		}
	} else {
		err := con.Save(&data).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
func DeleteObj_KTCN(id int8) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	obj, err := GetObjById_KTCN(id)
	if err != nil {
		return false, err
	}
	err = con.Delete(obj).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
