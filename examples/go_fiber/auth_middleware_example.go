// examples/go_fiber/auth_middleware_example.go
package middleware

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Claims struct to hold JWT claims
type Claims struct {
	UserID uint `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// JWTAuthMiddleware authenticates requests using JWT.
func JWTAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header required"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // No "Bearer " prefix
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Malformed token"})
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Ensure token signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token signature"})
			}
			if time.Now().Unix() > claims.ExpiresAt.Unix() { // Check expiration explicitly
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token expired"})
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token: " + err.Error()})
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Store user info in Fiber locals for later use
		c.Locals("user_id", claims.UserID)
		c.Locals("user_role", claims.Role)

		return c.Next() // Continue to the next handler
	}
}

// RoleBasedAuthMiddleware checks if the user has the required role.
func RoleBasedAuthMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole, ok := c.Locals("user_role").(string)
		if !ok || userRole != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access forbidden: Insufficient role"})
		}
		return c.Next()
	}
}