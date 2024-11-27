package healthcheck

import responses "backend/src/system/responses"

type HealthCheckController struct{}

func (h HealthCheckController) controller() *responses.Response {
	return responses.OK()
}
