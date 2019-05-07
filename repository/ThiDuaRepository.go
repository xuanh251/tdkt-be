package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
)

func NapDanhSachTapTheTD() ([]*models.TapTheQuaTrinh, error) {
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
		for _, dh := range elem.ListDanhHieu {
			obj := models.DanhHieuThiDua{}
			err := con.Find(&obj, dh).Error
			if err != nil {
				return nil, err
			}
			elem.DanhHieuThiDua = append(elem.DanhHieuThiDua, obj)
		}
	}
	if err != nil {
		return nil, err
	}
	return listtt, nil
}

func NapDanhSachCaNhanTD() ([]*models.CaNhanQuaTrinh, error) {
	con := driver.Connect()
	defer con.Close()
	listtt := []*models.CaNhanQuaTrinh{}
	err := con.Find(&listtt).Error
	for _, elem := range listtt {
		err = con.Find(&elem.CanBo, elem.MaCanBo).Error
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
		err = con.Find(&elem.NamHoc, elem.MaNamHoc).Error
		if err != nil {
			return nil, err
		}
		for _, dh := range elem.ListDanhHieu {
			obj := models.DanhHieuThiDua{}
			err := con.Find(&obj, dh).Error
			if err != nil {
				return nil, err
			}
			elem.DanhHieuThiDua = append(elem.DanhHieuThiDua, obj)
		}
	}
	if err != nil {
		return nil, err
	}
	return listtt, nil
}
