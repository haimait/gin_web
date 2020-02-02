package service

import (
	"fmt"
	"github.com/haimait/gin_web/lib/jwt"
	"time"
)

func test()  {
	fmt.Println("server test_fun")

}
//CreateToken @desc //生成一个新的token
func CreateToken() string {
	// 判断用户信息是否补全

	claim := jwt.CustomClaims{
		DeviceUuid: "LastLoginDevice",
		UserId:     100000001,
		LinkBookId: "100000001",
		Password:  "123456",
		PlateForm:  "app",
		EndTime:    time.Now().Unix(),
	}
	// 返回token和用户的基本信息
	token, _ := jwt.NewJWT().CreateToken(claim)
	return token
}
