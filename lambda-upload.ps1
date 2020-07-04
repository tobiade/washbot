$env:GOOS = "linux"
$env:CGO_ENABLED = "0"
$env:GOARCH = "amd64"
go build -o main cmd/washbot/main.go
build-lambda-zip.exe -output main.zip main
aws lambda update-function-code --function-name washbot-alexa --zip-file fileb://main.zip