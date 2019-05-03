package controllers

import (
	"TDKTServer/models"
	"TDKTServer/repository"
	"TDKTServer/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetListBauChonThiDuaTapTheByDanhHieu(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	m := make(map[string]interface{})
	err := json.Unmarshal(body, &m)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	maDanhHieu := int8(m["maDanhHieu"].(float64))
	maHoiDong := m["maHoiDong"].(string)
	tt := []models.XetThiDuaTapThe{}
	tt, err = repository.GetListBauChonThiDuaTapTheByDanhHieu(maDanhHieu, maHoiDong)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, tt)
}
