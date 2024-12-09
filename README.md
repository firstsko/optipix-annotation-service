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

Annnotation Data Sample:
```
{
    "project_id": "1",
    "part_id": "1001",
    "img": "../src/assets/drew/draw-any-autocad-mechanical-drawings-for-you-1.jpg",
    "tag": "mechanical part",
    "title": "This is a mechanical part, with precise dimensions for fitting and securing two components together.",
    "authors": [
        {
            "name": "Hui Tu",
            "avatar": "/static/images/avatar/3.jpg"
        }
    ],
    "annotations": [
        {
            "id": 1,
            "x": 420,
            "y": 534.6666564941406,
            "width": 146,
            "height": 79,
            "imageWidth": 1152,
            "imageHeight": 600
        },
        {
            "id": 2,
            "x": 708,
            "y": 502.6666564941406,
            "width": 105,
            "height": 73,
            "imageWidth": 1152,
            "imageHeight": 600
        },
        {
            "id": 3,
            "x": 916,
            "y": 454.6666564941406,
            "width": 82,
            "height": 82,
            "imageWidth": 1152,
            "imageHeight": 600
        }
    ]
}
```