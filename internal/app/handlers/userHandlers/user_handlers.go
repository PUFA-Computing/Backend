package userHandlers

import (
	"Backend/internal/app/domain/roles"
	"Backend/internal/app/domain/user"
	user2 "Backend/internal/app/interfaces/service/userService"
	"Backend/internal/utils/token"
	"github.com/gofiber/fiber/v2"
	"time"
)

type UserHandlers struct {
	userService user2.UserServices
}

func NewUserHandlers(userService user2.UserServices) *UserHandlers {
	return &UserHandlers{userService: userService}
}

func (h *UserHandlers) RegisterUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user user.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error parsing userService"})
		}

		user.Role = roles.RoleUser

		if err := h.userService.RegisterUser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error creating userService", "error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
	}
}

func (h *UserHandlers) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var loginData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&loginData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error parsing login data"})
		}

		user, err := h.userService.AuthenticateUser(loginData.Email, loginData.Password)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid email or password", "error": err.Error()})
		}

		sessionToken, err := token.GenerateJWTToken(user.User.ID, user.User.Role)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error generating session token"})
		}

		expirationTime := time.Now().Add(token.SessionDuration)
		if err := token.StoreSessionData(user.User.ID, sessionToken, expirationTime); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error storing session data", "error": err.Error()})
		}

		c.Cookie(&fiber.Cookie{
			Name:     "session_token",
			Value:    sessionToken,
			Expires:  expirationTime,
			Path:     "/",
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login successful", "session_token": sessionToken})
	}
}

func (h *UserHandlers) Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "session_token",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
			Secure:   true,
		})

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logout successful"})
	}
}