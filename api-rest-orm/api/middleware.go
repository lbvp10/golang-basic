package api

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"orm/api"
	"orm/logger"
	"os"
	"time"
)

func LogApi(ctx *fiber.Ctx) error {

	t := time.Now()

	logStruct := api.RequestApiStruct{
		IP:        ctx.IP(),
		URL:       ctx.OriginalURL(),
		StartTime: t.Format(api.DATE_FORMAT),
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
	logStruct.EndTime = time.Now().Format(api.DATE_FORMAT)
	logStruct.Duration = time.Since(t).Milliseconds()

	logStr, _ := json.Marshal(logStruct)
	logger.Info(fmt.Sprintf("%s", string(logStr)))

	return result
}
func GetConfigLogsRequest() bool {
	if os.Getenv("LOG_FULL") == "false" {
		return false
	}
	return true
}
