package server

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

type LogApiStruct struct {
	IP            string            `json:"ip"`
	URL           string            `json:"url"`
	StartTime     string            `json:"start_time"`
	EndTime       string            `json:"end_time"`
	Duration      int64             `json:"duration"`
	Status        int               `json:"status"`
	Method        string            `json:"method"`
	RequestId     string            `json:"request_id"`
	RequestBody   string            `json:"request_body"`
	RequestHeader map[string]string `json:"request_header"`
	RequestParam  string            `json:"request_param"`
	ResponseBody  string            `json:"response_body"`
}

func LogApi(ctx *fiber.Ctx) error {

	t := time.Now()

	logStruct := LogApiStruct{
		IP:        ctx.IP(),
		URL:       ctx.OriginalURL(),
		StartTime: t.Format(DATE_FORMAT),
		Method:    ctx.Method(),
		RequestId: ctx.Locals("requestid").(string),
	}
	result := ctx.Next()
	if GetConfigLogs() {
		logStruct.RequestHeader = ctx.GetReqHeaders()
		logStruct.RequestParam = ctx.Request().URI().QueryArgs().String()
		logStruct.ResponseBody = string(ctx.Response().Body()[:])
	}

	logStruct.Status = ctx.Response().StatusCode()
	logStruct.EndTime = time.Now().Format(DATE_FORMAT)
	logStruct.Duration = time.Since(t).Milliseconds()

	logStr, _ := json.Marshal(logStruct)
	log.Printf("%s", string(logStr))

	return result
}
