package api

import (
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
	Type          string            `json:"type"`
}

func NewLogApi() LogApiStruct {

	t := time.Now()

	logStruct := LogApiStruct{
		IP:        "localhost",
		URL:       "http:/ssssssssss",
		StartTime: t.Format("2006-01-02T15:04:05-0700"),
		Method:    "ctx.Method()",
		RequestId: "ctx.Locals().(string)",
	}

	logStruct.RequestParam = " ctx.Request().URI().QueryArgs().String()"
	logStruct.ResponseBody = "string(ctx.Response().Body()[:])"

	logStruct.Status = 200
	logStruct.EndTime = time.Now().Format("2006-01-02T15:04:05-0700")
	logStruct.Duration = time.Since(t).Milliseconds()

	//logStr, _ := json.Marshal(logStruct)

	return logStruct
}
