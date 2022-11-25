package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"orm/logger"
	"os"
	"time"
)

func LogApi(ctx *fiber.Ctx) error {

	t := time.Now()

	logStruct := RequestApiStruct{
		IP:        ctx.IP(),
		URL:       ctx.OriginalURL(),
		StartTime: t.Format(DateFormat),
		Method:    ctx.Method(),
		RequestId: ctx.Locals("requestid").(string),
	}
	result := ctx.Next()
	if GetConfigLogsRequest() {
		logStruct.RequestHeader = ctx.GetReqHeaders()
		logStruct.RequestParam = ctx.Request().URI().QueryArgs().String()
		logStruct.ResponseBody = string(ctx.Response().Body()[:])
	}

	logStruct.Status = ctx.Response().StatusCode()
	logStruct.EndTime = time.Now().Format(DateFormat)
	logStruct.Duration = time.Since(t).Milliseconds()

	logger.InfoF(fmt.Sprintf("%s %s", ctx.Method(), logStruct.URL), logStruct)

	return result
}
func GetConfigLogsRequest() bool {
	if os.Getenv("LOG_FULL") == "false" {
		return false
	}
	return true
}
