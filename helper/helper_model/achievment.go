package helper_model

import (
	"github.com/faridlan/web-basketball-extra/model/domain"
	"github.com/faridlan/web-basketball-extra/model/web/web_achievment"
)

func AchievmentResponse(role domain.Achievment) web_achievment.AchievmentResponse {
	return web_achievment.AchievmentResponse{
		Id:   role.Id,
		Desc: role.Desc,
	}
}

func AchievmentResponses(achievs []domain.Achievment) []web_achievment.AchievmentResponse {
	achievmentResponse := []web_achievment.AchievmentResponse{}
	for _, achiev := range achievs {
		achievmentResponse = append(achievmentResponse, AchievmentResponse(achiev))
	}

	return achievmentResponse
}
