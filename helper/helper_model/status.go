package helper_model

import (
	"github.com/faridlan/web-basketball-extra/model/domain"
	"github.com/faridlan/web-basketball-extra/model/web/web_status"
)

func StatusResponse(role domain.Status) web_status.StatusResponse {
	return web_status.StatusResponse{
		Id:   role.Id,
		Name: role.Name,
	}
}

func StatusResponses(stats []domain.Status) []web_status.StatusResponse {
	statusResponse := []web_status.StatusResponse{}
	for _, status := range stats {
		statusResponse = append(statusResponse, StatusResponse(status))
	}

	return statusResponse
}
