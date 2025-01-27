package userController

import (
	"net/http"

	"github.com/TIM-DEBUG-ProjectSprintBatch3/TutupLapak/user/src/exceptions"
	"github.com/TIM-DEBUG-ProjectSprintBatch3/TutupLapak/user/src/helper"
	functionCallerInfo "github.com/TIM-DEBUG-ProjectSprintBatch3/TutupLapak/user/src/logger/helper"
	loggerZap "github.com/TIM-DEBUG-ProjectSprintBatch3/TutupLapak/user/src/logger/zap"
	"github.com/TIM-DEBUG-ProjectSprintBatch3/TutupLapak/user/src/model/dtos/request"
	fileService "github.com/TIM-DEBUG-ProjectSprintBatch3/TutupLapak/user/src/services/external/file"
	userService "github.com/TIM-DEBUG-ProjectSprintBatch3/TutupLapak/user/src/services/user"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do/v2"
)

type UserControllerInterface interface {
	RegisterByEmail(C *fiber.Ctx) error
	RegisterByPhone(C *fiber.Ctx) error
	LoginByEmail(C *fiber.Ctx) error
	LoginByPhone(C *fiber.Ctx) error

	LinkEmail(C *fiber.Ctx) error
	LinkPhone(C *fiber.Ctx) error
	GetUserProfile(C *fiber.Ctx) error
	UpdateUserProfile(C *fiber.Ctx) error
}

type UserController struct {
	userService userService.UserServiceInterface
	fileService fileService.FileServiceInterface
	logger      loggerZap.LoggerInterface
}

func NewUserController(userService userService.UserServiceInterface, fileService fileService.FileServiceInterface, logger loggerZap.LoggerInterface) UserControllerInterface {
	return &UserController{userService: userService, logger: logger}
}

func NewUserControllerInject(i do.Injector) (UserControllerInterface, error) {
	_userService := do.MustInvoke[userService.UserServiceInterface](i)
	_fileService := do.MustInvoke[fileService.FileServiceInterface](i)
	_logger := do.MustInvoke[loggerZap.LoggerInterface](i)
	return NewUserController(_userService, _fileService, _logger), nil
}

// Authentications

// Auth godoc
// @Summary
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.AuthByEmailRequest true "Payload"
// @Success 201 {object} response.AuthResponse "success response"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal server error"
// @Router /v1/register/email [post]
func (uc *UserController) RegisterByEmail(c *fiber.Ctx) error {
	userRequestParse := request.AuthByEmailRequest{}

	if err := c.BodyParser(&userRequestParse); err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerRegisterByEmail)
		return c.Status(http.StatusBadRequest).JSON(exceptions.ErrBadRequest(err.Error()))
	}

	response, err := uc.userService.RegisterByEmail(c, userRequestParse)
	if err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerRegisterByEmail, userRequestParse)
		return c.Status(int(err.(exceptions.ErrorResponse).StatusCode)).
			JSON(err)
	}

	c.Set(helper.X_AUTHOR_HEADER_KEY, helper.X_AUTHOR_HEADER_VALUE)
	return c.Status(fiber.StatusCreated).JSON(response)
}

// Auth godoc
// @Summary
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.AuthByPhoneRequest true "Payload"
// @Success 201 {object} response.AuthResponse "success response"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal server error"
// @Router /v1/register/phone [post]
func (uc *UserController) RegisterByPhone(c *fiber.Ctx) error {
	userRequestParse := request.AuthByPhoneRequest{}

	if err := c.BodyParser(&userRequestParse); err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerRegisterByPhone)
		return c.Status(http.StatusBadRequest).JSON(exceptions.ErrBadRequest(err.Error()))
	}

	response, err := uc.userService.RegisterByPhone(c, userRequestParse)
	if err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerRegisterByPhone, userRequestParse)
		return c.Status(int(err.(exceptions.ErrorResponse).StatusCode)).
			JSON(err)
	}

	c.Set(helper.X_AUTHOR_HEADER_KEY, helper.X_AUTHOR_HEADER_VALUE)
	return c.Status(fiber.StatusCreated).JSON(response)
}

// Auth godoc
// @Summary
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.AuthByEmailRequest true "Payload"
// @Success 200 {object} response.AuthResponse "success response"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal server error"
// @Router /v1/login/email [post]
func (uc *UserController) LoginByEmail(c *fiber.Ctx) error {
	userRequestParse := request.AuthByEmailRequest{}

	if err := c.BodyParser(&userRequestParse); err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerLoginByEmail)
		return c.Status(http.StatusBadRequest).JSON(exceptions.ErrBadRequest(err.Error()))
	}

	response, err := uc.userService.LoginByEmail(c, userRequestParse)
	if err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerLoginByEmail, userRequestParse)
		return c.Status(int(err.(exceptions.ErrorResponse).StatusCode)).JSON(err)
	}

	c.Set(helper.X_AUTHOR_HEADER_KEY, helper.X_AUTHOR_HEADER_VALUE)
	return c.Status(fiber.StatusOK).JSON(response)
}

// Auth godoc
// @Summary
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.AuthByPhoneRequest true "Payload"
// @Success 200 {object} response.AuthResponse "success response"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal server error"
// @Router /v1/login/phone [post]
func (uc *UserController) LoginByPhone(c *fiber.Ctx) error {
	userRequestParse := request.AuthByPhoneRequest{}

	if err := c.BodyParser(&userRequestParse); err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerRegisterByPhone)
		return c.Status(http.StatusBadRequest).JSON(exceptions.ErrBadRequest(err.Error()))
	}

	response, err := uc.userService.LoginByPhone(c, userRequestParse)
	if err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerLoginByPhone, userRequestParse)
		return c.Status(int(err.(exceptions.ErrorResponse).StatusCode)).JSON(err)
	}

	c.Set(helper.X_AUTHOR_HEADER_KEY, helper.X_AUTHOR_HEADER_VALUE)
	return c.Status(fiber.StatusOK).JSON(response)
}

// end Authentications

// User Profile

// UserProfile godoc
// @Summary
// @Description
// @Tags UserProfile
// @Accept json
// @Produce json
// @Param request body request.LinkEmailRequest true "Payload"
// @Success 200 {object} response.UserResponse "success response"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal server error"
// @Router /v1/user/link/email [post]
func (uc *UserController) LinkEmail(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(exceptions.ErrUnauthorized)
	}

	userRequestParse := request.LinkEmailRequest{}

	if err := c.BodyParser(&userRequestParse); err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerLinkEmail)
		return c.Status(http.StatusBadRequest).JSON(exceptions.ErrBadRequest(err.Error()))
	}

	response, err := uc.userService.LinkEmail(c, userRequestParse, userId)
	if err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerLinkEmail, userRequestParse)
		return c.Status(int(err.(exceptions.ErrorResponse).StatusCode)).JSON(err)
	}

	c.Set(helper.X_AUTHOR_HEADER_KEY, helper.X_AUTHOR_HEADER_VALUE)
	return c.Status(fiber.StatusOK).JSON(response)
}

// UserProfile godoc
// @Summary
// @Description
// @Tags UserProfile
// @Accept json
// @Produce json
// @Param request body request.LinkPhoneRequest true "Payload"
// @Success 200 {object} response.UserResponse "success response"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal server error"
// @Router /v1/user/link/phone [post]
func (uc *UserController) LinkPhone(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(exceptions.ErrUnauthorized)
	}

	userRequestParse := request.LinkPhoneRequest{}

	if err := c.BodyParser(&userRequestParse); err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerLinkPhone)
		return c.Status(http.StatusBadRequest).JSON(exceptions.ErrBadRequest(err.Error()))
	}

	response, err := uc.userService.LinkPhone(c, userRequestParse, userId)
	if err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerLinkPhone, userRequestParse)
		return c.Status(int(err.(exceptions.ErrorResponse).StatusCode)).JSON(err)
	}

	c.Set(helper.X_AUTHOR_HEADER_KEY, helper.X_AUTHOR_HEADER_VALUE)
	return c.Status(fiber.StatusOK).JSON(response)
}

// UserProfile godoc
// @Summary
// @Description
// @Tags UserProfile
// @Accept json
// @Produce json
// @Success 200 {object} response.UserResponse "success response"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal server error"
// @Router /v1/user [get]
func (uc *UserController) GetUserProfile(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(exceptions.ErrUnauthorized)
	}

	response, err := uc.userService.GetUserProfile(c, userId)
	if err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerGetUserProfile)
		return c.Status(int(err.(exceptions.ErrorResponse).StatusCode)).JSON(err)
	}

	c.Set(helper.X_AUTHOR_HEADER_KEY, helper.X_AUTHOR_HEADER_VALUE)
	return c.Status(fiber.StatusOK).JSON(response)
}

// UserProfile godoc
// @Summary
// @Description
// @Tags UserProfile
// @Accept json
// @Produce json
// @Param request body request.UpdateUserProfileRequest true "Update Profile Payload"
// @Success 200 {object} response.UserResponse "success response"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 500 {object} map[string]interface{} "internal server error"
// @Router /v1/user [post]
func (uc *UserController) UpdateUserProfile(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(exceptions.ErrUnauthorized)
	}

	userRequestParse := request.UpdateUserProfileRequest{}

	if err := c.BodyParser(&userRequestParse); err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerUpdateUserProfile)
		return c.Status(http.StatusBadRequest).JSON(exceptions.ErrBadRequest(err.Error()))
	}

	response, err := uc.userService.UpdateUserProfile(c, userRequestParse, userId)
	if err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerUpdateUserProfile)
		return c.Status(int(err.(exceptions.ErrorResponse).StatusCode)).JSON(err)
	}

	c.Set(helper.X_AUTHOR_HEADER_KEY, helper.X_AUTHOR_HEADER_VALUE)
	return c.Status(fiber.StatusOK).JSON(response)
}

// end User Profile
