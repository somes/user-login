package controller

import (
	"encoding/hex"
	"log"
	"net/http"
	"server/internal/common"
	"server/internal/dto"
	"server/internal/model"
	"server/internal/response"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

func PingController(ctx *gin.Context) {
	msg := service.PingService()
	ctx.JSON(200, gin.H{"msg": msg})
}

func LoginController(ctx *gin.Context) {
	// username := ctx.PostForm("username")
	// password := ctx.PostForm("password")
	var requestBody map[string]interface{}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		response.Response(ctx, http.StatusInternalServerError, -1, nil, err.Error())
		return
	}
	username := requestBody["username"].(string)
	password := requestBody["password"].(string)
	if len(username) == 0 || len(password) == 0 {
		// ctx.JSON(http.StatusOK, gin.H{"code": -1, "msg": "username and password cannot be empty"})
		response.Response(ctx, http.StatusOK, -1, nil, "username and password cannot be empty")
		return
	}
	log.Println(username, password)
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

	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, -1, nil, "internal server error")
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": "login success"})
	response.Success(ctx, gin.H{"token": token}, "login success")
}

func InfoController(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
}

func LoginCertificateController(ctx *gin.Context) {
	cn := ctx.Request.TLS.PeerCertificates[0].Subject.CommonName
	serialNumber := ctx.Request.TLS.PeerCertificates[0].SerialNumber.String()
	hexSerialNumber := hex.EncodeToString(ctx.Request.TLS.PeerCertificates[0].SerialNumber.Bytes())
	log.Println(cn, serialNumber, hexSerialNumber)
}
