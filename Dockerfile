# Build the manager binary
FROM golang:1.17.7 as builder

WORKDIR /workspace

# Run this with docker build --build_arg $(go env GOPROXY) to override the goproxy
ARG goproxy=https://proxy.golang.org
ENV GOPROXY=$goproxy

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY pkg/ pkg/

# Build
ARG LDFLAGS
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -a -trimpath -ldflags "${LDFLAGS} -extldflags '-static'" \
  -o manager .

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER nobody

ENTRYPOINT ["/manager"]
