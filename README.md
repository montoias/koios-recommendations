# koios-recommendations

## About

Micro-service for Koios recommendations.

## Stack:
* GoLang
* Docker

## Development

### Install dependencies
* Install [Docker](https://beta.docker.com/docs)

### Running
```bash
# to startup a shell in the container
$ ./environment.sh
```

### Building the binary
```bash
$ cd delivery/api && go build -v
```

### Run the code
To run without building the binary:
```bash
$ go run delivery/api/main.go
```

### Testing
To run all test cases in the project except the vendor directory
```bash
$ ./test.sh
```

### Dependencies
This project uses [tools/godep](https://github.com/tools/godep) package to as a dependency manager tool.

If adding any new dependencies, just run the following command and push the new files added to your vendor directory.

```bash
$ godep save $(go list ./... | grep -v /vendor/)
```
