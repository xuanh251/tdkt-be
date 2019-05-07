package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
)

func NapDanhSachTapTheKT() ([]*models.TapTheQuaTrinh, error) {
	con := driver.Connect()
	defer con.Close()
	listtt := []*models.TapTheQuaTrinh{}
	err := con.Find(&listtt).Error
	for _, elem := range listtt {
		err = con.Find(&elem.DonVi, elem.MaDonVi).Error
		if err != nil {
			return nil, err
		}
		err = con.Find(&elem.NamHoc, elem.MaNamHoc).Error
		if err != nil {
			return nil, err
		}
		for _, kt := range elem.ListKhenThuong {
			obj := models.HinhThucKhenThuong{}
			err := con.Find(&obj, kt).Error
			if err != nil {
				return nil, err
			}
			elem.HinhThucKhenThuong = append(elem.HinhThucKhenThuong, obj)
		}
	}
	if err != nil {
		return nil, err
	}
	return listtt, nil
}
func NapDanhSachCaNhanKT() ([]*models.CaNhanQuaTrinh, error) {
	con := driver.Connect()
	defer con.Close()
	listtt := []*models.CaNhanQuaTrinh{}
	err := con.Find(&listtt).Error
	for _, elem := range listtt {
		err = con.Find(&elem.CanBo, elem.MaCanBo).Error
		if err != nil {
			return nil, err
		}
		err = con.Find(&elem.NamHoc, elem.MaNamHoc).Error
		if err != nil {
			return nil, err
		}
		err = con.Find(&elem.CanBo.ChucVu, elem.CanBo.MaChucVu).Error
		if err != nil {
			return nil, err
		}
		err = con.Find(&elem.CanBo.DonVi, elem.CanBo.MaDonVi).Error
		if err != nil {
			return nil, err
		}
		for _, kt := range elem.ListKhenThuong {
			obj := models.HinhThucKhenThuong{}
			err := con.Find(&obj, kt).Error
			if err != nil {
				return nil, err
			}
			elem.HinhThucKhenThuong = append(elem.HinhThucKhenThuong, obj)
		}
	}
	if err != nil {
		return nil, err
	}
	return listtt, nil
}
