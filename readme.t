- Crafted Simple Web API Implementation Using Golang Echo Framework.
- Crafted with db migrations for database version control
- Crafted with automated unit and coverage test on handler, usecase, and repository level on respective domain
- Crafted with Docker Container and CI Script to run build and test the apps

## To Run Unit Test and Coverage Test
go test -v -coverprofile=coverage.out ./calculations/handler/ ./calculations/usecase/

## To create report of Unit and Coverage Test
go tool cover -html=coverage.out -o coverage.html
