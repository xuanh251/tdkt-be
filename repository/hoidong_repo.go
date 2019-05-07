package repository

import (
	"TDKTServer/driver"
	"TDKTServer/models"
	"errors"
	"fmt"
	"strconv"
	"time"
)

var (
	ErrHoiDongNotFound = errors.New("Không tìm thấy hội đồng này!")
)

func GetAllHoiDong() ([]models.HoiDong, error) {
	con := driver.Connect()
	defer con.Close()

	listHoiDong := []models.HoiDong{}
	rs := []models.HoiDong{}
	err := con.Order("ngay_ky desc").Find(&listHoiDong).Error
	if err != nil {
		return nil, err
	}
	for _, value := range listHoiDong {
		value.NamHoc, err = GetNamHocByID(value.MaNamHoc)
		if err != nil {
			return nil, err
		}
		rs = append(rs, value)
	}

	return rs, nil
}
func CreateHoiDong(HoiDong models.HoiDong) (string, error) {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(HoiDong)
	timeNow := int64(time.Now().Unix()) * 1000
	if addRc {
		HoiDong.MaHoiDong = strconv.FormatInt(timeNow, 10)
		HoiDong.NgayTao = timeNow
		err := con.Create(&HoiDong).Error
		if err != nil {
			return "", err
		}
	}
	return HoiDong.MaHoiDong, nil
}
func CreateHoiDongCanCuRelate(obj models.HoiDongCanCu) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(obj)
	if addRc {
		err := con.Create(&obj).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
func CreateThanhPhanHoiDong(obj models.ThanhPhanHoiDong) (bool, error) {
	con := driver.Connect()
	defer con.Close()
	addRc := con.NewRecord(obj)
	if addRc {
		err := con.Create(&obj).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
func GetListThanhVienByHoiDong(maHoiDong string) ([]models.ThanhPhanHoiDong, error) {
	listThanhPhanHoiDong := []models.ThanhPhanHoiDong{}
	var listdata []models.ThanhPhanHoiDong
	con := driver.Connect()
	defer con.Close()
	err := con.Where("ma_hoi_dong=?", maHoiDong).Find(&listThanhPhanHoiDong).Error
	if err != nil {
		return nil, err
	}
	for _, thanhphan := range listThanhPhanHoiDong {
		thanhphan.CanBo, err = GetCanBoByID(thanhphan.MaCanBo)
		if err != nil {
			return nil, err
		}
		thanhphan.ChucDanh, err = GetChucDanhByID(thanhphan.MaChucDanh)
		if err != nil {
			return nil, err
		}
		listdata = append(listdata, thanhphan)
	}
	return listdata, nil
}
func GetHoiDongById(maHoiDong string) (models.HoiDong, error) {
	con := driver.Connect()
	defer con.Close()
	hoiDong := models.HoiDong{}
	err := con.Where("ma_hoi_dong=?", maHoiDong).Find(&hoiDong).Error
	if err != nil {
		return models.HoiDong{}, err
	}
	hoiDong.NamHoc, err = GetNamHocByID(hoiDong.MaNamHoc)
	if err != nil {
		return models.HoiDong{}, err
	}
	return hoiDong, nil
}
func GetLastHoiDong() (models.HoiDong, error) {
	con := driver.Connect()
	defer con.Close()
	hoiDong := models.HoiDong{}
	err := con.Last(&hoiDong).Error
	if err != nil {
		return models.HoiDong{}, err
	}
	//if hoiDong.TrangThai!=0 {
	//	return models.HoiDong{}, nil
	//}
	hoiDong.NamHoc, err = GetNamHocByID(hoiDong.MaNamHoc)
	if err != nil {
		return models.HoiDong{}, err
	}
	return hoiDong, nil
}
func GetListCanCuByHoiDong(maHoiDong string) ([]models.CanCu, error) {
	con := driver.Connect()
	defer con.Close()
	hdcc := []models.HoiDongCanCu{}
	rs := []models.CanCu{}
	err := con.Where("ma_hoi_dong=?", maHoiDong).Find(&hdcc).Error
	if err != nil {
		return nil, err
	}
	for _, cc := range hdcc {
		obj := models.CanCu{}
		obj, err = GetCanCuByID(cc.MaCanCu)
		if err != nil {
			return nil, err
		}
		rs = append(rs, obj)
	}
	return rs, nil
}
func GetListDHTDByHoiDong(maHoiDong string) ([]models.DanhHieuThiDua, error) {
	con := driver.Connect()
	defer con.Close()
	listkq := []models.DanhHieuThiDua{}
	listdhtd, err := GetAllDanhHieu()
	if err != nil {
		return nil, err
	}
	for _, elem := range listdhtd {
		tt := []models.XetThiDuaTapThe{}
		cn := []models.XetThiDuaCaNhan{}
		err = con.Where("ma_hoi_dong=? AND ma_danh_hieu=?", maHoiDong, elem.MaDanhHieu).Find(&tt).Error
		if err != nil {
			return nil, err
		}
		err = con.Where("ma_hoi_dong=? AND ma_danh_hieu=?", maHoiDong, elem.MaDanhHieu).Find(&cn).Error
		if err != nil {
			return nil, err
		}
		if len(tt) != 0 || len(cn) != 0 {
			listkq = append(listkq, elem)
		}
	}
	return listkq, nil
}
func GetListDanhHieuThiDuaTTByHoiDong(maHoiDong string) ([]models.DanhHieuThiDua, error) {
	con := driver.Connect()
	defer con.Close()
	list := []models.DanhHieuThiDua{}
	tt := []models.XetThiDuaTapThe{}
	err := con.Where("ma_hoi_dong=?", maHoiDong).Find(&tt).Error
	if err != nil {
		return nil, err
	}
	for _, elem := range tt {
		dh := models.DanhHieuThiDua{}
		err = con.Find(&dh, elem.MaDanhHieu).Error
		if err != nil {
			return nil, err
		}
		list = append(list, dh)
	}
	return list, nil
}
func MoHoiDong(maHoiDong string) error {
	myTime := time.Now().Unix() * 1000
	con := driver.Connect()
	defer con.Close()
	hd := models.HoiDong{}
	err := con.Where("ma_hoi_dong=?", maHoiDong).Find(&hd).Error
	if err != nil {
		return err
	}
	hd.TrangThai = 2 //Mở  hội đồng
	hd.ThoiGianMo = myTime
	err = con.Save(&hd).Error
	if err != nil {
		return err
	}

	fmt.Printf("Hoi dong %s da mo &v", hd.MaHoiDong, time.Now())

	listThanhVienHoiDong, err := GetListThanhVienByHoiDong(hd.MaHoiDong)
	if err != nil {
		return err
	}
	//update trạng thái của các cán bộ thuộc thành viên hội đồng
	for _, value := range listThanhVienHoiDong {
		if value.ChucDanh.ThamGiaBoPhieu == true {
			canbo, err := GetCanBoByID(value.CanBo.MaCanBo)
			if err != nil {
				return err
			}
			canbo.MaQuyen = 3 //Thành viên hội đồng
			err = con.Save(&canbo).Error
			if err != nil {
				return err
			}
		}
	}
	//ctab := crontab.New()
	//ctab.MustAddJob("* * * * *", func() {
	//	con := driver.Connect()
	//	defer con.Close()
	//	hd := models.HoiDong{}
	//	err := con.Last(&hd).Error
	//	if err != nil {
	//		return
	//	}
	//	timeNow := int64(time.Now().Unix() * 1000)
	//	if timeNow >= hd.ThoiGianMo+24000000{
	//		if timeNow >= hd.ThoiGianMo+30000000 {
	//			hd.TrangThai = 1
	//			err = con.Save(&hd).Error
	//			if err != nil {
	//				return
	//			}
	//			fmt.Printf("Hoi dong %s da dong in %v", hd.MaHoiDong, time.Now())
	//			jsonMessage, _ := json.Marshal(&utils.Message{Content: "Hội đồng "+hd.MaHoiDong+" đã đóng."})
	//			utils.Manager.Broadcast <- jsonMessage
	//
	//			listThanhVienHoiDong, err := GetListThanhVienByHoiDong(hd.MaHoiDong)
	//			if err != nil {
	//				return
	//			}
	//			//update trạng thái của các cán bộ thuộc thành viên hội đồng
	//			for _, value := range listThanhVienHoiDong {
	//				if value.ChucDanh.ThamGiaBoPhieu == true {
	//					canbo, err := GetCanBoByID(value.CanBo.MaCanBo)
	//					if err != nil {
	//						return
	//					}
	//					canbo.MaQuyen = 1 //Thành viên hội đồng
	//					err = con.Save(&canbo).Error
	//					if err != nil {
	//						return
	//					}
	//				}
	//			}
	//			ctab.Clear()
	//		}else{
	//			jsonMessage, _ := json.Marshal(&utils.Message{Content: "Hội đồng bầu chọn sẽ đóng sau 1 phút nữa!"})
	//			utils.Manager.Broadcast <- jsonMessage
	//		}
	//	}
	//
	//})
	return nil
}
func DongHoiDong(maHoiDong string) error {
	con := driver.Connect()
	defer con.Close()
	hd := models.HoiDong{}
	err := con.Where("ma_hoi_dong=?", maHoiDong).Find(&hd).Error
	if err != nil {
		return err
	}
	hd.TrangThai = 1 //Đóng  hội đồng
	err = con.Save(&hd).Error
	if err != nil {
		return err
	}
	listThanhVienHoiDong, err := GetListThanhVienByHoiDong(maHoiDong)
	if err != nil {
		return err
	}
	//update trạng thái của các cán bộ thuộc thành viên hội đồng
	for _, value := range listThanhVienHoiDong {
		if value.ChucDanh.ThamGiaBoPhieu == true {
			canbo, err := GetCanBoByID(value.CanBo.MaCanBo)
			if err != nil {
				return err
			}
			canbo.MaQuyen = 1 //Thành viên hội đồng
			err = con.Save(&canbo).Error
			if err != nil {
				return err
			}
		}
	}
	listdh, err := GetAllDanhHieu()
	for _, value := range listdh {
		value.TrangThai = false
		value.DaBauChon = 0
		value.ThoiGianMo = 0
		value.SoPhut = 0
		value.SlBoPhieu = 0
		err = con.Save(&value).Error
		if err != nil {
			return err
		}
	}
	listkt, err := GetAllHinhThucKhenThuong()
	for _, value := range listkt {
		value.TrangThai = false
		value.DaBauChon = 0
		value.ThoiGianMo = 0
		value.SoPhut = 0
		err = con.Save(&value).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func ActiveDanhHieu(maDanhHieu int) (models.DanhHieuThiDua, error) {
	con := driver.Connect()
	defer con.Close()
	dh := models.DanhHieuThiDua{}
	err := con.Find(&dh, maDanhHieu).Error
	if err != nil {
		return models.DanhHieuThiDua{}, err
	}
	dh.TrangThai = true
	dh.DaBauChon = 1
	timenow := int64(time.Now().Unix() * 1000)
	dh.ThoiGianMo = timenow
	err = con.Save(&dh).Error
	if err != nil {
		return models.DanhHieuThiDua{}, err
	}
	listtp := []models.ThanhPhanHoiDong{}
	err = con.Where("bau_chon = true").Find(&listtp).Error
	if err != nil {
		return models.DanhHieuThiDua{}, err
	}
	for _, elem := range listtp {
		elem.BauChon = false
		err = con.Save(&elem).Error
		if err != nil {
			return models.DanhHieuThiDua{}, err
		}
	}
	return dh, nil
}
func DeactiveDanhHieu(maDanhHieu int) (models.DanhHieuThiDua, error) {
	con := driver.Connect()
	defer con.Close()
	dh := models.DanhHieuThiDua{}
	err := con.Find(&dh, maDanhHieu).Error
	if err != nil {
		return models.DanhHieuThiDua{}, err
	}
	dh.TrangThai = false
	dh.DaBauChon = 2
	listtp := []models.ThanhPhanHoiDong{}
	err = con.Where("bau_chon = true").Find(&listtp).Error
	if err != nil {
		return models.DanhHieuThiDua{}, err
	}
	dh.SlBoPhieu = len(listtp)
	err = con.Save(&dh).Error
	if err != nil {
		return models.DanhHieuThiDua{}, err
	}
	return models.DanhHieuThiDua{}, nil
}

func CapNhatDiemDanh(maThanhPhan int, coMat bool) error {
	con := driver.Connect()
	defer con.Close()
	tp := models.ThanhPhanHoiDong{}
	err := con.Find(&tp, maThanhPhan).Error
	if err != nil {
		return err
	}
	tp.CoMat = coMat
	err = con.Save(&tp).Error
	if err != nil {
		return err
	}
	return nil
}
func CapNhatThoiGianBauChon(maDanhHieu int, soPhut int) error {
	con := driver.Connect()
	defer con.Close()
	dh := models.DanhHieuThiDua{}
	err := con.Find(&dh, maDanhHieu).Error
	if err != nil {
		return err
	}
	dh.SoPhut = soPhut
	err = con.Save(&dh).Error
	if err != nil {
		return err
	}
	return nil
}
