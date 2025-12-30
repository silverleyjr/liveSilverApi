FROM ubuntu:latest
COPY silverApi /silverApi
RUN apt-get update
RUN apt install -y curl 
RUN curl https://go.dev/dl/go1.22.1.linux-amd64.tar.gz 
RUN tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz 
RUN export PATH=$PATH:/usr/local/go/bin 
RUN cd silverApi
RUN go build
CMD ["go run cmd/api/main.go"]
