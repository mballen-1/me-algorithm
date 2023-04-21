### M.E Sample Project

# How to compile and run the program

### Using Docker:
1. Install [Docker Desktop](https://www.docker.com/products/docker-desktop/)
2. Go to the project root directory `(me-algorithm)` after cloning it.
2. Using your terminal run `docker build --tag me-algorithm .` from the same location.
3. Run `docker run -i me-algorithm` 
4. Follow the program instructions and copy-paste (or type) the input content you want to test.
- Stop the container when you are done testing (ctrl+c / cmd+c).

### Using a local installation of Go
- Please go to the [Golang](https://go.dev/dl/) releases page and download the 1.19.6 version for your operative
  system.

--- Once Go is installed, please follow the instructions of your preference:
1. #### Using the command line
Run `go run main.go` from the project root directory.
Then copy paste the input content you want to test in the command line.

2. #### Using Goland IDE
- Install and run [Goland](https://www.jetbrains.com/es-es/go/download/#section=windows)
- Open main.go file and click on the run button, then start using the application after a window is prompted.

# How to run the tests of the solution

1. ### Using the command line and a Go local installation
- Follow and complete the instructions from [Using a local installation of Go](#Using-a-local-installation-of-Go). Then run
  `go test ./... ` from the project root directory

2. ### Using Goland IDE
Follow and complete the instructions from [Using a local installation of Go](#Using-a-local-installation-of-Go)
and [Using Goland IDE](#using-goland-ide) then
Configure Goroot and then run tests using the green icon buttons from
- `src/sum2_test.go`
