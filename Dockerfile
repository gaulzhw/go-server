# build image, run this with docker build --build-arg builder_image=<golang:x.y.z>
ARG builder_image
FROM ${builder_image} as builder
WORKDIR /workspace

# Build
ARG package=.

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 go build -ldflags=${ldflags} -o manager ${package}

# runtime image
FROM alpine:3
WORKDIR /bin
COPY --from=builder /workspace/manager .

CMD ["manager"]
