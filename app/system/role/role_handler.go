package role

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golden-ocean/ocean-admin/pkg/common/constants"
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
	if err := utils.ValidateStruct(r); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	if err := h.service.Delete(r); err != nil {
		return err
	}
	return c.JSON(response.OK(DeletedSuccess))
}

func (h *Handler) QueryPage(c *fiber.Ctx) error {
	w := &WhereParams{
		Current:  constants.Current,
		PageSize: constants.PageSize,
	}
	if err := c.QueryParser(w); err != nil {
		return err
	}
	if err := utils.ValidateStruct(w); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	es, total, err := h.service.QueryPage(w)
	if err != nil {
		return err
	}
	return c.JSON(response.Page(es, w.Current, w.PageSize, total))
}

func (h *Handler) QueryAll(c *fiber.Ctx) error {
	w := &WhereParams{}
	if err := c.QueryParser(w); err != nil {
		return err
	}
	if err := utils.ValidateStruct(w); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	es, err := h.service.QueryAll(w)
	if err != nil {
		return err
	}
	return c.JSON(response.OK(es))
}

func (h *Handler) QueryMenus(c *fiber.Ctx) error {
	w := &RoleMenuInput{}
	if err := c.ParamsParser(w); err != nil {
		return err
	}
	es, err := h.service.QueryMenus(w)
	if err != nil {
		return err
	}
	return c.JSON(response.OK(es))
}

func (h *Handler) GrantMenus(c *fiber.Ctx) error {
	r := &RoleMenuInput{}
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := utils.ValidateStruct(r); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	if err := h.service.GrantMenus(r); err != nil {
		return err
	}
	return c.JSON(response.OK(CreatedSuccess))
}
