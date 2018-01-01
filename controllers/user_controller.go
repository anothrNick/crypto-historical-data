package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/anothrnick/generic-serviceapi-go/db"
  	"github.com/anothrnick/generic-serviceapi-go/models"
  	"github.com/anothrnick/generic-serviceapi-go/utils"
	"net/http"
	"time"
)

type LoginCreds struct {
	Email 		string 	`json:"email"`
	Password 	string 	`json:"password"`
}

func UserTokenClaims(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
			"user_id": c.MustGet("user_id"),
			"name": c.MustGet("name"),
			"email": c.MustGet("email"),
			"active": c.MustGet("active"),
			"exp": c.MustGet("exp"),
		})
	return
}

func LoginUser(c *gin.Context) {
	var user models.User
	var creds LoginCreds

	// TODO: basic authn
	c.BindJSON(&creds)

	if creds.Email == "" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	db := db.Database()
    defer db.Close()
	if err := db.Where("email = ?", creds.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	if (!utils.CheckPasswordHash(creds.Password, user.Password)) {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	// create jwt
	tokenString, err := utils.GetToken(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating token."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
			"message": "Successfully Logged in.", 
			"authorization_token": tokenString, 
			"user_id": user.ID, 
			"email": user.Email, 
			"name": user.Name,
			"active": user.Active,
			"joined": user.Joined,
		})
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	user.Password, _ = utils.HashPassword(user.Password)
	user.Joined = time.Now()

	db := db.Database()
    defer db.Close()
	db.Save(&user)

	tokenString, err := utils.GetToken(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating token."})
		return
	}

	// _, err = utils.SendConfirmation(user.Email, user.Name)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "Error sending confirmation email."})
	// 	return
	// }

	c.JSON(http.StatusCreated, gin.H{
			"message" : "User created successfully!", 
			"authorization_token": tokenString, 
			"user_id": user.ID, 
			"email": user.Email, 
			"name": user.Name,
			"active": user.Active,
			"joined": user.Joined,
		})
}

func GetSingleUser(c *gin.Context) {
	var user models.User
	var _users []models.TransformedUser
	user_id := c.Param("user_id")

	db := db.Database()
    defer db.Close()
	if err := db.Where("id = ?", user_id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	_users = append(_users, models.TransformedUser{ID: user.ID, Email: user.Email, Name: user.Name, Active: user.Active, Joined: user.Joined})
	c.JSON(http.StatusOK, gin.H{"data": _users})
}

// admin...
func GetAllUsers(c *gin.Context) {
	var users []models.User
	_users := make([]models.TransformedUser, 0)

	db := db.Database()
    defer db.Close()
	db.Find(&users)

	if (len(users) <= 0) {
		c.JSON(http.StatusNotFound, gin.H{"message": "No Users."})
		return
	}

	for _, user := range users {
		_users = append(
			_users, 
			models.TransformedUser{
				ID: user.ID, 
				Email: user.Email, 
				Name: user.Name, 
				Active: user.Active,
				Joined: user.Joined,
				},
				)
	}

	c.JSON(http.StatusOK, gin.H{"data": _users})
}