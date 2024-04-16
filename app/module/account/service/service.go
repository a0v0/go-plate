package service

import (
	"go_plate/app/module/account/biz"
	_ "go_plate/app/module/account/request"
	"go_plate/pkg/response"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Service struct {
	biz *biz.Biz
	log *zap.Logger
}

func NewService(b *biz.Biz, log *zap.Logger) *Service {
	return &Service{
		biz: b,
		log: log,
	}
}

// @Summary Notify server about new user
// @Description After registering a new user at supertokens, the client must inform the server about the new user, so that server can issue an account for the user
// @Tags Server
// @Produce json
// @Success 200 {array} request.GetSelfAccountResponse
// @Failure 400,401,403 {object} response.Error "Error"
// @Router /notify [post]
func (svc *Service) NotifyServer(c *fiber.Ctx) error {
	err := svc.biz.NotifyServer("")
	if err != nil {
		svc.log.Error("Error while notifying server about new user", zap.Error(err))
		return err
	}

	return nil
}

// GetSelfAccount returns account information of the currently logged in user
// @Description Returns account information of the currently logged in user
// @Summary Get Self Account
// @Router /user/account [get]
// @Tags Account
// @Success 200  {object}   response.Response
// @Failure 400,401,403 {object} response.Error "Error"
func (svc *Service) GetSelfAccount(c *fiber.Ctx) error {
	account, err := svc.biz.GetSelfAccount("")
	if err != nil {
		svc.log.Error("Error while getting self account", zap.Error(err))
		return err
	}

	return response.Resp(c, response.Response{
		Messages: nil,
		Data:     account,
	})
}

// @Summary      Update Self Profile
// @Description  Update the  profile of currently logged in user
// @Description  All the fields are required to be specified, even if you don't want to update them.
// @Description  If you want to unset a field, set it to the null value depending on the field type.
// @Router       /user/account [post]
// @Accept       json
// @Produce      json
// @Success      200  {object}   response.Response
// @Failure 400,401 {object} response.Error "Error"
// @Param        body  body     request.AccountUpdateRequest true  "body"
// @Tags         Account
func (svc *Service) UpdateSelfAccount(c *fiber.Ctx) error {
	err := svc.biz.UpdateSelfAccount("", c)
	if err != nil {
		svc.log.Error("Error while updating self account", zap.Error(err))
		return err
	}
	return nil
}
