package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
)

func GetListBauChonThiDuaTapTheByDanhHieu(maDanhHieu int8, maHoiDong string) ([]*models.XetThiDuaTapThe, error) {
	con := driver.Connect()
	defer con.Close()
	tt := []*models.XetThiDuaTapThe{}

	err := con.Where("ma_danh_hieu=? and ma_hoi_dong=?", maDanhHieu, maHoiDong).Find(&tt).Error
	if err != nil {
		return nil, err
	}
	for _, elem := range tt {
		err = con.Find(&elem.DonVi, elem.MaDonVi).Error
		if err != nil {
			return nil, err
		}
		err = con.Where("ma_xet=?", elem.ID).Find(&elem.BauChonThiDuaTapThe).Error
		if err != nil {
			return nil, err
		}
	}
	return tt, nil
}
func GetListDaBauChonTDTTByTVHD(maThanhPhan int8) ([]*models.BauChonThiDuaTapThe, error) {
	con := driver.Connect()
	defer con.Close()
	listBC := []*models.BauChonThiDuaTapThe{}
	err := con.Where("ma_thanh_phan=?", maThanhPhan).Find(&listBC).Error
	if err != nil {
		return nil, err
	}
	for _, elem := range listBC {
		err = con.Find(&elem.XetThiDuaTapThe, elem.MaXet).Error
		if err != nil {
			return nil, err
		}
		err = con.Find(&elem.XetThiDuaTapThe.DonVi, elem.XetThiDuaTapThe.MaDonVi).Error
		if err != nil {
			return nil, err
		}
	}
	return listBC, nil
}
func GetListBauChonThiDuaCaNhanByDanhHieu(maDanhHieu int8, maHoiDong string) ([]models.XetThiDuaCaNhan, error) {
	con := driver.Connect()
	defer con.Close()
	cn := []models.XetThiDuaCaNhan{}
	dcn := []models.XetThiDuaCaNhan{}
	err := con.Where("ma_danh_hieu=? and ma_hoi_dong=?", maDanhHieu, maHoiDong).Find(&cn).Error
	if err != nil {
		return nil, err
	}
	for _, elem := range cn {
		elem.CanBo, err = GetCanBoByID(elem.MaCanBo)
		if err != nil {
			return nil, err
		}
		dcn = append(dcn, elem)
	}
	return dcn, nil
}
func AddBauChonThiDuaTapThe(obj models.BauChonThiDuaTapThe) error {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(&obj)
	if addRc {
		err := con.Create(&obj).Error
		if err != nil {
			return err
		}
	}
	return nil
}
func AddBauChonThiDuaCaNhan(obj models.BauChonThiDuaCaNhan) error {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(&obj)
	if addRc {
		err := con.Create(&obj).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func AddBauChon2QuaTrinh(obj models.TapTheQuaTrinh) error {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(&obj)
	if addRc {
		err := con.Create(&obj).Error
		if err != nil {
			return err
		}
	}

	return nil
}
