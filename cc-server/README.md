# Imacry Command & Control Server

**WARNING: these instructions work best on a unix based system**

## **THIS WAS MADE FOR EDUCATIONAL PURPOSES**

## Setup

We use [Earthly](https://earthly.dev) to build files and images

You need Earthly dependencies
* Docker
* Git 

## Building

You can build a docker container or the normal binary

to build the container run 

``` sh
earthly +docker
```

This will create an image called `imacry-cc-server:latest`. This container exposes port 80 so you gotta bind the port

You can also just build the binary with

``` sh
earthly +build-cc-server
```

## Configuration

You can use your own configuration file but one is already provided.

By default, if no configuration path is specified, it will look in `/etc/imacry-cc-server` for `cc-server.yaml`. 

The docker container **ON BUILD** uses `./cc-server.yaml` as its configuration. All options can be overruled by flags or env vars.

## Running

This could be ran by a number of ways

* command line on bare metal
* command line from inside docker container
* systemd launching on bare metal
* systemd launching a docker container

### command line on bare metal

This is just the binary by itself. Im sure you can figure this one out

### command line from inside docker container

``` sh
docker run -p <port num>:80 imacry-cc-server:latest
```
You can bind to whatever port number you want

