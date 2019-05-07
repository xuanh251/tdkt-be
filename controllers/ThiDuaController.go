package controllers

import (
	"TDKTServer/repository"
	"TDKTServer/utils"
	"net/http"
)

func NapDanhSachTapTheThiDua(w http.ResponseWriter, r *http.Request) {
	listtt, err := repository.NapDanhSachTapTheTD()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, listtt)
}
func NapDanhSachCaNhanThiDua(w http.ResponseWriter, r *http.Request) {
	listcn, err := repository.NapDanhSachCaNhanTD()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, listcn)
}
