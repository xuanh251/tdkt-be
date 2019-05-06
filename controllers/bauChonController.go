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
	tt := []*models.XetThiDuaTapThe{}
	tt, err = repository.GetListBauChonThiDuaTapTheByDanhHieu(maDanhHieu, maHoiDong)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, tt)
}
func GetListBauChonThiDuaCaNhanByDanhHieu(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	m := make(map[string]interface{})
	err := json.Unmarshal(body, &m)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	maDanhHieu := int8(m["maDanhHieu"].(float64))
	maHoiDong := m["maHoiDong"].(string)
	cn := []models.XetThiDuaCaNhan{}
	cn, err = repository.GetListBauChonThiDuaCaNhanByDanhHieu(maDanhHieu, maHoiDong)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, cn)
}

func BauChonThiDuaTapThe(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	m := make(map[string]interface{})
	err := json.Unmarshal(body, &m)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	data := m["Request"].([]interface{})
	for _, elem := range data {
		obj := models.BauChonThiDuaTapThe{}
		jsonString, _ := json.Marshal(elem.(map[string]interface{}))
		json.Unmarshal(jsonString, &obj)
		err = repository.AddBauChonThiDuaTapThe(obj)
		if err != nil {
			utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
			return
		}
	}
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, true)
}
func GetListDaBauChonTDTTByTVHD(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	listBC := []*models.BauChonThiDuaTapThe{}
	listBC, err = repository.GetListDaBauChonTDTTByTVHD(int8(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, listBC)
}
func AddListTDTTDatYeuCau(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	m := make(map[string]interface{})
	err := json.Unmarshal(body, &m)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	list := m["list"].([]interface{})
	for _, elem := range list {
		madv := elem.(map[string]interface{})["ma_don_vi"]
		ttqt := models.TapTheQuaTrinh{}
		ttqt.MaNamHoc = int8(m["ma_nam_hoc"].(float64))
		ttqt.MaHoiDong = m["ma_hoi_dong"].(string)
		ttqt.MaDonVi = int8(madv.(float64))
		ttqt.ListDanhHieu = append(ttqt.ListDanhHieu, int64(m["ma_danh_hieu"].(float64)))
		err = repository.AddBauChon2QuaTrinh(ttqt)
		if err != nil {
			utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
			return
		}
	}

	utils.ToJson(w, true)
}
