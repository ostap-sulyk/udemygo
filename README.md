# Project README

## Getting Started

To build and run the project:

1. **Build** the application:
   ```bash
   make build
   ```

2. You can simply run the server with:
   ```bash
   make run
   ```

The server will start on `http://localhost:8080`.

---

### Port Configuration
If the port '":8080"' is already in use, modify the port constant in:  

`udemygo/cmd/web/main.go`

```go
const port = ":8080"
```
Change it to an available port (e.g., `":9090"`), then rebuild and rerun.

Here is a simple flow chart:

