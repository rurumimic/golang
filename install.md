# Install

1. Install [gvm](https://github.com/moovweb/gvm)
2. List all Go versions: `gvm listall`
3. Install go: `gvm install goX.Y.Z --binary`
4. Go versions: `gvm list`
5. Select: `gvm use goX.Y.Z`
6. Env vars: `go env`

## Go Version Manager

```bash
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

### Restart a terminal session

```bash
source /Users/$USER/.gvm/scripts/gvm
```

### Install

```bash
gvm install go1.16 -B
gvm alias create 16 go1.16
gvm use 16 # --default
```

### List installed

```bash
gvm list

gvm gos (installed)

   16
=> go1.16
```

---

## Visual Studio Code Extension

Extension: [Go](https://marketplace.visualstudio.com/items?itemName=golang.go)

```bash
go env

GOPATH="/Users/$USER/.gvm/pkgsets/go1.16/global"
GOROOT="/Users/$USER/.gvm/gos/go1.16"
```

**Preferences: Configure Language Specific Settings** `⇧⌘P` → Go

`settings.json`

```json
  "go.gopath": "/Users/$USER/.gvm/pkgsets/go1.16/global",
  "go.goroot": "/Users/$USER/.gvm/gos/go1.16",
  "[go]": {
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    },
    "editor.suggest.snippetsPreventQuickSuggestions": false,
    "editor.defaultFormatter": "golang.go"
  },
  "gopls": {
    "experimentalWorkspaceModule": true
  }
```

### gopls

VS Code should handle that step for you.

```bash
Tools environment: GOPATH=/Users/$USER/.gvm/pkgsets/go1.16/global
Installing 1 tool at /Users/$USER/.gvm/pkgsets/go1.16/global/bin in module mode.
  gopls
  gopkgs
  go-outline
  dlv
  staticcheck

Installing golang.org/x/tools/gopls (/Users/$USER/.gvm/pkgsets/go1.16/global/bin/gopls) SUCCEEDED
# ...

All tools successfully installed. You are ready to Go :).
```
