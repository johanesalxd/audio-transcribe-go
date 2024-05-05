package audiotranscribe

import "encoding/json"

type Transcript struct {
	Result     string  `json:"result"`
	Confidence float32 `json:"confidence"`
	LogMessage string  `json:"log_message"`
}

type tempTranscript struct {
	result     []string
	confidence []float32
}

func (t Transcript) toJSONString() string {
	jsonTranscript, _ := json.Marshal(t)

	return string(jsonTranscript)
}

func (t tempTranscript) avgConfidence() float32 {
	var sum float32

	if len(t.confidence) == 0 {
		return float32(0)
	}

	for i := range t.confidence {
		sum += t.confidence[i]
	}

	return sum / float32(len(t.confidence))
}
