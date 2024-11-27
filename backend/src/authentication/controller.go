package authentication

import (
	interfaces "backend/src/authentication/interfaces"
	"backend/src/authentication/models"
	response "backend/src/system/responses"
	"backend/src/system/utils"
)

type AuthController struct {
	AuthService interfaces.IAuthService
}

func (a *AuthController) register(req models.RegisterRequest) *response.Response {
	if badSchema := utils.ValidateSchema(&req); badSchema != nil {
		return response.BadRequest(badSchema.Error())
	}
	return response.OK()
}
