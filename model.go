package audiotranscribe

type BigQueryRequest struct {
	RequestID          string            `json:"requestId"`
	Caller             string            `json:"caller"`
	SessionUser        string            `json:"sessionUser"`
	UserDefinedContext map[string]string `json:"userDefinedContext"`
	Calls              [][]interface{}   `json:"calls"`
}

type BigQueryResponse struct {
	Replies      []string `json:"replies"`
	ErrorMessage string   `json:"errorMessage"`
}

type Transcript struct {
	Result     string  `json:"result"`
	Confidence float32 `json:"confidence"`
	LogMessage string  `json:"log_message"`
}

type TempTranscript struct {
	Result     []string
	Confidence []float32
}
