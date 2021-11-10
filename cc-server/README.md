# Imacry Command & Control Server

**WARNING: these instructions work best on a unix based system**

## **THIS WAS MADE FOR EDUCATIONAL PURPOSES**

## Setup

We use [Earthly](https://earthly.dev) to build files and images

You need Earthly dependencies
* Docker
* Git 

## Building

## Configuration
You can use your own configuration file but one is already provided.

By default, if no configuration path is specified, it will look in `/etc/imacry-cc-server` for `cc-server.yaml`. 

The docker container **ON BUILD** uses `./cc-server.yaml` as its configuration. All options can be overruled by flags or env vars.

## Running
This could be ran by a number of ways

* command line on bare metal
* command line from inside docker container
* systemd launching a on bare metal
* systemd launching a docker container

### 
