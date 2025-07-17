FROM --platform=${BUILDPLATFORM} golang:1.24.5-alpine3.22 AS build

ARG TARGETARCH
ENV GOARCH="${TARGETARCH}"
ENV GOOS=linux
WORKDIR /
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod go mod download
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build -ldflags="-w -s" -o /tasks-api

FROM alpine:3.22
ARG ENV_TAG
ENV BUILD_VERSION=${ENV_TAG}
COPY --from=build /tasks-api /tasks-api

EXPOSE 8080
ENTRYPOINT ["/tasks-api"]
