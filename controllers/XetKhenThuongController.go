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

func CapNhatXetKhenThuongTapThe(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	data := models.XetKhenThuongTapThe{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = repository.CapNhatXetKhenThuongTapThe(data)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, "Đã cập nhật tập thể xét khen thưởng!")
}
func GetDanhSach_KTTT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	maHoiDong := params["id"]
	list, err := repository.GetDanhSach_KTTT(maHoiDong)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, list)
}
func GetObjById_KTTT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	obj, err := repository.GetObjById_KTTT(int8(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, obj)
}

func DeleteObj_KTTT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	_, err := repository.DeleteObj_KTTT(int8(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{"Xoá dữ liệu thành công!", http.StatusAccepted})
}

//-------------------------------------------------
//cá nhân
func CapNhatXetKhenThuongCaNhan(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	data := models.XetKhenThuongCaNhan{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = repository.CapNhatXetKhenThuongCaNhan(data)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, "Đã cập nhật tập thể xét khen thưởng!")
}
func GetDanhSach_KTCN(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	maHoiDong := params["id"]
	list, err := repository.GetDanhSach_KTCN(maHoiDong)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, list)
}
func GetObjById_KTCN(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	obj, err := repository.GetObjById_KTCN(int8(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, obj)
}

func DeleteObj_KTCN(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	_, err := repository.DeleteObj_KTCN(int8(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{"Xoá dữ liệu thành công!", http.StatusAccepted})
}
