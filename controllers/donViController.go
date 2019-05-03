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

func CreateDonVi(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var donvi models.DonVi
	err := json.Unmarshal(body, &donvi)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = repository.CreateDonVi(donvi)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, "Tạo đơn vị thành công!")
}
func GetAllDonVi(w http.ResponseWriter, r *http.Request) {
	list_donvi, err := repository.GetAllDonVi()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, list_donvi)
}
func GetDonViById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	donvi, err := repository.GetDonViByID(uint32(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, donvi)
}
func UpdateDonVi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ma_don_vi, _ := strconv.ParseInt(params["id"], 10, 12)
	body, _ := ioutil.ReadAll(r.Body)
	var donvi models.DonVi
	err := json.Unmarshal(body, &donvi)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	donvi.MaDonVi = uint32(ma_don_vi)
	_, err = repository.UpdateDonVi(donvi)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, "Cập nhật thành công!")
}
func DeleteDonVi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ma_don_vi, _ := strconv.ParseInt(params["id"], 10, 12)
	_, err := repository.DeleteDonVi(uint32(ma_don_vi))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{"Xoá đơn vị thành công!", http.StatusAccepted})
}
