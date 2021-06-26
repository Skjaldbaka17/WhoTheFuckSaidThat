# WhoTheFuckSaidThat


## Usage

Create .env file in root with:

PORT=PORT
API_KEY=MY_API_KEY
API_BASE_URL=http://quotel-api.com

and run 

```go
go run main.go
```

## AWS Lambda Functions

Uploading each function as .zip:

```zsh
GOARCH=amd64 GOOS=linux go build main.go &&
zip main.zip main 
```
