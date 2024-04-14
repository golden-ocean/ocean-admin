package menu

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golden-ocean/ocean-admin/pkg/common/response"
	"github.com/golden-ocean/ocean-admin/pkg/utils"
)

type Handler struct {
	service *Service
}

func NewHandler() *Handler {
	return &Handler{
		service: NewService(),
	}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	r := &CreateInput{}
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := utils.ValidateStruct(r); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	if err := h.service.Create(r); err != nil {
		return err
	}
	return c.JSON(response.OK(CreatedSuccess))
}

func (h *Handler) Update(c *fiber.Ctx) error {
	r := &UpdateInput{}
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := utils.ValidateStruct(r); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	if err := h.service.Update(r); err != nil {
		return err
	}
	return c.JSON(response.OK(UpdatedSuccess))
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	r := &DeleteInput{}
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := h.service.Delete(r); err != nil {
		return err
	}
	return c.JSON(response.OK(DeletedSuccess))
}

func (h *Handler) QueryTree(c *fiber.Ctx) error {
	w := &WhereParams{}
	if err := c.QueryParser(w); err != nil {
		return err
	}
	if err := utils.ValidateStruct(w); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	es, err := h.service.QueryTree(w)
	if err != nil {
		return err
	}
	return c.JSON(response.OK(es))
}
