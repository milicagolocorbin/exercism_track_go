# Exercism Track Go - SOLUTIONS

## Organization

Problems are divided in two folders:

1. `learning_exercises` - folder that explores specific Go features. Each exercise has separate folder with solution code file and test code file.
2. `practice_exercises` - folder that reinforces what is learned by solving larger, more open-ended problems. Each exercise has separate folder with solution code file and test code file.

## Run Tests

To run all tests:

```bash
# to run all tests from the level of go.mod file:
go test ./...

# to run all tests from the level of go.mod file with verbose flag:
go test -v ./...

```

To run individual tests:

```bash
# to run all tests in specific folder:
cd learning_exercises/01_hello_world/
go test

# to run all tests in specific folder with verbose flag:
cd learning_exercises/01_hello_world/
go test -v

```
