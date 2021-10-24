
# getting the dependencies
deps:
	FROM golang:1.16
	WORKDIR /app
	COPY go.mod go.sum .
	RUN go mod download

# builds the client binary
build:
	# build off of deps stage
	FROM +deps
	# copy main file
	COPY main.go .
	# copy cmd and whatever else in dir mode
	# this is like `cp -r`
	COPY --dir cmd/ pkg/ ./
	# build to file `imacry`
	RUN go build -o imacry main.go
	# save file to outside container
	SAVE ARTIFACT imacry

# gets the binary from build and then saves it to local machine
save-binary:
	FROM scratch
	COPY +build/imacry imacry
	SAVE ARTIFACT imacry AS LOCAL imacry

# make docker container
docker:
	FROM ubuntu:21.04
	COPY +build/imacry imacry

	SAVE IMAGE imacry-run:latest
