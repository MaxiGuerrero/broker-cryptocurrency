package authentication

import (
	interfaces "backend/src/authentication/interfaces"
	"backend/src/authentication/models"
	response "backend/src/system/responses"
	"backend/src/system/utils"
)

type AuthController struct {
	authService interfaces.IAuthService
}

func NewAuthController(authService interfaces.IAuthService) *AuthController {
	return &AuthController{
		authService,
	}
}

func (a *AuthController) Register(req *models.RegisterRequest) *response.Response {
	if badSchema := utils.ValidateSchema(req); badSchema != nil {
		return response.BadRequest(badSchema.Error())
	}
	userInfo, err := a.authService.Register(req.Username, req.Password, req.Email)
	if err != nil {
		return response.BadRequest(err.Error())
	}
	return response.OK_WITH_DATA(userInfo)
}
