package routes

import (
	cont "github.com/Modifa/KO_TICKET_LOGGER/controller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	V1 := r.Group("/api/devfinder/")
	{
		V1.POST("TestDB", cont.TestDB)

	}
}
