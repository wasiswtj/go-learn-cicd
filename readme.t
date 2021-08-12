## To Run Unit Test and Coverage Test
go test -v -coverprofile=coverage.out ./calculations/handler/ ./calculations/usecase/

## To create report of Unit and Coverage Test
go tool cover -html=coverage.out -o coverage.html