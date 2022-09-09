# Debugging

I debug in vscode + [go extension](https://code.visualstudio.com/docs/languages/go#_debugging) + [Delve](https://github.com/go-delve/delve/tree/master/Documentation/installation).

## Prerequisites

Delve: `go install github.com/go-delve/delve/cmd/dlv@latest`

## Launch Profile

```json
// .vscode/launch.json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "go-api",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "go",
            "args": [
                "${workspaceFolderBasename}/main.go",
                "server",
            ],
        }
    ]
}
```