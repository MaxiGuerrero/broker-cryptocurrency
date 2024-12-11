package authentication_test

import (
	"backend/src/authentication"
	"backend/src/authentication/models"
	"backend/tests/authentication/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var fakeEncrypter = &mocks.FakeEncrypter{}
var fakeAuthRepository = &mocks.FakeAuthRepository{}
var service = authentication.NewAuthService(fakeEncrypter, fakeAuthRepository)

const username = "maxi"
const password = "123"
const email = "maxi@123"
const passwordHashed = "hashed"

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

func Test_Service_Register_Success(t *testing.T) {
	// Arrenge
	fakeEncrypter.On("GenerateHash", password).Return(passwordHashed)
	fakeAuthRepository.On("FindUserByUsernameAndEmail", username, email).Return(nil).Once()
	fakeAuthRepository.On("CreateUser", username, passwordHashed, email).Return(userInfoMock).Once()
	// Act
	userinfo, err := service.Register(username, password, email)
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, username, userinfo.Username)
	assert.Equal(t, email, userinfo.Email)
	assert.Equal(t, models.USER, userinfo.Role)
	assert.Equal(t, models.ACTIVE, userinfo.Status)
}

func Test_Service_Register_Error_User_Not_Found(t *testing.T) {
	// Arrenge
	fakeAuthRepository.On("FindUserByUsernameAndEmail", username, email).Return(userMock).Once()
	// Act
	userinfo, err := service.Register(username, password, email)
	// Assert
	assert.Nil(t, userinfo)
	assert.Error(t, err)
	fakeAuthRepository.AssertCalled(t, "FindUserByUsernameAndEmail", username, email)
	fakeAuthRepository.AssertNotCalled(t, "CreateUser")
	fakeEncrypter.AssertNotCalled(t, "GenerateHash")
}
