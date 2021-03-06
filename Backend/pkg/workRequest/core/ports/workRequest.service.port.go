package workRequest_ports

import (
	models "zetting/pkg/workRequest/core/models"
)

type WorkRequestService interface {
	CreateWorkRequest(workRequest models.WorkRequest) (*models.WorkRequest, error)
	GetWorkRequests(referenceId interface{}, status string, page int, number int, document string) (models.WorkRequests, error)
	AnswerWorkRequest(workRequest models.WorkRequest) error
	GetWorkRequest(workRequestId interface{}) (*models.WorkRequest, error)
}
