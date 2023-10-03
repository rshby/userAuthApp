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

type AccountHandler struct {
	AccountService service.InterfaceAccountService
}

// create function provider
func NewAccountHandler(accService service.InterfaceAccountService) *AccountHandler {
	return &AccountHandler{
		AccountService: accService,
	}
}

// method insert / register
func (a *AccountHandler) Register(c *gin.Context) {
	var request dto.CreateAccountRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		msg, ok := err.(validator.ValidationErrors)
		errorObject := []helper.ErrorValidation{}
		if ok {
			for _, errMsg := range msg {
				errorObject = append(errorObject, helper.ErrorValidation{
					Field:        errMsg.Field(),
					ErrorMessage: helper.ErrorMessageFromTag(errMsg.Tag()),
				})
			}

			var errMessges []string
			for _, value := range errorObject {
				errMessges = append(errMessges, value.ErrorMessage)
			}

			responseAPI := &dto.ApiMessage{
				StatusCode: http.StatusBadRequest,
				Status:     "bad request",
				Message:    strings.Join(errMessges, ". "),
			}
			c.JSON(http.StatusBadRequest, responseAPI)
			return
		}
	}

	// call function to insert in service
	result, err := a.AccountService.InsertAccount(c, &request)
	if err != nil {
		response := &dto.ApiMessage{
			StatusCode: http.StatusInternalServerError,
			Status:     "internal server error",
			Message:    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// success insert
	response := &dto.ApiMessage{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Message:    "success create account",
		Data:       result,
	}

	c.JSON(http.StatusOK, response)
}

// metod login
func (a *AccountHandler) Login(c *gin.Context) {
	var request dto.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		msg, ok := err.(validator.ValidationErrors)
		errorObject := []helper.ErrorValidation{}
		if ok {
			for _, value := range msg {
				errorObject = append(errorObject, helper.ErrorValidation{
					Field:        value.Field(),
					ErrorMessage: helper.ErrorMessageFromTag(value.Tag()),
				})
			}

			var errMessages []string
			for _, value := range errorObject {
				errMessages = append(errMessages, value.ErrorMessage)
			}

			response := dto.ApiMessage{
				StatusCode: http.StatusBadRequest,
				Status:     "bad reuqest",
				Message:    strings.Join(errMessages, ". "),
			}

			c.JSON(http.StatusBadRequest, &response)
			return
		}
	}

	// call procedure login in service
	result, err := a.AccountService.Login(c, &request)
	if err != nil {
		// not found
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			response := dto.ApiMessage{
				StatusCode: http.StatusNotFound,
				Status:     "not found",
				Message:    err.Error(),
			}

			c.JSON(http.StatusNotFound, &response)
			return
		}

		// bad request
		if strings.Contains(strings.ToLower(err.Error()), "not match") {
			response := dto.ApiMessage{
				StatusCode: http.StatusBadRequest,
				Status:     "bad request",
				Message:    err.Error(),
			}

			c.JSON(http.StatusBadRequest, &response)
			return
		}

		// internal server error
		response := dto.ApiMessage{
			StatusCode: http.StatusInternalServerError,
			Status:     "internal server error",
			Message:    err.Error(),
		}

		c.JSON(http.StatusInternalServerError, &response)
		return
	}

	// success login
	response := dto.ApiMessage{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Message:    "success login",
		Data:       result,
	}

	c.JSON(http.StatusOK, &response)
}
