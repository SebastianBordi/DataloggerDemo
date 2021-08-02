FROM golang:1.16.3-alpine AS build
RUN apk add git
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/google/uuid
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/websocket
RUN go get gorm.io/driver/mysql
RUN go get gorm.io/gorm
WORKDIR /src
COPY ./src .
RUN go build -o /out .

FROM alpine:3.13.5 AS bin
COPY --from=build /out /bin/app
RUN ["mkdir", "/ftp"]
CMD ["/bin/app"]
# docker build . -t docker.emconsol.com/dlm:<version>