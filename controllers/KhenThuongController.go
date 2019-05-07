package controllers

import (
	"TDKTServer/repository"
	"TDKTServer/utils"
	"net/http"
)

func NapDanhSachTapTheKT(w http.ResponseWriter, r *http.Request) {
	listtt, err := repository.NapDanhSachTapTheKT()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, listtt)
}
func NapDanhSachCaNhanKT(w http.ResponseWriter, r *http.Request) {
	listcn, err := repository.NapDanhSachCaNhanKT()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, listcn)
}
