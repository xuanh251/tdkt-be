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

func CreateChucVu(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var chucvu models.ChucVu
	err := json.Unmarshal(body, &chucvu)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = repository.CreateChucVu(chucvu)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, "Tạo chức vụ thành công!")
}
func GetAllChucVu(w http.ResponseWriter, r *http.Request) {
	list_chucvu, err := repository.GetAllChucVu()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, list_chucvu)
}
func GetChucVuById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	chucvu, err := repository.GetChucVuByID(uint32(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, chucvu)
}
func UpdateChucVu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ma_chuc_vu, _ := strconv.ParseInt(params["id"], 10, 12)
	body, _ := ioutil.ReadAll(r.Body)
	var chucvu models.ChucVu
	err := json.Unmarshal(body, &chucvu)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	chucvu.MaChucVu = uint32(ma_chuc_vu)
	_, err = repository.UpdateChucVu(chucvu)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, "Cập nhật thành công!")
}
func DeleteChucVu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ma_chuc_vu, _ := strconv.ParseInt(params["id"], 10, 12)
	_, err := repository.DeleteChucVu(uint32(ma_chuc_vu))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{"Xoá chức vụ thành công!", http.StatusAccepted})
}
