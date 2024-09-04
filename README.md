# Go Application with gRPC and Docker

This project is a Go-based microservice that uses gRPC for communication. The service calculates scores based on ratings in different categories. The project also includes Docker support for containerization.

### Prerequisites
* Go 1.23
* Docker

### Project Structure
This Project uses Go Workspaces to structure the code into separate modules
* **protos** - Contains the ProtoBuffer files and gRPC related code generation.
* **app** - Contains the server application, including the database and gRPC services.
* **client** - Contains a simple test client for quick sandbox testing.

### Protos Module
To update the Protocol Buffer messages, run the following command

```
    make build_proto
```

### Protos Module
To update the Protocol Buffer messages, run the following command

```
    make build_proto
```

### App Module
To run the app module locally, run the following command from the project root directory

```
    go run ./app/main.go
```

The app module is separated into three directories
* **db** Contains the database layer.
* **grpc_service** Contains the gRPC layer.
* **ticket_score_service** Contains the domain layer, including the algorithms and service logic.

### Dockerization

**Build the Docker Image**

```
    docker build -t zenklaus:latest . 
```

**Run the Docker Container

```
    docker run -p 9000:9000 zenklaus:latest
```

Your application should now be accessible at `http://localhost:9000`.

To run a quick test, You can use the client module

```
    go run ./client/client.go
```


### Running Tests

From the project root, run the following command

```
    go test ./app/...
```

