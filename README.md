# My Go language laboratory
    > go mod init name
    > go run .
    > go get library
    > go mod tidy

This command specifies that `example.com/greetings` should be replaced with `../greetings` for the purpose of locating the dependency. After you run the command, the `go.mod` file in the hello directory should include a replace directive:
    > go mod edit -replace example/greetings=../greetings
then run this command to synchronize the `example/hello` module's dependencies, adding those required by the code, but not yet tracked in the module:
    > go mod tidy