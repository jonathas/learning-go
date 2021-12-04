# Writing tests in Go

Show coverage in the terminal:
go test -cover

Generate coverage report:
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
