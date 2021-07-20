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
    - firstly just until load config file
  
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
    - dddddd

- Write `tesh.sh`
    - Start with `#!/bin/sh`
    - `grpcurl`: give the path where can look for the file to execute (ex. bsbLookup.bsbLookup.Validate)