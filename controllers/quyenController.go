package controllers

import (
	"TDKTServer/repository"
	"TDKTServer/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAllQuyen(w http.ResponseWriter, r *http.Request) {
	CanBos, err := repository.GetListQuyenLimit()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, CanBos)
}
func GetQuyenById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 12)
	chucdanh, err := repository.GetChucDanhByID(uint32(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, chucdanh)
}
