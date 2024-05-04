Audio Transcribe Go
-----------------------------
Speech-to-text application written in Go (based on Function Framework) to be used primarily by BigQuery Remote Function (BQ RF).

# Project structures
* **audio_transcribe\*.go** are the **main/logic functions**.
* **server\*.go** and others are the **generic functions** (boilerplate).

# How to run
## Run locally
```
FUNCTION_TARGET=AudioTranscribe go run cmd/main.go
```

## Run locally with Pack and Docker
```
pack build --builder=gcr.io/buildpacks/builder audio-transcribe-go

gcloud auth application-default login

ADC=~/.config/gcloud/application_default_credentials.json && \
docker run -p8080:8080 \
-e GOOGLE_APPLICATION_CREDENTIALS=/tmp/keys/secret.json \
-v ${ADC}:/tmp/keys/secret.json \
audio-transcribe-go
```

## Test locally (accept BQ RF [request contract](https://cloud.google.com/bigquery/docs/remote-functions#input_format))
Upload this [sample audio](https://www.voiptroubleshooter.com/open_speech/american.html) (click [here](https://www.cs.columbia.edu/~hgs/audio/harvard.html) for more details) to your GCS bucket and run the local test.
```
Please see examples.http for more details. You can also execute it directly with something like REST Client in VS Code (ext install humao.rest-client).
```

## Run on Cloud Function
```
gcloud functions deploy audio-transcribe-go \
    --gen2 \
    --concurrency=8 \
    --runtime=go122 \
    --region=us-central1 \
    --source=. \
    --entry-point=AudioTranscribe \
    --trigger-http \
    --allow-unauthenticated
```

## Run on Cloud Run
[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)

# Additional notes
* Unit test created by leveraging Gemini.
  
## Related links
* https://cloud.google.com/bigquery/docs/remote-functions
* https://cloud.google.com/functions/docs/concepts/go-runtime
* https://cloud.google.com/docs/buildpacks/build-function
