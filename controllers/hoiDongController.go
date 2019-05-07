package controllers

import (
	"TDKTServer/models"
	"TDKTServer/repository"
	"TDKTServer/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetAllHoiDong(w http.ResponseWriter, r *http.Request) {
	listHoiDong, err := repository.GetAllHoiDong()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, listHoiDong)
}
func CreateHoiDong(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var hoidong models.HoiDong
	err := json.Unmarshal(body, &hoidong)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	MaHoiDong, err := repository.CreateHoiDong(hoidong)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	namhoc, err := repository.GetNamHocByID(hoidong.MaNamHoc)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	namhoc.TuNgay = utils.ConvertSchoolYear(namhoc)
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, struct {
		MaHoiDong string `json:"ma_hoi_dong"`
		NamHoc    string `json:"nam_hoc"`
	}{MaHoiDong: MaHoiDong,
		NamHoc: namhoc.TuNgay})
}
func CapNhatCanCuHoiDong(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	m := make(map[string]interface{})
	err := json.Unmarshal(body, &m)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	listIdCC := []int64{}

	dataIdHoiDong := m["idHoiDong"].(string)
	dataIdccc := m["listCCC"].([]interface{})
	dataccr := m["listCCR"].([]interface{})

	//thêm id vào list ccc
	for _, value := range dataIdccc {
		listIdCC = append(listIdCC, int64(value.(float64)))
	}
	//Thêm căn cứ riêng vào bảng căn cứ-->trả về id căn cứ-->thêm id vào list ccr
	for _, value := range dataccr {
		objcc := models.CanCu{}
		jsonString, _ := json.Marshal(value.(map[string]interface{}))
		json.Unmarshal(jsonString, &objcc)
		idcc, _ := repository.CreateCanCu(objcc)
		listIdCC = append(listIdCC, idcc)
	}
	//thêm các list id vào bảng hội đồng căn cứ
	for _, value := range listIdCC {
		obj := models.HoiDongCanCu{}
		obj.MaHoiDong = dataIdHoiDong
		obj.MaCanCu = value
		_, err := repository.CreateHoiDongCanCuRelate(obj)
		if err != nil {
			utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, "Cập nhật danh sách căn cứ thành công!")
}
func CapNhatThanhVienHoiDong(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	m := make(map[string]interface{})
	err := json.Unmarshal(body, &m)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	obj := models.ThanhPhanHoiDong{}
	jsonString, _ := json.Marshal(m)
	json.Unmarshal(jsonString, &obj)
	_, err = repository.CreateThanhPhanHoiDong(obj)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, "Đã cập nhật thành viên hội đồng!")
}
func GetListThanhVienByHoiDong(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	m := make(map[string]interface{})
	err := json.Unmarshal(body, &m)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	rs := []models.ThanhPhanHoiDong{}
	maHoiDong := m["ma_hoi_dong"].(string)
	rs, err = repository.GetListThanhVienByHoiDong(maHoiDong)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rs)
}
func GetListCanCuByHoiDong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	list, err := repository.GetListCanCuByHoiDong(params["id"])
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, list)
}

func GetHoiDongById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hoidong, err := repository.GetHoiDongById(params["id"])
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, hoidong)
}
func GetListThanhVien(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	list, err := repository.GetListThanhVienByHoiDong(params["id"])
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, list)
}
func GetLastHoiDong(w http.ResponseWriter, r *http.Request) {
	hoiDong, err := repository.GetLastHoiDong()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, hoiDong)
}
func GetListDanhHieuThiDuaByHoiDong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	list, err := repository.GetListDanhHieuThiDuaTTByHoiDong(params["id"])
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, list)
}
func MoHoiDong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := repository.MoHoiDong(params["id"])
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	} else {
		utils.ToJson(w, "Hội đồng đã mở!")
	}
}
func DongHoiDong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := repository.DongHoiDong(params["id"])
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	} else {
		utils.ToJson(w, "Hội đồng đã đóng!")
	}
}

func GetListDHTDByHoiDong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	listkq, err := repository.GetListDHTDByHoiDong(params["id"])
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, listkq)
}
func ActiveDanhHieu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	dh, err := repository.ActiveDanhHieu(id)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, dh)
}
func DeactiveDanhHieu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	dh, err := repository.DeactiveDanhHieu(id)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, dh)
}

func CapNhatDiemDanh(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	m := make(map[string]interface{})
	err := json.Unmarshal(body, &m)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	list := m["request"].([]interface{})
	for _, elem := range list {
		maThanhPhan := elem.(map[string]interface{})["ma_thanh_phan"]
		coMat := elem.(map[string]interface{})["co_mat"]
		err = repository.CapNhatDiemDanh(int(maThanhPhan.(float64)), coMat.(bool))
		if err != nil {
			utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
			return
		}
	}
	utils.ToJson(w, "Đã cập nhật trạng thái các thành viên hội đồng!")
}

func CapNhatThoiGianBauChon(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	m := make(map[string]interface{})
	err := json.Unmarshal(body, &m)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	maDanhHieu := m["ma_danh_hieu"].(float64)
	soPhut := m["so_phut"].(float64)
	err = repository.CapNhatThoiGianBauChon(int(maDanhHieu), int(soPhut))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, true)
}
