package controller

import (
	"btpn-backend-go/config"
	"btpn-backend-go/database"
	"btpn-backend-go/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()

	contentType := config.GetContentType(c)
	_, _ = db, contentType
	User := model.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"email":    User.Email,
		"username": User.Username,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := config.GetContentType(c)
	_, _ = db, contentType
	User := model.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email / password",
		})

		return
	}

	comparePass := config.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email / password",
		})

		return
	}

	token := config.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := config.GetContentType(c)
	_, _ = db, contentType

	User := model.User{}
	NewUser := model.User{}

	id := c.Param("userId")

	err := db.First(&User, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&NewUser)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = db.Model(&User).Updates(NewUser).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         User.ID,
		"email":      User.Email,
		"username":   User.Username,
		"updated_at": User.UpdatedAt,
	})
}

func UserDelete(c *gin.Context) {
	db := database.GetDB()
	contentType := config.GetContentType(c)
	_, _ = db, contentType
	User := model.User{}

	id := c.Param("userId")

	err := db.First(&User, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = db.Delete(&User).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
