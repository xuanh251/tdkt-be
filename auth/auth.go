package auth

import (
	"TDKTServer/config"
	"TDKTServer/driver"
	"TDKTServer/models"
	"TDKTServer/repository"
	"encoding/json"

	//"TDKTServer/repository"
	"TDKTServer/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	ErrInvalidPassword = errors.New("Password invalid!")
)

type Auth struct {
	CanBo   models.CanBo `json:"can_bo"`
	Token   string       `json:"token"`
	IsValid bool         `json:"is_valid"`
	Message string       `json:"message"`
}

var configs = config.LoadConfigs()

func SignIn(canbo models.CanBo) (Auth, error) {
	con := driver.Connect()
	defer con.Close()
	password := canbo.Password
	canbo, err := repository.GetCanBoByEmail(canbo.Email)
	if canbo.TrangThai {
		return Auth{IsValid: false, Message: "isUsing"}, nil
	}
	if err != nil {
		return Auth{IsValid: false}, err
	}
	err = utils.IsPassword(canbo.Password, password)
	if err != nil {
		return Auth{IsValid: false}, ErrInvalidPassword
	}
	token, err := GenerateJWT(canbo)
	if err != nil {
		return Auth{IsValid: false}, err
	}
	canbo.TrangThai = true
	err = con.Save(&canbo).Error
	if err != nil {
		return Auth{IsValid: false}, err
	}
	jsonMessage, _ := json.Marshal(&utils.Message{Content: "Cán bộ " + canbo.HoTen + " đã online!"})
	utils.Manager.Broadcast <- jsonMessage
	return Auth{canbo, token, true, "Đăng nhập thành công"}, nil
}
func GenerateJWT(canbo models.CanBo) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["info"] = canbo
	claims["exp"] = time.Now().Add(time.Minute * 60 * 24).Unix()
	return token.SignedString(configs.Jwt.SecretKey)
}
