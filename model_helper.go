package audiotranscribe

import "encoding/json"

func (t *Transcript) ToJSONString() string {
	jsonTranscript, _ := json.Marshal(t)

	return string(jsonTranscript)
}

func (t *TempTranscript) AvgConfidence() float32 {
	var sum float32

	if len(t.Confidence) == 0 {
		return float32(0)
	}

	for i := range t.Confidence {
		sum += t.Confidence[i]
	}

	return sum / float32(len(t.Confidence))
}
