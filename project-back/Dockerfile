FROM golang:1.17.2-alpine

ENV PORT 8080
ENV APP_NAME myproject
ENV GOPATH /go

RUN apk add --no-cache gcc musl-dev

COPY . /go/src/${APP_NAME}
COPY ./connection /go/src/${APP_NAME}/connection
COPY ./customhttp /go/src/${APP_NAME}/customhttp
COPY ./dataacces /go/src/${APP_NAME}/database

WORKDIR /go/src/${APP_NAME}
RUN go build -o ${APP_NAME}

CMD /go/src/${APP_NAME}/${APP_NAME}

EXPOSE ${PORT}