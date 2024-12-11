package authentication_test

import (
	"backend/src/authentication"
	"backend/src/authentication/models"
	responses "backend/src/system/responses"
	"backend/tests/authentication/fakes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fakeService = &fakes.FakeService{}
var controller = authentication.NewAuthController(fakeService)

var userInfo = &models.UserInfo{
	UserId:   "1",
	Username: "maxi",
	Email:    "maxi@email.com",
	Role:     models.USER.String(),
	Status:   models.ACTIVE.String(),
}

var tokenService = "token"
var responseToken = &models.LoginResponse{
	Token: tokenService,
}

func Test_Controller_Register_Success(t *testing.T) {
	// Arrenge
	var responseSuccessExpected = responses.OK_WITH_DATA(userInfo)
	var req = &models.RegisterRequest{Username: "maxi", Password: "123456", Email: "maxi@email.com"}
	fakeService.On("Register", req.Username, req.Password, req.Email).Return(userInfo, nil).Once()
	// Act
	response := controller.Register(req)
	// Assert
	assert.Equal(t, responseSuccessExpected, response)
}

func Test_Controller_Register_Error_BadRequest_ErrorService(t *testing.T) {
	// Arrenge
	var req = &models.RegisterRequest{Username: "maxi", Password: "123456", Email: "maxi@email.com"}
	fakeService.On("Register", req.Username, req.Password, req.Email).Return(nil, errors.New("A service error")).Once()
	// Act
	response := controller.Register(req)
	// Assert
	assert.Equal(t, 400, response.Code)
	assert.NotNil(t, response.Message)
}

func Test_Controller_Register_Error_BadRequest_BadSchema(t *testing.T) {
	// Arrenge
	var req = &models.RegisterRequest{Username: "m", Password: "123", Email: "maxiemail.com"}
	// Act
	response := controller.Register(req)
	// Assert
	assert.Equal(t, 400, response.Code)
	assert.NotNil(t, response.Message)
	fakeService.AssertNotCalled(t, "Register")
}

func Test_Controller_Login_Success(t *testing.T) {
	// Arrenge
	var responseSuccessExpected = responses.OK_WITH_DATA(responseToken)
	var req = &models.LoginRequest{Username: "maxi", Password: "123456"}
	fakeService.On("Login", req.Username, req.Password).Return(&tokenService, nil).Once()
	// Act
	response := controller.Login(req)
	// Assert
	assert.Equal(t, responseSuccessExpected, response)
}

func Test_Controller_Login_Error_BadRequest_ErrorService(t *testing.T) {
	// Arrenge
	var req = &models.LoginRequest{Username: "maxi", Password: "123456"}
	fakeService.On("Login", req.Username, req.Password).Return(nil, errors.New("A service error")).Once()
	// Act
	response := controller.Login(req)
	// Assert
	assert.Equal(t, 400, response.Code)
	assert.NotNil(t, response.Message)
}

func Test_Controller_Login_Error_BadRequest_BadSchema(t *testing.T) {
	// Arrenge
	var req = &models.LoginRequest{Username: "ma", Password: "1"}
	// Act
	response := controller.Login(req)
	// Assert
	assert.Equal(t, 400, response.Code)
	assert.NotNil(t, response.Message)
	fakeService.AssertNotCalled(t, "Login")
}
