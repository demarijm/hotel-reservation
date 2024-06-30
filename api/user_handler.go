package api

import (
	"context"

	"github.com/demarijm/hotel-reservation/db"
	"github.com/demarijm/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

// ListAccounts lists all existing accounts
//
//	@Summary      List accounts
//	@Description  get accounts
//	@Tags         accounts
//	@Accept       json
//	@Produce      json
//	@Param        id    query     string  false  "name search by id"  Format(id)
//	@Success      200  {array}   types.User
//	@Router       /user/{id} [get]
func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)

	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "Foot",
	}
	return c.JSON(u)
}
