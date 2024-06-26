package audiotranscribe

import (
	"context"
	"fmt"
	"log"
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

		go func(j int, url string) {
			defer wait.Done()

			for {
				select {
				case <-ctx.Done():
					log.Printf("Got cancellation signal in Goroutine #%d", j)

					return
				default:
					log.Printf("Running in Goroutine #%d for URL: %v", j, url)

					transcript := transcribe(ctx, client, url)
					transcripts[j] = transcript

					return
				}
			}
		}(i, fmt.Sprint(call[0]))
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
