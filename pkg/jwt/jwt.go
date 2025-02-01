package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	TokenExpireDuration = time.Hour * 2
)

var mySecret = []byte("呜噜噜猪")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

// // GenToken 生成JWT
//
//	func GenToken(userID int64, username string) (string, error) {
//		// 创建一个我们自己的声明
//		claims := MyClaims{
//			userID, // 自定义字段
//			username,
//			jwt.RegisteredClaims{
//				ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
//				Issuer:    "blueball", // 签发人
//			},
//		}
//		// 使用指定的签名方法创建签名对象
//		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//		// 使用指定的secret签名并获得完整的编码后的字符串token
//		return token.SignedString(mySecret)
//	}
//

// GenToken ⽣生成access token 和 refresh token
func GenToken(userID int64, username string) (aToken, rToken string, err error) {
	// 创建⼀一个我们⾃自⼰己的声明
	c := MyClaims{
		userID, // ⾃自定义字段
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), // 过期时间
			Issuer:    "bluebell",
		}}
	// 加密并获得完整的编码后的字符串串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
	// 签发⼈人
	// refresh token 不不需要存任何⾃自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 30)), // 过期时间
		Issuer:    "bluebell",
	}).SignedString(mySecret)
	// 使⽤用指定的secret签名并获得完整的编码后的字符串串token
	return
	// 签发⼈人
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	mc := new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token⽆无效直接返回
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}
	// 从旧access token中解析出claims数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 当access token是过期错误 并且 refresh token没有过期时就创建⼀一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID, claims.Username)
	}
	return
}
