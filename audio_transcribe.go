package audiotranscribe

import (
	"context"
	"encoding/json"
	"net/http"

	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
)

var config = &speechpb.RecognitionConfig{
	Encoding:        speechpb.RecognitionConfig_LINEAR16,
	SampleRateHertz: 8000,
	LanguageCode:    "en-US",
}

func AudioTranscribe(w http.ResponseWriter, r *http.Request) {
	bqReq := new(BigQueryRequest)
	if err := json.NewDecoder(r.Body).Decode(bqReq); err != nil {
		sendError(w, err, http.StatusBadRequest)

		return
	}

	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)

		return
	}

	bqResp := transcribes(ctx, client, bqReq)
	sendSuccess(w, bqResp)
}
