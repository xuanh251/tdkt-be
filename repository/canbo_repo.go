package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
	"TDKTServer/utils"
	"encoding/json"
	"errors"
)

var (
	ErrCanBoNotFound = errors.New("Không tìm thấy cán bộ này!")
)

func NewCanBo(CanBo models.CanBo) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	CanBo.Password = "123456"
	hashedPassword, err := utils.Bcrypt(CanBo.Password)
	if err != nil {
		return false, err
	}
	CanBo.Password = string(hashedPassword)
	addRc := con.NewRecord(CanBo)
	if addRc {
		err := con.Create(&CanBo).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
func GetAllCanBo() ([]models.CanBo, error) {
	con := driver.Connect()
	defer con.Close()
	var listcanbo []models.CanBo
	var listdata []models.CanBo
	err := con.Find(&listcanbo).Error
	if err != nil {
		return nil, err
	}
	for _, canbo := range listcanbo {
		canbo.DonVi, err = GetDonViByID(canbo.MaDonVi)
		if err != nil {
			return nil, err
		}
		canbo.ChucVu, err = GetChucVuByID(canbo.MaChucVu)
		if err != nil {
			return nil, err
		}
		listdata = append(listdata, canbo)
	}
	return listdata, nil
}

func GetCanBoByID(id uint32) (models.CanBo, error) {
	con := driver.Connect()
	defer con.Close()
	canbo := models.CanBo{}
	err := con.First(&canbo, id).Error
	if err != nil {
		return models.CanBo{}, err
	}
	if canbo.MaCanBo == 0 {
		return models.CanBo{}, ErrCanBoNotFound
	}
	canbo.NgaySinh = canbo.NgaySinh[0:10]
	canbo.ChucVu, err = GetChucVuByID(canbo.MaChucVu)
	if err != nil {
		return models.CanBo{}, err
	}
	canbo.DonVi, err = GetDonViByID(canbo.MaDonVi)
	if err != nil {
		return models.CanBo{}, err
	}
	canbo.Quyen, err = GetQuyenByID(canbo.MaQuyen)
	if err != nil {
		return models.CanBo{}, err
	}
	return canbo, nil
}
func GetCanBoByInfo(cb *models.CanBo) (models.CanBo, error) {
	con := driver.Connect()
	defer con.Close()
	canbo := models.CanBo{}
	err := con.First(&canbo, cb.MaCanBo).Error
	if err != nil {
		return models.CanBo{}, err
	}
	if canbo.MaCanBo == 0 {
		return models.CanBo{}, ErrCanBoNotFound
	}
	canbo.NgaySinh = canbo.NgaySinh[0:10]
	canbo.ChucVu, err = GetChucVuByID(canbo.MaChucVu)
	if err != nil {
		return models.CanBo{}, err
	}
	canbo.DonVi, err = GetDonViByID(canbo.MaDonVi)
	if err != nil {
		return models.CanBo{}, err
	}
	canbo.Quyen, err = GetQuyenByID(canbo.MaQuyen)
	if err != nil {
		return models.CanBo{}, err
	}
	return canbo, nil
}
func GetCanBoByEmail(email string) (models.CanBo, error) {
	con := driver.Connect()
	defer con.Close()
	acc := models.CanBo{}
	err := con.Where("email=?", email).First(&acc).Error
	if err != nil {
		return models.CanBo{}, err
	}
	return acc, nil
}

func UpdateCanBo(canbo models.CanBo) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	acc2Update, err := GetCanBoByID(canbo.MaCanBo)
	if err != nil {
		return 0, err
	}
	canbo.Password = acc2Update.Password
	canbo.AnhDaiDien = acc2Update.AnhDaiDien
	err = con.Save(&canbo).Error
	if err != nil {
		return 0, err
	}
	return con.RowsAffected, nil
}
func DeleteCanBo(id uint32) (int64, error) {
	con := driver.Connect()
	defer con.Close()
	canbo, err := GetCanBoByID(id)
	if err != nil {
		return 0, err
	}
	err = con.Delete(&canbo).Error
	if err != nil {
		return 0, err
	}
	return con.RowsAffected, nil
}
func SetActive(maCanBo uint32) error {
	con := driver.Connect()
	defer con.Close()
	canbo, err := GetCanBoByID(maCanBo)
	if err != nil {
		return err
	}
	canbo.TrangThai = true
	err = con.Save(&canbo).Error
	if err != nil {
		return err
	}
	jsonMessage, _ := json.Marshal(&utils.Message{Content: "Cán bộ " + canbo.HoTen + " đã online!"})
	utils.Manager.Broadcast <- jsonMessage
	return nil
}
func SetDeactive(maCanBo uint32) error {
	con := driver.Connect()
	defer con.Close()
	canbo, err := GetCanBoByID(maCanBo)
	if err != nil {
		return err
	}
	canbo.TrangThai = false
	err = con.Save(&canbo).Error
	if err != nil {
		return err
	}
	jsonMessage, _ := json.Marshal(&utils.Message{Content: "Cán bộ " + canbo.HoTen + " đã thoát!"})
	utils.Manager.Broadcast <- jsonMessage
	return nil
}
