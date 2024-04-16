package biz

import (
	"go_plate/app/module/account/data"
	"go_plate/app/module/account/request"
	"go_plate/internal/ent"
	"go_plate/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type Biz struct {
	Data *data.Data
}

func NewBiz(data *data.Data) *Biz {
	return &Biz{
		Data: data,
	}
}

func (b *Biz) NotifyServer(userID string) error {
	// check if this user is registered at supertokens

	return nil
}

func (b *Biz) GetSelfAccount(userID string) (resp *request.GetSelfAccountResponse, err error) {
	a, err := b.Data.GetAccount(userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, response.NotFoundError
		}
		return nil, response.InternalError
	}

	p, err := b.Data.GetProfile(a)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, response.NotFoundError
		}
		return nil, response.InternalError
	}

	// Copy the account data.
	resp = &request.GetSelfAccountResponse{}
	err = copier.Copy(&resp, &a)
	if err != nil {
		return nil, err
	}

	// Copy the profile data.
	err = copier.Copy(&resp.Profile, &p)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Biz) UpdateSelfAccount(userID string, c *fiber.Ctx) error {
	// validate the request
	validate := validator.New()

	r := new(request.AccountUpdateRequest)
	if err := c.BodyParser(&r); err != nil {
		return response.InvalidRequestError
	}

	if r.Username != "" {

		if !s.Data.IsUsernameAvailable(r.Username) {
			return response.UsernameUnavailableError
		}
	}

	err := validate.Var(r.DisplayName, "max=50")
	if err != nil {
		return response.DisplayNameInvalidError
	}

	err = validate.Var(r.Bio, "max=160")
	if err != nil {
		return response.BioInvalidError
	}

	// prepare the data for update
	profile, err := s.Data.GetProfileByUserID(userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return response.NotFoundError
		}
		return response.InternalError
	}

	if err := c.BodyParser(&profile); err != nil {
		return err
	}

	newProfileData := &data.NewProfileData{}
	err = copier.Copy(&newProfileData, &profile)
	if err != nil {
		return response.InternalError
	}

	err = s.Data.UpdateProfile(userID, newProfileData)
	if err != nil {
		if ent.IsNotFound(err) {
			return response.NotFoundError
		}
		return response.InternalError
	}

	return nil
}
