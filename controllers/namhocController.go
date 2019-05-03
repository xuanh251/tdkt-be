package controllers

import (
	"TDKTServer/models"
	"TDKTServer/repository"
	"TDKTServer/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateNamHoc(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var namhoc models.NamHoc
	err := json.Unmarshal(body, &namhoc)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = repository.CreateNamHoc(namhoc)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.ToJson(w, "Tạo năm học thành công!")
}
func GetAllNamHoc(w http.ResponseWriter, r *http.Request) {
	list_donvi, err := repository.GetAllNamHoc()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, list_donvi)
}

//func GetNamHocById(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	id, _ := strconv.ParseInt(params["id"], 10, 12)
//	namhoc, err := repository.GetChucDanhByID(uint32(id))
//	if err != nil {
//		utils.ErrorResponse(w, err, http.StatusBadRequest)
//		return
//	}
//	utils.ToJson(w, namhoc)
//}
