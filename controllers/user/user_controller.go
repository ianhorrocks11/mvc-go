package userController

import (
	"mvc-go/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserById(c *gin.Context) {
	log.Debug("User id: " + c.Param("id")) //recibe un parametro que es el id

	var userDto dto.UserDto
	userDto.Name = "Edu"
	userDto.LastName = "Gaite"

	c.JSON(http.StatusOK, userDto) // Devuelve un json con status ok
}

func GetUsers(c *gin.Context) {

	c.JSON(http.StatusOK, "Users:")
}

func OrderInsert(c *gin.Context) {
	var orderDto dto.OrderDto
}
func UserInsert(c *gin.Context) {
	var userDto dto.UserDto //marshall
	err := c.BindJSON(&userDto)

	log.Debug(userDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, userDto)
}
/*
debug
info
error
fatal
crash*/ 