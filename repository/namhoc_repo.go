package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
	"TDKTServer/utils"
)

func CreateNamHoc(NamHoc models.NamHoc) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(NamHoc)
	if addRc {
		err := con.Create(&NamHoc).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
func GetAllNamHoc() ([]models.NamHoc, error) {
	con := driver.Connect()
	defer con.Close()
	listData := []models.NamHoc{}
	listNamHoc := []models.NamHoc{}
	err := con.Find(&listData).Error
	if err != nil {
		return nil, err
	}

	for _, elem := range listData {
		elem.TuNgay = utils.ConvertSchoolYear(elem)
		listNamHoc = append(listNamHoc, elem)
	}
	return listNamHoc, nil
}
func GetNamHocByID(ma_nam_hoc uint32) (models.NamHoc, error) {
	con := driver.Connect()
	defer con.Close()
	namHoc := models.NamHoc{}
	err := con.Find(&namHoc, ma_nam_hoc).Error
	if err != nil {
		return models.NamHoc{}, err
	}
	return namHoc, nil
}
