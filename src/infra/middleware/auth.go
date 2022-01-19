package middleware

import (
	"helley/src/app/repository"
	"helley/src/app/security"
	"helley/src/domain"
	"strings"

	"github.com/labstack/echo"
)

const authorizationHeaderKey = "authorization"
const authorizationTypeBearer = "bearer"

type AuthMiddleware interface {
	Middleware(next echo.HandlerFunc) echo.HandlerFunc
}

type authMiddleware struct {
	tokenManager   security.TokenManager
	userRepository repository.UserRepository
}

func NewAuthMiddleware(tokenManager security.TokenManager, userRepository repository.UserRepository) AuthMiddleware {
	return &authMiddleware{
		tokenManager:   tokenManager,
		userRepository: userRepository,
	}
}

func (m *authMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			return security.ErrUnauthorized
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			return security.ErrUnauthorized
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			return security.ErrUnauthorized
		}

		accessToken := fields[1]
		claims, err := m.tokenManager.ValidateAccessToken(accessToken)
		if err != nil {
			return security.ErrUnauthorized
		}

		userId := claims["sub"].(string)

		user, err := m.userRepository.FindById(userId)
		if err != nil {
			return security.ErrUnauthorized
		}

		c.Set("User", domain.UserInfo{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})

		return next(c)
	}
}
