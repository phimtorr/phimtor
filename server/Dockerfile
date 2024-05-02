FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN cd server && CGO_ENABLED=0 GOOS=linux go build -o server -ldflags "-s -w -X github.com/phimtorr/phimtor/server/version.Version=$VERSION"

FROM scratch

WORKDIR /app

COPY --from=build /app/server .

EXPOSE 8080

CMD ["./server"]

