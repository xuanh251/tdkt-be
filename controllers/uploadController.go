package controllers

//import (
//	"TDKTServer/repository"
//	"TDKTServer/utils"
//	"errors"
//	"fmt"
//	"net/http"
//)
//
//var (
//	ErrEmptyAccount = errors.New("Tài khoản không xác định!")
//)
//
//func UploadFile(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//
//
//	file, handle, err := r.FormFile("image")
//	userid := r.PostFormValue("uid");
//	if userid == "" {
//		utils.ErrorResponse(w, ErrEmptyAccount, http.StatusUnprocessableEntity)
//		return
//	}
//	uid := utils.ConvertString2Uint32(userid)
//
//
//	if err != nil {
//		fmt.Println("Error Retrieving file from form-data")
//		fmt.Println(err)
//		utils.ErrorResponse(w, err, http.StatusConflict)
//		return
//	}
//	defer file.Close()
//	_,err = repository.UploadUserPhoto(file, handle, uid)
//	if err!=nil {
//		utils.ErrorResponse(w, err, http.StatusConflict)
//		return
//	}
//	utils.ToJson(w, utils.DefaultResponse{"Upload thanh cong!",http.StatusOK})
//}
