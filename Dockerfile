FROM cgr.dev/chainguard/go:latest AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o elfexplorer

FROM cgr.dev/chainguard/static:latest
COPY --from=build /app/elfexplorer /app
COPY ./addresses.txt /addresses.txt

EXPOSE 8080

ENTRYPOINT [ "/app" ]