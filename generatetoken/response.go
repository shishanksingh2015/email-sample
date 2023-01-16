package generatetoken

type JWT struct {
	Token   string `json:"token"`
	TraceId string `json:"traceId"`
}