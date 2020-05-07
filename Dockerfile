FROM golang:1.14-alpine as builder

# run it before so that we can re-use the cached layer
RUN apk add build-base
WORKDIR /src
ADD . /src

RUN make binary

FROM gcr.io/distroless/static
COPY --from=builder /src/build/app /app
ENTRYPOINT ["/app"]
