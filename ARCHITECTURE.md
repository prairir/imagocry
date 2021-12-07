# Architecture

**This document describes the high-level architecture of this project**

If you want to familiarize yourself with the code base and _generally_ how it works, this is a good place to be.

## High Level TLDR

This is a simple state machine. It has 5 major states, init, encrypt, wait, decrypt, and exit. These states corespond to a specific action or series of actions to be performed.

Command & Control Server is a very basic web server with websocket event loop.

## Sequence Diagram

![Sequence Diagram](pictures/imacry_sequence_diagram.png)

## Code Map

#### Code Map Legend

`<file name>` for a file name

`<folder name>/` for a folder

`<folder name>/<file name>` for a file within a folder

### `README.md`

Readme of the project, it goes over a bunch of things

### `CONTRIBUTING.md`

Documentation about contributing to this project

### `Earthfile`

Earthly is a build system that uses containers to create reporducible environments. This is our configuration file for Earthly. It can do tons of things like building the binary, building the docker container, download dependencies, etc

### `LICENSE`

Were using MIT license

### `main.go`

Base of our code. It launches the CMD system

### `cmd/`

Command line interface module. It pretains to handling the command line interface and all the logic and stuff that goes with it.

### `cmd/root.go`

Root command. This launches the state machine. It also handles the flags as well as the `-h` message

### `cmd/encrypt.go`

Encrypt subcommand. This just encrypts the file system starting at config.Config.Base. It requires `config.Config.Password` to be setup.

### `cmd/decrypt.go`

Decrypt subcommand. This just decrypts the file system starting at config.Config.Base. It requires `config.Config.Password` to be setup.

### `pkg/config/`

Configuration system. It holds the constants and types as well as a central configuration struct.

### `pkg/config/config.go`

This is the file for above.

### `pkg/walk/`

File system walking module. It has an interface `FileAction` which is the file action to happen on every file that isnt a directory `walk.Walk` encounters.

### `pkg/walk/walk.go`

This is the file for above.

### `pkg/encryptfile/`

File encryption system. It has a struct which implements `walk.FileAction`.

### `pkg/encryptfile/encryptfile.go`

This is the file for above.

### `pkg/decryptfile/`

File decryption system. It has a struct which implements `walk.FileAction`.

### `pkg/decryptfile/decryptfile.go`

This is the file for above.

### `pkg/state/`

States for state machine.

### `pkg/state/init.go`

Initialize connection with cc-server, get password if it isnt supplied, launch heartbeat system, and move to next state.

### `pkg/state/encrypt.go`

Encrypt state. Encrypt file system starting at `config.Config.Base`

### `pkg/state/decrypt.go`

Decrypt state. Decrypt file system starting at `config.Config.Base`

### `pkg/state/wait.go`

Wait state. Continuesly waits until either `config.Config.Signal` or `config.Config.HBError` can be read in. If its `config.Config.HBError` then return the error.

### `pkg/state/exit.go`

Exit state. Close `config.Config.Conn`.

### `cc-server/`

Command and control server folder.

### `cc-server/README.md`

Readme for command and control server

### `cc-server/Earthfile`

Earthfile for command and control server.

### `cc-server/main.go`

Base of command and control server. It launches the cmd system.

### `cc-server/cmd/`

Command line interface module. It pretains to handling the command line interface and all the logic and stuff that goes with it.

### `cc-server/cmd/root.go`

File for above.

### `cc-server/web/web.go`

Web server for websockets. Has the event loop for connection handling.

### `cc-server/pkg/config/`

Configuration system.

### `cc-server/pkg/config/config.go`

File for above.

### `cc-server/pkg/handler/`

Event handlers for specific events.

### `cc-server/pkg/handler/init.go`

Init event handler. Responds with the `config.Config.Password`

### `cc-server/pkg/handler/heartbeat.go`

Heartbeat event handler. Responds whether in trigger period.

### `cc-server/services/`

Systemd services for docker container or binary.
