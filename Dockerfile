FROM golang:1.23

LABEL maintainer="By Mohsen Taheri"
LABEL Email="m.rozbehano@outlook.com"

RUN apt update && apt install vim -y

ENV APP_ENV=test
ENV PORT=9090
ENV IP="0.0.0.0"

EXPOSE 9090

WORKDIR app

RUN mkdir -p ./config/file

COPY ./cmd/rest_service .
COPY ./config/file ./config/file

CMD ["./rest_service"]