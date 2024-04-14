package auth

import (
	"time"

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

func (h *Handler) Login(c *fiber.Ctx) error {
	req := &LoginInput{}
	if err := c.BodyParser(req); err != nil {
		return err
	}
	if err := utils.ValidateStruct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	e, err := h.service.Login(req)
	if err != nil {
		return err
	}
	tokens, err := utils.GenerateNewTokens(e.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(response.OK(tokens))
}

// func (h *Handler) Logout(c *fiber.Ctx) error {
// 	req := &UpdateInput{}
// 	if err := c.BodyParser(req); err != nil {
// 		return err
// 	}
// 	if err := utils.ValidateStruct(req); err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err...)
// 	}
// 	if err := h.service.Update(req); err != nil {
// 		return err
// 	}
// 	return c.JSON(response.OK(UpdatedSuccess))
// }

func (h *Handler) QueryInfo(c *fiber.Ctx) error {
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	e, err := h.service.QueryInfo(claims.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(response.OK(e))
}

func (h *Handler) Refresh(c *fiber.Ctx) error {
	now := time.Now().Unix()
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return err
	}
	if now < claims.Expire {
		return fiber.NewError(fiber.StatusUnauthorized, ErrorCheckExpiredTime)
	}
	req := &RefreshInput{}
	if err := c.BodyParser(req); err != nil {
		return err
	}
	refresh_time, err := utils.ParseRefreshToken(req.Refresh)
	if err != nil {
		return err
	}
	if now < refresh_time {
		return fiber.NewError(fiber.StatusUnauthorized, ErrorRefreshExpiredTime)
	}
	// 转到service 中，在redis中查找相应的过期时间，合适则返回并redis更新时间
	// e, err := h.service.Refresh(id)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }
	tokens, err := utils.GenerateNewTokens(claims.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(response.OK(tokens))
}

func (h *Handler) Test(c *fiber.Ctx) error {
	// []int{1, 3, 4, 5}, []int{6}
	return c.JSON(response.OK("Xxx"))
}
