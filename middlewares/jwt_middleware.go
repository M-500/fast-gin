package middlewares

import (
	"errors"
	"fast-gin/app/config"
	"fast-gin/pkg/comm/constant"
	"fast-gin/pkg/comm/e"
	"fast-gin/pkg/comm/response"
	"fast-gin/pkg/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2023/12/5 11:33
type JWT struct {
	SigningKey []byte
}

// CustomClaims
// @Description: Jwt相关数据
type CustomClaims struct {
	ID          uint
	NickName    string
	UserName    string
	AuthorityId uint
	jwt.StandardClaims
}

func NewJwt() *JWT {
	jwtCfg := config.ConfigInstance.Jwt
	return &JWT{
		[]byte(jwtCfg.SecretKey), // 可以设置过期时间
	}
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(e.GetMessage(e.TokenMalformed))
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errors.New(e.GetMessage(e.TokenExpired))
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(e.GetMessage(e.TokenNotValidYet))
			} else {
				return nil, errors.New(e.GetMessage(e.TokenInvalid))
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, errors.New(e.GetMessage(e.TokenInvalid))
}

func JwtAuth() gin.HandlerFunc {
	jwtCfg := config.ConfigInstance.Jwt
	return func(c *gin.Context) {
		// 从请求头中获取token
		token := c.Request.Header.Get(jwtCfg.JwtHeaderKey)
		//token := c.Request.Header.Get("")
		// 从url参数中获取token
		fmt.Println(token)
		// 从body中获取token
		if utils.IsBlank(token) {
			// 直接报错用户未登录
			response.JsonFailCode(c, e.NOT_LOGIN)
			c.Abort()
			return
		}
		j := NewJwt()
		// 解析token信息
		claims, err := j.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    e.ERROR,
				"message": err.Error(),
				"success": false,
			})
			c.Abort()
			return
		}
		// 将用户信息插入上下文中
		c.Set("claims", claims)
		c.Set(constant.JWT_INFO_KEY, claims.ID)
		c.Next()
	}
}

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
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", errors.New(e.GetMessage(e.TokenInvalid))
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
