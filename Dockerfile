FROM cgr.dev/chainguard/go:latest AS build

WORKDIR /work

COPY ./ .
RUN go build -o app .

FROM cgr.dev/chainguard/static:latest

COPY --from=build /work/app /app
CMD ["/app"]