package audiotranscribe

import "encoding/json"

func (t *Transcript) toJSONString() string {
	jsonTranscript, _ := json.Marshal(t)

	return string(jsonTranscript)
}

func (t *tempTranscript) avgConfidence() float32 {
	var sum float32

	for i := range t.confidence {
		sum += t.confidence[i]
	}

	return sum / float32(len(t.confidence))
}
