package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
)

func GetDanhSach_TDTT(maHoiDong string) ([]models.XetThiDuaTapThe, error) {
	con := driver.Connect()
	defer con.Close()
	data := []models.XetThiDuaTapThe{}
	list := []models.XetThiDuaTapThe{}
	err := con.Where("ma_hoi_dong=?", maHoiDong).Find(&data).Error
	if err != nil {
		return nil, err
	}
	for _, obj := range data {
		obj.DonVi, err = GetDonViByID(obj.MaDonVi)
		if err != nil {
			return nil, err
		}
		obj.DanhHieuThiDua, err = GetDanhHieuByID(obj.MaDanhHieu)
		if err != nil {
			return nil, err
		}
		list = append(list, obj)
	}
	return list, nil
}
func GetObjById_TDTT(id int8) (models.XetThiDuaTapThe, error) {
	con := driver.Connect()
	defer con.Close()
	data := models.XetThiDuaTapThe{}
	err := con.Find(&data, id).Error
	if err != nil {
		return models.XetThiDuaTapThe{}, err
	}
	data.DonVi, err = GetDonViByID(data.MaDonVi)
	if err != nil {
		return models.XetThiDuaTapThe{}, err
	}
	data.DanhHieuThiDua, err = GetDanhHieuByID(data.MaDanhHieu)
	if err != nil {
		return models.XetThiDuaTapThe{}, err
	}
	return data, nil
}
func CapNhatXetThiDuaTapThe(data models.XetThiDuaTapThe) (bool, error) {
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
func DeleteObj_TDTT(id int8) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	obj, err := GetObjById_TDTT(id)
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

func GetDanhSach_TDCN(maHoiDong string) ([]models.XetThiDuaCaNhan, error) {
	con := driver.Connect()
	defer con.Close()
	data := []models.XetThiDuaCaNhan{}
	list := []models.XetThiDuaCaNhan{}
	err := con.Where("ma_hoi_dong=?", maHoiDong).Find(&data).Error
	if err != nil {
		return nil, err
	}
	for _, obj := range data {
		obj.CanBo, err = GetCanBoByID(obj.MaCanBo)
		if err != nil {
			return nil, err
		}
		obj.DanhHieuThiDua, err = GetDanhHieuByID(obj.MaDanhHieu)
		if err != nil {
			return nil, err
		}
		list = append(list, obj)
	}
	return list, nil
}
func GetObjById_TDCN(id int8) (models.XetThiDuaCaNhan, error) {
	con := driver.Connect()
	defer con.Close()
	data := models.XetThiDuaCaNhan{}
	err := con.Find(&data, id).Error
	if err != nil {
		return models.XetThiDuaCaNhan{}, err
	}
	data.CanBo, err = GetCanBoByID(data.MaCanBo)
	if err != nil {
		return models.XetThiDuaCaNhan{}, err
	}
	data.DanhHieuThiDua, err = GetDanhHieuByID(data.MaDanhHieu)
	if err != nil {
		return models.XetThiDuaCaNhan{}, err
	}
	return data, nil
}
func CapNhatXetThiDuaCaNhan(data models.XetThiDuaCaNhan) (bool, error) {
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
func DeleteObj_TDCN(id int8) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	obj, err := GetObjById_TDCN(id)
	if err != nil {
		return false, err
	}
	err = con.Delete(obj).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
