# storage
It is an application that calculates the size of files under a directory and returns them in JSON format.

## How use it?
### Running the server
```sh
$ go run main.go
```
### Build & Run server
```sh
$ go build .
$ storage.exe
```

### Using Docker Compose
The Docker Compose file consists of the following.

```yaml
version: '3'

services: 
  storage:
    build: 
      context: ./
      dockerfile: ./Dockerfile
    image: storage-scale:0.1.0
    container_name: storage-scale
    ports:
      - 8080:8080
    environment:
        # Specify the origin to allow access to the resource
        ACCESS_CONTROL_ALLOW_ORIGIN: "http://localhost"
        # Specifies the base path of the directory to search
        BASE_DIR: "/share/misc"
        # Separate files and paths to exclude from the search, separated by commas
        EXCLUDE: ignore,ignore2
        # If the debug mode is set to true, the standard output shows the number, file name and file size (in bytes).
        # Example.
        #   No.1 .dockerignore 19 byte
        #   No.2 api/swagger.yaml 2420 byte
        IS_DEBUG_MODE: "false"
    # Specify the directory you want to search for in volume
    volumes: 
      - /tmp:/share/misc
```

Start the container with the following command and request URL (`https://localhost :8080/v1/weight`).
```sh
$ docker-compose build
$ docker-compose up -d
```

### cURL
```sh
curl -i -X POST \
   -H "Content-Type:application/json" \
   -H "Accept:application/json" \
   -d \
'{
  "target_path": "/scales",
  "ignore_paths": [
    	".git",
    	".DS_Store",
    	"__debug_bin"
  ]
}' \
 'http://localhost:8080/v1/weight'
```

## Document
Import the following YAML file into the [Swagger Editor](https://editor.swagger.io/).
- [Storage Scales](./api/swagger.yaml)

## LICENSE
[MIT](LICENSE)
