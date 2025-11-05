package handlers

import (
	"base_crud_api/internals/models"
	"base_crud_api/internals/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service services.UserServiceInterface
}

// Prevalidation using validator

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (uh *UserHandler) SignUpUser(r *gin.Context) {

	var body models.LoginRequest

	if err := r.ShouldBindJSON(&body); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "request body is missing paramters or empty"})
		return

	}

	service, err := uh.Service.CreateUser(body)

	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	r.JSON(http.StatusOK, gin.H{"username": string(service)})

}

// comiong back with an EOF for some reason when there is nothing passed
func (uh *UserHandler) LoginUser(c *gin.Context) {

	// if c.Request.Body == nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is empty"})
	// 	return
	// }

	var body models.LoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is missing paramters or empty"})
		// fmt.Println("error from handler itself", err)
		return
	}

	token, err := uh.Service.LoginUser(body)
	if err != nil {

		// The thing with this is that we need to be able to filter out the different kinds of errors that can come back from the service layer here
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
