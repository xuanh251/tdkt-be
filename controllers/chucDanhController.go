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

func CreateChucDanh(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var chucdanh models.ChucDanh
	err := json.Unmarshal(body, &chucdanh)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = repository.CreateChucDanh(chucdanh)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, "Tạo chức vụ thành công!")
}
func GetAllChucDanh(w http.ResponseWriter, r *http.Request) {
	list_chucdanh, err := repository.GetAllChucDanh()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, list_chucdanh)
}
func GetChucDanhById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	chucdanh, err := repository.GetChucDanhByID(uint32(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, chucdanh)
}
func UpdateChucDanh(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ma_chuc_danh, _ := strconv.ParseInt(params["id"], 10, 12)
	body, _ := ioutil.ReadAll(r.Body)
	var chucdanh models.ChucDanh
	err := json.Unmarshal(body, &chucdanh)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	chucdanh.MaChucDanh = int8(ma_chuc_danh)
	_, err = repository.UpdateChucDanh(chucdanh)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, "Cập nhật thành công!")
}
func DeleteChucDanh(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ma_chuc_danh, _ := strconv.ParseInt(params["id"], 10, 12)
	_, err := repository.DeleteChucDanh(uint32(ma_chuc_danh))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{"Xoá chức vụ thành công!", http.StatusAccepted})
}
