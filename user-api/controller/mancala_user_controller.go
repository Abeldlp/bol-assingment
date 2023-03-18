package controller

import (
	"net/http"

	"github.com/Abeldlp/bol-assignment/user-api/model"
	"github.com/gin-gonic/gin"
)

func GetMancalaUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.MancalaUser
		if err := model.GetUserById(&user, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong",
			})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func CreateMancalaUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.MancalaUser
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Could not serialize payload",
			})
			return
		}

		if err := model.CreateUser(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong",
			})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}
