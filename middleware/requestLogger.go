package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/cryptoAPI/database"
)

/*RequestLogger : Save the request in the database*/
func RequestLogger(c *fiber.Ctx) error {
	method := string(c.Request().Header.Method())
	url := string(c.Request().Header.RequestURI())
	database.SaveLog(method, url)
	return c.Next()
}
