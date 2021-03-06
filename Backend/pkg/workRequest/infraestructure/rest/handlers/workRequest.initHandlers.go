package workRequest_handlers

import (
	ports "zetting/pkg/workRequest/core/ports"
)

import "github.com/gofiber/fiber/v2"

type WorkRequestHandler interface {
	CreateWorkRequest(c *fiber.Ctx) error
	GetWorkRequestsByWorker(c *fiber.Ctx) error
	GetWorkRequestsByProject(c *fiber.Ctx) error
	GetWorkRequestsById(c *fiber.Ctx) error
	AnswerWorkRequest(c *fiber.Ctx) error
}
type workRequestHandler struct {
	workRequestService ports.WorkRequestService
}

func NewWorkRequestHandler(workRequestService ports.WorkRequestService) WorkRequestHandler {
	return &workRequestHandler{
		workRequestService,
	}
}
