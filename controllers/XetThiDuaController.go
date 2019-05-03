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

func CapNhatXetThiDuaTapThe(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	data := models.XetThiDuaTapThe{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = repository.CapNhatXetThiDuaTapThe(data)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, "Đã cập nhật tập thể xét thi đua!")
}
func GetDanhSach_TDTT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	maHoiDong := params["id"]
	list, err := repository.GetDanhSach_TDTT(maHoiDong)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, list)
}
func GetObjById_TDTT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	obj, err := repository.GetObjById_TDTT(int8(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, obj)
}

func DeleteObj_TDTT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	_, err := repository.DeleteObj_TDTT(int8(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{"Xoá dữ liệu thành công!", http.StatusAccepted})
}

//-------------------------------------------------
//cá nhân
func CapNhatXetThiDuaCaNhan(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	data := models.XetThiDuaCaNhan{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = repository.CapNhatXetThiDuaCaNhan(data)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, "Đã cập nhật tập thể xét thi đua!")
}
func GetDanhSach_TDCN(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	maHoiDong := params["id"]
	list, err := repository.GetDanhSach_TDCN(maHoiDong)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, list)
}
func GetObjById_TDCN(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	obj, err := repository.GetObjById_TDCN(int8(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, obj)
}

func DeleteObj_TDCN(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	_, err := repository.DeleteObj_TDCN(int8(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{"Xoá dữ liệu thành công!", http.StatusAccepted})
}
