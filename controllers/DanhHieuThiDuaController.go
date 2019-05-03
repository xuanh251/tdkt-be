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

func CreateDanhHieuThiDua(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var danhhieu models.DanhHieuThiDua
	err := json.Unmarshal(body, &danhhieu)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = repository.CreateDanhHieu(danhhieu)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, "Tạo danh hiệu thi đua thành công!")
}
func GetAllDanhHieuThiDua(w http.ResponseWriter, r *http.Request) {
	list_danhhieu, err := repository.GetAllDanhHieu()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, list_danhhieu)
}
func GetDanhHieuThiDuaById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	danhhieu, err := repository.GetDanhHieuByID(uint32(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, danhhieu)
}
func UpdateDanhHieuThiDua(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ma_danh_hieu, _ := strconv.ParseInt(params["id"], 10, 12)
	body, _ := ioutil.ReadAll(r.Body)
	var danhhieu models.DanhHieuThiDua
	err := json.Unmarshal(body, &danhhieu)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	danhhieu.MaDanhHieu = uint32(ma_danh_hieu)
	_, err = repository.UpdateDanhHieu(danhhieu)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, "Cập nhật thành công!")
}
func DeleteDanhHieuThiDua(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ma_danh_hieu, _ := strconv.ParseInt(params["id"], 10, 12)
	_, err := repository.DeleteDanhHieu(uint32(ma_danh_hieu))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{"Xoá danh hiệu thi đua thành công!", http.StatusAccepted})
}
