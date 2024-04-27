Audio Transcribe Go
-----------------------------
Speech-to-text application written in Go (based on Function Framework) to be used primarily by BigQuery Remote Function.

# How to run
## Run locally:
```
go run cmd/main.go
```

## Run locally with Pack and Docker:
```
pack build --builder=gcr.io/buildpacks/builder audio-transcribe-go

gcloud auth application-default login

ADC=~/.config/gcloud/application_default_credentials.json \
docker run -p8080:8080 \
-e GOOGLE_APPLICATION_CREDENTIALS=/tmp/keys/secret.json \
-v ${ADC}:/tmp/keys/secret.json \
audio-transcribe-go
```

## Test locally (accept BQ RF [request contract](https://cloud.google.com/bigquery/docs/remote-functions#input_format)):
```
curl -m 60 -X POST localhost:8080 \
-H "Content-Type: application/json" \
-d '{
  "requestId": "",
  "caller": "",
  "sessionUser": "",
  "userDefinedContext": {},
  "calls": [
    ["gcs_audio_wav_8khz_uri_1"],
    ["gcs_audio_wav_8khz_uri_2"],
    ["gcs_audio_wav_8khz_uri_n"]
  ]
  }'
```

## Run on Cloud Run:
[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)

## Run on Cloud Function:
```
gcloud functions deploy audio-transcribe-go \
    --gen2 \
    --runtime=go122 \
    --region=us-central1 \
    --source=. \
    --entry-point=AudioTranscribe \
    --trigger-http \
    --allow-unauthenticated
```

# Additional notes
* https://cloud.google.com/bigquery/docs/remote-functions
* https://cloud.google.com/functions/docs/concepts/go-runtime
* https://cloud.google.com/docs/buildpacks/build-function
* https://www.voiptroubleshooter.com/open_speech/american.html
* https://www.cs.columbia.edu/~hgs/audio/harvard.html
