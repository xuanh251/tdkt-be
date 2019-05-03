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

func CreateHinhThucKhenThuong(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var hinhthuc models.HinhThucKhenThuong
	err := json.Unmarshal(body, &hinhthuc)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = repository.CreateHinhThucKhenThuong(hinhthuc)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, "Tạo hình thức khen thưởng thành công!")
}
func GetAllHinhThucKhenThuong(w http.ResponseWriter, r *http.Request) {
	list_hinhthuc, err := repository.GetAllHinhThucKhenThuong()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, list_hinhthuc)
}
func GetHinhThucKhenThuongById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	hinhthuc, err := repository.GetHinhThucKhenThuongByID(uint32(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, hinhthuc)
}
func UpdateHinhThucKhenThuong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ma_hinh_thuc, _ := strconv.ParseInt(params["id"], 10, 12)
	body, _ := ioutil.ReadAll(r.Body)
	var hinhthuc models.HinhThucKhenThuong
	err := json.Unmarshal(body, &hinhthuc)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	hinhthuc.MaHinhThuc = uint32(ma_hinh_thuc)
	_, err = repository.UpdateHinhThucKhenThuong(hinhthuc)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, "Cập nhật thành công!")
}
func DeleteHinhThucKhenThuong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ma_hinh_thuc, _ := strconv.ParseInt(params["id"], 10, 12)
	_, err := repository.DeleteHinhThucKhenThuong(uint32(ma_hinh_thuc))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{"Xoá hình thức khen thưởng thành công!", http.StatusAccepted})
}
