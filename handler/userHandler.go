package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"userAuthApp/helper"
	"userAuthApp/model/dto"
	"userAuthApp/service"
)

type UserHandler struct {
	UserService service.InterfaceUserService
}

// create function provider
func NewUserHandler(userService service.InterfaceUserService) *UserHandler {
	return &UserHandler{
		userService,
	}
}

// handler create user
func (u *UserHandler) CreateUser(c *gin.Context) {
	var request dto.CreateUserRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		var errorObject []helper.ErrorValidation
		errValidation, ok := err.(validator.ValidationErrors)
		if ok {
			for _, value := range errValidation {
				errorObject = append(errorObject, helper.ErrorValidation{
					Field:        value.Field(),
					ErrorMessage: helper.ErrorMessageFromTag(value.Tag()),
				})
			}

			var errorMessages []string
			for _, value := range errorObject {
				errorMessages = append(errorMessages, value.ErrorMessage)
			}

			response := dto.ApiMessage{
				StatusCode: http.StatusBadRequest,
				Status:     "bad request",
				Message:    strings.Join(errorMessages, ". "),
			}

			c.JSON(http.StatusBadRequest, &response)
			return
		}
	}

	// call procedure insert in service
	result, err := u.UserService.InsertUser(c, &request)
	if err != nil {
		response := dto.ApiMessage{
			StatusCode: http.StatusInternalServerError,
			Status:     "internal server error",
			Message:    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, &response)
		return
	}

	// success create new user
	response := dto.ApiMessage{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Message:    "success create new user",
		Data:       result,
	}

	c.JSON(http.StatusOK, &response)
}

func (u *UserHandler) GetUser(c *gin.Context) {
	// call procedure get details in service
	userDetail, err := u.UserService.Getuser(c)
	if err != nil {
		response := dto.ApiMessage{
			StatusCode: http.StatusNotFound,
			Status:     "not found",
			Message:    err.Error(),
		}

		c.JSON(http.StatusNotFound, &response)
		return
	}

	// success get data
	response := dto.ApiMessage{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Message:    "success get data",
		Data:       userDetail,
	}

	c.JSON(http.StatusOK, &response)
}
