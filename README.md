# Validate and look up a bank BSB number

### Step by step
- Create a proto file bsbLookup.proto
    - `syntax`
    - `package`
    - `option go_package`
    - `service`
    - `message`
    
- Create `go.mod` file
    - `git init` a directory, push current project to remote
    - `go mod init <directory path>`
    - add requires
    - ```go get vendor```
  
- Create `config/config.yaml`
    - start with ```---```: indicate the start of a new yaml document
    - include `port: 8080`
  
- Create `config.go` file in `cmd/config/config.go`
    - in type `Config` define field `Port` and `mapstructure:"port"` - tell config file find in `config.yaml` file the field `port`
  
