package middleware

import (
	"encoding/json"
	"log"
	"mithril/model"
	"mithril/service"

	// "os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "account"

type JwtAuthorizator func(data interface{}, c *gin.Context) bool

func GinJWTMiddlewareInit(jwtAuthorizator JwtAuthorizator) (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Minute * 5,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.UserCheck); ok {
				//get claims from username
				v.UserClaims = model.GetUserClaims(v.UserName)
				jsonClaim, _ := json.Marshal(v.UserClaims)
				//maps the claims in the JWT
				return jwt.MapClaims{
					"userName":   v.UserName,
					"userClaims": string(jsonClaim),
				}
			}
			return jwt.MapClaims{}
		},
		//PayloadFunc: func(data interface{}) jwt.MapClaims {
		//	if v, ok := data.(*model.User); ok {
		//		return jwt.MapClaims{
		//			identityKey: v.Account,
		//		}
		//	}
		//	return jwt.MapClaims{}
		//},
		// IdentityHandler: func(c *gin.Context) interface{} {
		// 	claims := jwt.ExtractClaims(c)
		// 	//extracts identity from claims
		// 	jsonClaim := claims["userClaims"].(string)
		// 	var userClaims []model.Claims
		// 	json.Unmarshal([]byte(jsonClaim), &userClaims)
		// 	//Set the identity
		// 	return &model.UserCheck{
		// 		UserName:   claims["userName"].(string),
		// 		UserClaims: userClaims,
		// 	}
		// },
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				Account: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			//handles the login logic. On success LoginResponse is called, on failure Unauthorized is called
			var loginVals service.UserLoginService
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			account := loginVals.Account
			password := loginVals.Password

			if model.CheckAuth(account, password) {
				return &model.User{
					Account: account,
				}, nil
			}
			//var service service.UserLoginService
			//if err := c.ShouldBind(&service); err == nil {
			//	res := service.Login(c)
			//	c.JSON(200, res)
			//		return &model.User{
			//			Account: service.Account,
			//		}, nil
			//}

			return nil, jwt.ErrFailedAuthentication
		},
		//receives identity and handles authorization logic
		Authorizator: jwtAuthorizator,
		//Authorizator: func(data interface{}, c *gin.Context) bool {
		//	if v, ok := data.(*model.User); ok && v != nil {
		//		return true
		//	}
		//	return false
		//},

		//handles unauthorized logic
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return
}

////role is admin can access
//func AdminAuthorizator(data interface{}, c *gin.Context) bool {
//	if v, ok := data.(*model.UserCheck); ok {
//		for _, itemClaim := range v.UserClaims {
//			if itemClaim.Type == "role" && itemClaim.Value == "admin" {
//				return true
//			}
//		}
//	}
//
//	return false
//}
//
////username is test can access
//func TestAuthorizator(data interface{}, c *gin.Context) bool {
//	if v, ok := data.(*model.UserCheck); ok && v.UserName == "test" {
//		return true
//	}
//
//	return false
//}

//
func AllUserAuthorizator(data interface{}, c *gin.Context) bool {
	return true
}

// //404 handler
// func NoRouteHandler(c *gin.Context) {
// 	code := e.PAGE_NOT_FOUND
// 	c.JSON(404, gin.H{"code": code, "message": e.GetMsg(code)})
// }
