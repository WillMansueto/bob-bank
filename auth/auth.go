package auth

import(
	"errors"
	"time"

	"bob-bank/models"
	"bob-bank/utils"
	"bob-bank/config"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	ErrInvalidPassword = errors.New("Senha Inválida")
)

type Auth struct {
	User    models.User `json:"user"`
	Token   string		`json:"token"`
	IsValid bool		`json:"is_valid"`
}

var configs = config.LoadConfigs()

func SignIn(user models.User) (Auth, error) {
	password := user.Password
	user, err := models.GetUserByEmail(user.Email)
	if err != nil {
		return Auth{IsValid:false}, err
	}
	err = utils.IsPassword(user.Password, password)
	if err != nil {
		return Auth{IsValid:false}, ErrInvalidPassword
	}
	token, err := GenerateJWT(user)
	if err != nil {
		return Auth{IsValid:false}, err
	}
	return Auth{user, token, true}, nil
}

func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized" : true,
		"userId" : user.UID,
		"exp" : time.Now().Add(time.Minute * 60).Unix(),
	})
	return token.SignedString(configs.Jwt.SecretKey)
}