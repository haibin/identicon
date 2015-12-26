# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/haibin/identicon

WORKDIR /go/src/github.com/haibin/identicon

RUN go get github.com/tools/godep

RUN godep go install github.com/haibin/identicon

# Run the outyet command by default when the container starts.
CMD /go/bin/identicon

# Document that the service listens on port 5000.
EXPOSE 5000