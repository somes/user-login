package controller

import (
	"encoding/hex"
	"log"
	"net/http"
	"server/internal/common"
	"server/internal/model"
	"server/internal/response"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

func PingController(ctx *gin.Context) {
	msg := service.PingService()
	ctx.JSON(200, gin.H{"msg": msg})
}

func IndexController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func LoginController(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if len(username) == 0 || len(password) == 0 {
		// ctx.JSON(http.StatusOK, gin.H{"code": -1, "msg": "username and password cannot be empty"})
		response.Response(ctx, http.StatusOK, -1, nil, "username and password cannot be empty")
		return
	}
	// log.Println(username, password)
	// log.Println(ctx.Request.TLS.PeerCertificates)

	DB := common.GetDB()
	// 判断 username 是否存在
	var user model.User
	DB.Where("username = ?", username).First(&user)
	if len(user.Username) == 0 {
		response.Response(ctx, http.StatusOK, -1, nil, "user does not exist")
		return
	}
	// 判断 password 是否正确
	if password != user.Password {
		response.Response(ctx, http.StatusOK, -1, nil, "incorrect password")
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": "login success"})
	response.Success(ctx, nil, "login success")
}

func LoginCertificateController(ctx *gin.Context) {
	cn := ctx.Request.TLS.PeerCertificates[0].Subject.CommonName
	serialNumber := ctx.Request.TLS.PeerCertificates[0].SerialNumber.String()
	hexSerialNumber := hex.EncodeToString(ctx.Request.TLS.PeerCertificates[0].SerialNumber.Bytes())
	log.Println(cn, serialNumber, hexSerialNumber)
}
