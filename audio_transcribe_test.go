package audiotranscribe

import (
	"testing"
)

func TestTranscriptToJSONString(t *testing.T) {
	tests := []struct {
		name string
		t    *Transcript
		want string
	}{
		{
			name: "empty transcript",
			t:    &Transcript{},
			want: `{"result":"","confidence":0,"log_message":""}`,
		},
		{
			name: "transcript with text",
			t: &Transcript{
				Result: "hello world",
			},
			want: `{"result":"hello world","confidence":0,"log_message":""}`,
		},
		{
			name: "transcript with confidence",
			t: &Transcript{
				Confidence: 0.9,
			},
			want: `{"result":"","confidence":0.9,"log_message":""}`,
		},
		{
			name: "transcript with text and confidence",
			t: &Transcript{
				Result:     "hello world",
				Confidence: 0.9,
			},
			want: `{"result":"hello world","confidence":0.9,"log_message":""}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.t.toJSONString()
			if got != test.want {
				t.Errorf("Transcript.toJSONString() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestTempTranscriptAvgConfidence(t *testing.T) {
	tests := []struct {
		name string
		t    *tempTranscript
		want float32
	}{
		{
			name: "empty transcript",
			t:    &tempTranscript{},
			want: 0,
		},
		{
			name: "transcript with one confidence",
			t: &tempTranscript{
				confidence: []float32{0.9},
			},
			want: 0.9,
		},
		{
			name: "transcript with multiple confidences",
			t: &tempTranscript{
				confidence: []float32{0.9, 0.8, 0.7},
			},
			want: 0.8,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.t.avgConfidence()
			if got != test.want {
				t.Errorf("TempTranscript.avgConfidence() = %v, want %v", got, test.want)
			}
		})
	}
}
