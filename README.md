# Validate and look up a bank BSB number

### Step by step
- Create a proto file bsbLookup.proto
    - `syntax`
    - `package`
    - `option go_package`
    - `service`
    - `message`
    - run `protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      proto/bsbLookup.proto` to create `bsbLookup_grpc.pb.go`
      
- Create `go.mod` file
    - `git init` a directory, push current project to remote
    - `go mod init <directory path>`
    - add requires
    - ```go mod vendor```
    - go to File -> Setting -> Go -> Go Module -> Enable Go Module integration
  
- Create `config/config.yaml` which has port number
    - start with ```---```: indicate the start of a new yaml document
    - include `port: 8080`
  
- Create `config.go` file in `cmd/config/config.go`
    - in type `Config` define field `Port` and `mapstructure:"port"` - tell config file find in `config.yaml` file the field `port`
  
- Use config in `main.go`
    - firstly just finishing load config file
  
    ```go
        configs, err := config.LoadConfig()
	    if err != nil {
		    fmt.Printf("Error loading configs file %s\n", err)
	    }
	    port := configs.Port
	    fmt.Printf("bsb-lookup started on port %v\n", port)
    ```
  
- Create `service/service.go`
    - create a new type that implements `bsbLookupServer` interface (which can be found in `bsbLookup_grpc.pb.go`)
    - define again `Validate` method
  
- Write test file `service_test.go`

- Come back with `main.go`
    - call `net.Listen` - listen in port 8080
    - create a grpc server and register lookupAPI service (`api`) to that server.
    - let server sever the call

- Write `tesh.sh`
    - Start with `#!/bin/sh`
    - `grpcurl`: give the path where can look for the file to execute (ex. bsbLookup.bsbLookup.Validate)
    - need to `cd ../`: go out of the `scripts` folder, go to main project to run the command
  
- Write test file for client site `blackbox_test`
    - `cc, err := grpc.Dial (target: "localhost:8080", grpc.WithInsecure())`
    Target is `localhost:8080` as the lookup service and blackbox test are running in the same localhost
      
- Create Dockerfile for service
    - Dockerfile start with `FROM golang:1.16.3-alpine3.13` - to inherit the base from golang alpine
    - Then other commands
    - build image: `docker build -t bsblookup .`
  At this stage, if we want to run docker for the new image, and then run `blackbox_test.go`, we need to specify the port 8080 in run command
  `docker run -d -p 8080:8080 bsblookup`
      
- Create Dockerfile for blackbox
    - `RUN go test -c ./testing/blackbox -o newblackbox`
    - `CMD["/app/newblackbox"]`
    - build image: `docker build -t blackbox -f ./testing/blackbox/Dockerfile .`
    - at this stage, if we create container for blackbox, we are dialling to `localhost:8080`, which is `docker0` of the blackbox image itself.
  We need to connect the 2 images toghether, by giving it IP address. To get the IP address of `bsblookup` container, we will inpsect its by name
      `docker ps` --> to get the name of container, which has image `bsblookup`
      `docker inspect <container's name>` --> get IPAddress ("172.17.0.2")
      change `localhost:8080` in `blackbox_test.go` file to `172.17.0.2:8080`
      build image blackbox again
      run : `docker run blackbox`
      
- Create `docker-compose.yml`
    - build docker-compose: `docker-compose build bsblookup blackbox`
    - run docker-compose: `docker-compose up blackbox`

