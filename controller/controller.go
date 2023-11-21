package controller

import (
	"fmt"
	"net/http"

	models "github.com/Modifa/KO_TICKET_LOGGER/models"
	services "github.com/Modifa/KO_TICKET_LOGGER/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func TestDB(c *gin.Context) {

	db := services.DB{}
	var u models.Employee_Reg

	var rb models.Returnblock
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	_, err := db.RegisterDeveloper("public.fn_register_employee", u)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		c.JSON(http.StatusBadRequest, rb.New(false, err.Error(), err))
		return
	} else {
		fmt.Println("QueryRow Succcess: ", "Reponse")
		c.JSON(http.StatusOK, rb.New(true, "Success", "Reponse"))
		return
	}
}
