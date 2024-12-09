# OptiPix Annotation Service

This Annotation Service is a WebSocket-based service designed to handle real-time communication for chat and annotation functionalities. It provides a lightweight and efficient server implementation to support WebSocket connections, enabling real-time message exchange and data synchronization.

## Requirements

- **Go**: Version 1.20 or higher.

## Launch
go run cmd/main.go 

## Directory Structure

```
```
# Directory Structure

project/
- cmd/  # Service entry directory
  - main.go  # Main service entry point
- internal/  # Internal module directory
  - annotation/  # Annotation logic
    - handler.go  # Annotation WebSocket handling
  - chat/  # Chat logic
    - handler.go  # Chat WebSocket handling
- pkg/  # Reusable components
  - websocket/  # WebSocket utilities and wrappers
  - database/  # Database utilities and wrappers
  - config/  # Config utilities and wrappers
- configs/  # Configuration directory
  - config.yaml  # Configuration file
- sql/  # SQL Schema
  - optipix.sql  # SQL Schema file
- scripts/  # SQL Schema
  - build.sh  # Compile script file
  - run.sh  # Run script file
- go.mod  # Go module file
- README.md  # Documentation
```

