package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
)

func GetListBauChonThiDuaTapTheByDanhHieu(maDanhHieu int8, maHoiDong string) ([]models.XetThiDuaTapThe, error) {
	con := driver.Connect()
	defer con.Close()
	tt := []models.XetThiDuaTapThe{}
	dtt := []models.XetThiDuaTapThe{}
	err := con.Where("ma_danh_hieu=? and ma_hoi_dong=?", maDanhHieu, maHoiDong).Find(&tt).Error
	if err != nil {
		return nil, err
	}
	for _, elem := range tt {
		elem.DonVi, err = GetDonViByID(elem.MaDonVi)
		if err != nil {
			return nil, err
		}
		dtt = append(dtt, elem)
	}
	return dtt, nil
}
func BauChonThiDuaTapThe(obj []models.BauChonThiDuaTapThe) error {
	con := driver.Connect()
	defer con.Close()
	for _, elem := range obj {
		addRc := con.NewRecord(&elem)
		if addRc {
			err := con.Create(&elem).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}
