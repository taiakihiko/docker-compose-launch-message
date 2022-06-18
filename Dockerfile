FROM golang:1.17.8

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y vim

RUN echo "alias ll='ls -l'" >> ~/.bashrc

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

EXPOSE 8080
