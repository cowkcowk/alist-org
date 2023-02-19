package server

import (
	"github.com/alist-org/alist/v3/internal/conf"
	"github.com/alist-org/alist/v3/server/common"
	"github.com/gin-gonic/gin"
)

func Init(e gin.Engine) {
	g := e.Group(conf.URL.Path)
	common.SecretKey = []byte(conf.Conf.JwtSecret)
	
}