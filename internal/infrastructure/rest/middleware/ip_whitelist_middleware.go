package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-template/internal/utils/ext"
	"go-fiber-template/pkg/logging"
)

var log = logging.New(logging.Config{
	FileName: "ip_wl_middleware",
	Name:     "ip_wl_middleware",
},
)

func WhileListMiddleware(allowedIPs, allowedHosts []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		host := c.Hostname()
		log.Info("Incoming source IP: %s; Host: %s", ip, host)
		if ext.Contains(allowedIPs, ip) || ext.Contains(allowedHosts, host) {
			log.Info("IP есть в Whetelist")
			return c.Next()
		}
		log.Info("IP отстусвует в Whetelist")
		return c.SendStatus(fiber.StatusForbidden)
	}
}
