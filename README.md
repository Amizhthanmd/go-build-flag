# go-build-flag

Build the application and pass metadata using ldflags:

```bash
go build -ldflags "-X main.version=1.0.0 -X main.buildDate=$(date +%Y-%m-%d) -X main.commit=$(git rev-parse --short HEAD)" -o app
```
- `-X main.version=1.0.0`: Sets the version variable.
- `-X main.buildDate=$(date +%Y-%m-%d)`: Sets the buildDate variable using the current date.
- `-X main.commit=$(git rev-parse --short HEAD)`: Sets the commit variable to the current Git commit hash.

- The `-ldflags` option injects metadata into the Go binary at build time.
