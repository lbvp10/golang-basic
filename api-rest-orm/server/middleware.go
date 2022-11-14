package server

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"time"
)

type LogStruct struct {
	IP            string            `json:"ip"`
	URL           string            `json:"url"`
	StartTime     string            `json:"start_time"`
	EndTime       string            `json:"end_time"`
	Duration      int64             `json:"duration"`
	Status        int               `json:"status"`
	Method        string            `json:"method"`
	ResquestId    string            `json:"resquest_id"`
	RequestBody   string            `json:"requet_body"`
	RequestHeader map[string]string `json:"request_header"`
	RequestParam  string            `json:"request_param"`
	ResponseBody  string            `json:"response_body"`
}

func Log(ctx *fiber.Ctx) error {

	t := time.Now()

	logStruct := LogStruct{
		IP:         ctx.IP(),
		URL:        ctx.OriginalURL(),
		StartTime:  t.Format(DATE_FORMAT),
		Method:     ctx.Method(),
		ResquestId: ctx.Locals("requestid").(string),
	}
	result := ctx.Next()
	logStruct.Status = ctx.Response().StatusCode()
	logStruct.EndTime = time.Now().Format(DATE_FORMAT)
	logStruct.Duration = time.Since(t).Milliseconds()

	if getConfigLogs() {
		logStruct.RequestHeader = ctx.GetReqHeaders()
		logStruct.RequestParam = ctx.Request().URI().QueryArgs().String()
		logStruct.ResponseBody = string(ctx.Response().Body()[:])
	}

	logStr, _ := json.Marshal(logStruct)
	log.Printf("%s", string(logStr))

	return result
}

func getConfigLogs() bool {
	if os.Getenv("LOG_FULL") == "false" {
		return false
	}
	return true
}
