package controllers

import (
	"TDKTServer/auth"
	"TDKTServer/models"
	"TDKTServer/repository"
	"TDKTServer/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var canbo models.CanBo
	err := json.Unmarshal(body, &canbo)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnauthorized)
		return
	}
	userAuthenticate, err := auth.SignIn(canbo)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnauthorized)
		return
	}
	//if userAuthenticate.Message == "isUsing" {
	//	utils.ToJson(w, userAuthenticate)
	//	return
	//}
	utils.ToJson(w, userAuthenticate)
}

func CreateCanBo(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var canbo models.CanBo
	err := json.Unmarshal(body, &canbo)
	if err != nil {
		fmt.Println(err)
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = repository.NewCanBo(canbo)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, "Thêm cán bộ thành công!")
}

func GetAllCanBos(w http.ResponseWriter, r *http.Request) {
	CanBos, err := repository.GetAllCanBo()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, CanBos)
}
func GetCanBoById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	CanBo, err := repository.GetCanBoByID(uint32(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, CanBo)
}
func UpdateCanBo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	body, _ := ioutil.ReadAll(r.Body)
	var CanBo models.CanBo
	err := json.Unmarshal(body, &CanBo)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	CanBo.MaCanBo = uint32(id)
	_, err = repository.UpdateCanBo(CanBo)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, "Cập nhật thành công!")
}
func DeleteCanBo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	_, err := repository.DeleteCanBo(uint32(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, "Xoá dữ liệu thành công!")
}
func SetActive(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	err := repository.SetActive(uint32(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, "Active")
}
func SetDeactive(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	err := repository.SetDeactive(uint32(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, "Deactive")
}
