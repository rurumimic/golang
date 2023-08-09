# Install

## Go Version Manager

```bash
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

### Restart a terminal session

`~/.zshrc`:

```bash
# Languages
## Go
source $HOME/.gvm/scripts/gvm
```

### Install

```bash
gvm listall
gvm install go1.21.0 --with-protobuf --with-build-tools --prefer-binary
gvm use go1.21.0 --default
```

#### on Apple Silicon

```bash
sudo port install go
# install protobuf3-cpp in ~/.local/bin & ~/.local/include : grpc.io/docs/protoc-installation
# export PATH="$PATH:$HOME/.local/bin"
gvm install go1.21.0 --with-protobuf --with-build-tools
gvm use go1.21.0 --default
sudo port uninstall go
```

### List installed

```bash
gvm list

gvm gos (installed)

=> go1.21.0
```

---

## Visual Studio Code Extension

Extension: [Go](https://marketplace.visualstudio.com/items?itemName=golang.go)

```bash
go env

GOPATH="~/.gvm/pkgsets/go1.21.0/global"
GOROOT="~/.gvm/gos/go1.21.0"
```

**Preferences: Configure Language Specific Settings** `⇧⌘P` → Go

`settings.json`

```json
  "go.gopath": "~/.gvm/pkgsets/go1.21.0/global",
  "go.goroot": "~/.gvm/gos/go1.21.0",
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
Tools environment: GOPATH=~/.gvm/pkgsets/go1.21.0/global
Installing 1 tool at ~/.gvm/pkgsets/go1.21.0/global/bin in module mode.
  gopls
  gopkgs
  go-outline
  dlv
  staticcheck

Installing golang.org/x/tools/gopls (~/.gvm/pkgsets/go1.21.0/global/bin/gopls) SUCCEEDED
# ...

All tools successfully installed. You are ready to Go :).
```

---

## SpaceVim

- [Use Vim as a Go IDE](https://spacevim.org/use-vim-as-a-go-ide/)

### Configuration

```toml
[[layers]]
  name = "format"

[[layers]]
  name = "lang#go"
```

- go run: `SPC l r`
- go build: `SPC l b`
- go test: `SPC l t`
- code coverage: `SPC l c`
- gofmt:`SPC b f`

### Install packages

in `nvim`

```bash
:GoInstallBinaries
```
