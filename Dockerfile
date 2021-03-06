FROM golang as build

ENV GOPROXY=https://goproxy.io

ADD . /mithril

WORKDIR /mithril

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine

ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV MysqlDSN=""
ENV GIN_MODE="release"
ENV PORT=3000

RUN mkdir -p /www/conf

WORKDIR /www

COPY --from=build /mithril/api_server /usr/bin/api_server
ADD ./conf /www/conf

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]