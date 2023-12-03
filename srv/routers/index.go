package routers

import (
	ctrl "github.com/nocode-nolunch/kube-codejam-framework/srv/controllers"

	"github.com/gin-gonic/gin"
)

func Init(srv *gin.Engine) {

	srv.GET("/run-payload", ctrl.Index)

	srv.POST("/run-payload/register/google/oauth2", ctrl.RegisterGoogleOauth)

	srv.GET("/run-payload/run/google/oauth2", ctrl.RunGoogleOauth)

}
