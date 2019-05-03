package utils

import (
	"TDKTServer/models"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

type DefaultResponse struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

func ErrorResponse(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	ToJson(w, struct {
		Message string `json:"message"`
	}{Message: err.Error()})
}
func ToJson(w http.ResponseWriter, data interface{}) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
func Bcrypt(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
func IsPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
func ConvertString2Uint32(str string) uint32 {
	rs, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(rs)
}
func ConvertSchoolYear(data models.NamHoc) string {
	rs := ""
	year, _ := strconv.Atoi(data.TuNgay[:4])
	rs = strconv.Itoa(year) + " - " + strconv.Itoa(year+1)
	return rs
}
