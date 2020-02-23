package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token")
	TokenOutTime     error  = errors.New("Token is out time")
	SignKey          string = "haimait"
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	DeviceUuid       string `json:"device_uuid"`
	UserId           uint  `json:"user_id"`
	LinkBookId       string `json:"link_book_id"`
	Password         string `json:"password"`
	EndTime          int64  `json:"end_time"`
	IsUpdateUserInfo int    `json:"is_update_user_info"`
	PlateForm        string `json:"plat_form"`
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}



// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	claims.EndTime = time.Now().Add(time.Hour * 24 * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	if strings.TrimSpace(tokenString) == "" {
		return nil, TokenMalformed
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if token == nil {
		return nil, TokenMalformed
	}

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// 验证token的过期时间
		if time.Now().Unix()-claims.EndTime >= 0 {
			return nil, TokenOutTime
		}
		fmt.Println(claims)
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.EndTime = time.Now().Add(time.Second).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
