FROM golang:1.23

LABEL maintainer="By Mohsen Taheri"
LABEL Email="m.rozbehano@outlook.com"

RUN apt update && apt install vim -y

ENV APP_ENV=test
ENV PORT=8585
ENV IP="0.0.0.0"

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

EXPOSE 8585

WORKDIR $GOPATH/src/usersrv

RUN mkdir -p ./config/file

COPY ./cmd/usersrv .

COPY ./config/file/* ./config/file

CMD ["./usersrv"]