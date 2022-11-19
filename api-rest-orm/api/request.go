package api

type RequestApiStruct struct {
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
