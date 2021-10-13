# Imacry

**WARNING: these instructions work best on a unix based system**

## **THIS WAS MADE FOR EDUCATIONAL PURPOSES**

## Setup

We use [Earthly](https://earthly.dev) to build files and images

You need Earthly dependencies
* Docker
* Git 

## Building

``` sh
earthly +client-build
```

## Testing

Because this is a ransomware bot, we highly suggest you run inside a container

To build the container, run

``` sh
earthly +docker
```

This stage creates an image called `imacry-run:latest`

and then to run

``` sh
docker run -it imacry-run:latest /bin/bash
```

## Contributing
You can read all about contributing to this project in `CONTRIBUTING.md`

## Architecture
You can read about it in `ARCHITECTURE.md`
