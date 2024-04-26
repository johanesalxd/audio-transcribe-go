export FUNCTION_TARGET=AudioTranscribe

curl -m 60 -X POST localhost:8080 \
-H "Content-Type: application/json" \
-d '{
  "requestId": "",
  "caller": "",
  "sessionUser": "",
  "userDefinedContext": {},
  "calls": [
    ["uri_1"],
    ["uri_2"],
    ["uri_n"]
  ]
  }'