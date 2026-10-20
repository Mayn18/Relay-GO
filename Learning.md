# Learning Go — Today's Notes

### 1. Serving Static Files

In Go, we can use `http.FileServer()` to serve static files from a directory.

```go
http.FileServer(http.Dir("./web"))
```

- `http.Dir("./web")` specifies the directory containing the files.
- `http.FileServer()` creates an HTTP handler that serves files from that directory.

### 2. Starting an HTTP Server

We can start a server and listen on a specific port using `http.ListenAndServe()`.

```go
http.ListenAndServe(":8080", fileServer)
```

- `:8080` specifies port 8080.
- `fileServer` handles incoming HTTP requests.
- `log.Fatal()` can be used to log an error and exit if the server fails to start or stops unexpectedly.

### 3. Go's Standard Library

The `net/http` package in Go's standard library provides functionality for building HTTP servers and clients.

We don't need to install an external package to serve static files.

### 4. Serving `index.html` Automatically

When using `http.FileServer(http.Dir("./web"))`, we generally don't need to specify `index.html` explicitly.

When a browser requests `/`, the file server serves `index.html` from the root of the `web` directory if it exists.

For example:

```text
project/
├── main.go
└── web/
    ├── index.html
    ├── style.css
    └── script.js
```

- `/` → `index.html`
- `/style.css` → `style.css`
- `/script.js` → `script.js`

**Important:** Pass the directory path (`./web`) to `http.Dir()`, not the path to `index.html` itself.
