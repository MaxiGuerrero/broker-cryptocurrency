package healthcheck

import responses "broker-cryptocurrency/backend/src/system/responses"

type HealthCheckController struct{}

func (h HealthCheckController) controller() *responses.Response {
	return responses.OK()
}
