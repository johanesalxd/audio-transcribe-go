package audiotranscribe

import "github.com/GoogleCloudPlatform/functions-framework-go/functions"

func init() {
	functions.HTTP("AudioTranscribe", AudioTranscribe)
}
