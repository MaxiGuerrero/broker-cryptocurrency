package register

import (
	response "broker-cryptocurrency/backend/src/system/server/responses"
	"broker-cryptocurrency/backend/src/system/utils"
	"broker-cryptocurrency/backend/src/users/models"
)

func CreateUserController(req models.CreateUserRequest, service ICreateUserService) *response.Response {
	if badSchema := utils.ValidateSchema(&req); badSchema != nil {
		return response.BadRequest(badSchema.Error())
	}
	businessErr := service.CreateUser(req.Username, req.Password, req.Email, req.Role)
	if businessErr != nil {
		return response.BadRequest(businessErr.Error())
	}
	return response.OK()
}
