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

func CreateCanCu(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var cancu models.CanCu
	err := json.Unmarshal(body, &cancu)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = repository.CreateCanCu(cancu)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, "Tạo căn cứ thành công!")
}
func GetAllCanCuChung(w http.ResponseWriter, r *http.Request) {
	list_cancu, err := repository.GetAllCanCuChung()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, list_cancu)
}
func GetAllCanCuRieng(w http.ResponseWriter, r *http.Request) {
	list_cancu, err := repository.GetCanCuRiengByHoiDong("123")
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, list_cancu)
}
func GetCanCuById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	cancu, err := repository.GetCanCuByID(id)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	cancu.NgayKy = cancu.NgayKy[0:10]
	utils.ToJson(w, cancu)
}
func UpdateCanCu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	body, _ := ioutil.ReadAll(r.Body)
	var cancu models.CanCu
	err := json.Unmarshal(body, &cancu)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	cancu.MaCanCu = id
	_, err = repository.UpdateCanCu(cancu)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, "Cập nhật thành công!")
}
func DeleteCanCu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	_, err := repository.DeleteCanCu(id)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{"Xoá căn cứ thành công!", http.StatusAccepted})
}
