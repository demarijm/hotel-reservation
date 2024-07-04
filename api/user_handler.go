package api

import (
	"errors"
	"fmt"

	"github.com/demarijm/hotel-reservation/db"
	"github.com/demarijm/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		id = c.Params("id")
	)

	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"error": "Not found"})
		}
		return err

	}
	return c.JSON(user)
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	var (
		values types.UpdateUserParams
		userID = c.Params("id")
	)

	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	if err := c.BodyParser(&values); err != nil {
		return err
	}

	filter := bson.M{"_id": oid}
	if err := h.userStore.UpdateUser(c.Context(), filter, values); err != nil {
		return err
	}
	return c.JSON(map[string]string{"updated": userID})
}
func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	userID := c.Params(("id"))
	if err := h.userStore.DeleteUser(c.Context(), userID); err != nil {
		return err
	}
	return c.JSON(map[string]string{"deleted": userID})
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	fmt.Println(users)
	return c.JSON(users)
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if err := params.Validate(); len(err) > 0 {
		return c.JSON(err)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}
