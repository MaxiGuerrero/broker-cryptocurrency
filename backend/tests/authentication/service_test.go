package authentication_test

import (
	"backend/src/authentication"
	"backend/src/authentication/models"
	"backend/src/system"
	"backend/tests/authentication/fakes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var fakeEncrypter = &fakes.FakeEncrypter{}
var fakeAuthRepository = &fakes.FakeAuthRepository{}
var fakeJWTBuilder = &fakes.FakeJWTBuilder{}
var service = authentication.NewAuthService(fakeEncrypter, fakeAuthRepository, fakeJWTBuilder)

const username = "maxi"
const password = "123"
const email = "maxi@123"
const passwordHashed = "hashed"
const token = "token"

var userMock = &models.User{
	ID:        "1",
	Username:  username,
	Password:  passwordHashed,
	Email:     email,
	Status:    models.ACTIVE,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	Role:      models.USER,
}
var userInfoMock = &models.UserInfo{
	UserId:   userMock.ID,
	Username: userMock.Username,
	Email:    userMock.Email,
	Status:   userMock.Status.String(),
	Role:     userMock.Role.String(),
}

var payload = &system.Payload{
	UserId:   userMock.ID,
	Username: userMock.Username,
	Role:     userMock.Role.String(),
}

func Test_Service_Register_Success(t *testing.T) {
	// Arrenge
	fakeEncrypter.On("GenerateHash", password).Return(passwordHashed)
	fakeAuthRepository.On("FindUserByUsername", username).Return(nil).Once()
	fakeAuthRepository.On("CreateUser", username, passwordHashed, email).Return(userInfoMock).Once()
	// Act
	userinfo, err := service.Register(username, password, email)
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, username, userinfo.Username)
	assert.Equal(t, email, userinfo.Email)
	assert.Equal(t, models.USER.String(), userinfo.Role)
	assert.Equal(t, models.ACTIVE.String(), userinfo.Status)
}

func Test_Service_Register_Error_User_Not_Found(t *testing.T) {
	// Arrenge
	fakeAuthRepository.On("FindUserByUsername", username).Return(userMock).Once()
	// Act
	userinfo, err := service.Register(username, password, email)
	// Assert
	assert.Nil(t, userinfo)
	assert.Error(t, err)
	fakeAuthRepository.AssertCalled(t, "FindUserByUsername", username)
	fakeAuthRepository.AssertNotCalled(t, "CreateUser")
	fakeEncrypter.AssertNotCalled(t, "GenerateHash")
}

func Test_Service_Login_Success(t *testing.T) {
	// Arrenge
	fakeAuthRepository.On("FindUserByUsername", username).Return(userMock).Once()
	fakeEncrypter.On("Compare", passwordHashed, password).Return(true).Once()
	fakeJWTBuilder.On("BuildToken", payload).Return(token).Once()
	// Act
	result, err := service.Login(username, password)
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, token, *result)
}

func Test_Service_Login_Error_User_Not_Found(t *testing.T) {
	// Arrenge
	fakeAuthRepository.On("FindUserByUsername", username).Return(nil).Once()
	// Act
	result, err := service.Login(username, password)
	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	fakeEncrypter.AssertNotCalled(t, "Compare")
	fakeJWTBuilder.AssertNotCalled(t, "BuildToken")
}

func Test_Service_Login_Error_User_Password_Wrong(t *testing.T) {
	// Arrenge
	fakeAuthRepository.On("FindUserByUsername", username).Return(userMock).Once()
	fakeEncrypter.On("Compare", passwordHashed, password).Return(false).Once()
	// Act
	result, err := service.Login(username, password)
	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	fakeEncrypter.AssertCalled(t, "Compare", passwordHashed, password)
	fakeJWTBuilder.AssertNotCalled(t, "BuildToken")
}

func Test_Service_Login_Error_User_Blocked(t *testing.T) {
	// Arrenge
	userBlocked := &models.User{
		Status: models.BLOCKED,
	}
	fakeAuthRepository.On("FindUserByUsername", username).Return(userBlocked).Once()
	// Act
	result, err := service.Login(username, password)
	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	fakeEncrypter.AssertNotCalled(t, "Compare")
	fakeJWTBuilder.AssertNotCalled(t, "BuildToken")
}

func Test_Service_Login_Error_User_Inactive(t *testing.T) {
	// Arrenge
	userInactive := &models.User{
		Status: models.INACTIVE,
	}
	fakeAuthRepository.On("FindUserByUsername", username).Return(userInactive).Once()
	// Act
	result, err := service.Login(username, password)
	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	fakeEncrypter.AssertNotCalled(t, "Compare")
	fakeJWTBuilder.AssertNotCalled(t, "BuildToken")
}
