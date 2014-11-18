package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	jwt "github.com/dgrijalva/jwt-go"

	"beeblog/configs"
)

// oprations for Jwt
type JwtController struct {
	beego.Controller
}

func (this *JwtController) URLMapping() {
	this.Mapping("IssueToken", this.IssueToken)
}

// @Title IssueToken
// @Description Issue a Json Web Token
// @Success 200 string
// @Failure 403 no privilege to access
// @Failure 500 server inner error
// @router /issue-token [get]
func (this *JwtController) IssueToken() {
	log := logs.NewLogger(10000)
	log.SetLogger("console", "")

	token := jwt.New(jwt.GetSigningMethod("RS256")) // Create a Token that will be signed with RSA 256.
	token.Claims["ID"] = "This is my super fake ID"
	token.Claims["exp"] = time.Now().Unix() + 36000
	// The claims object allows you to store information in the actual token.
	tokenString, _ := token.SignedString(configs.PrivateKey)
	// tokenString Contains the actual token you should share with your client.
	this.Data["json"] = map[string]string{"token": tokenString}

	//	log.Debug(err.Error())
	this.ServeJson()
}
