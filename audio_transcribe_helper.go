package audiotranscribe

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
)

func transcribes(ctx context.Context, client *speech.Client, bqReq *BigQueryRequest) *BigQueryResponse {
	transcripts := make([]string, len(bqReq.Calls))
	wait := new(sync.WaitGroup)

	for i, call := range bqReq.Calls {
		wait.Add(1)

		go func(j int) {
			defer wait.Done()

			for {
				select {
				case <-ctx.Done():
					return
				default:
					log.Printf("Goroutines #%v is started", j)

					transcript := transcribe(ctx, client, fmt.Sprint(call[0]))
					transcripts[j] = transcript

					return
				}
			}
		}(i)
	}
	wait.Wait()

	bqResp := new(BigQueryResponse)
	bqResp.Replies = transcripts

	return bqResp
}

func transcribe(ctx context.Context, client *speech.Client, uri string) string {
	audio := &speechpb.RecognitionAudio{
		AudioSource: &speechpb.RecognitionAudio_Uri{
			Uri: uri,
		},
	}

	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: config,
		Audio:  audio,
	})
	if err != nil {
		transcript := Transcript{
			LogMessage: err.Error(),
		}

		return transcript.toJSONString()
	}

	var temp tempTranscript

	for _, result := range resp.Results {
		temp.result = append(temp.result, result.Alternatives[0].Transcript)
		temp.confidence = append(temp.confidence, result.Alternatives[0].Confidence)
	}

	transcript := Transcript{
		Result:     strings.Join(temp.result, ""),
		Confidence: temp.avgConfidence(),
	}

	return transcript.toJSONString()
}

func sendError(w http.ResponseWriter, err error, code int) {
	bqResp := new(BigQueryResponse)
	bqResp.ErrorMessage = fmt.Sprintf("Got error with details: %v", err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(bqResp)
}

func sendSuccess(w http.ResponseWriter, bqResp *BigQueryResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bqResp)
}

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
