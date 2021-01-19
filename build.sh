rm app.zip
go test ./...
GOARCH=amd64 GOOS=linux go build -o app cmd/application.go
zip -r app.zip app
rm app