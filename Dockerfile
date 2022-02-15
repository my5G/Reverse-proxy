FROM golang:1.14.4-stretch

WORKDIR /workspace

RUN git clone https://github.com/my5G/Reverse-proxy.git \
    && cd Reverse-proxy \
    && go mod download


# Move to the binary path
WORKDIR /workspace/Reverse-proxy/cmd

RUN go build -o app

# Config files volume
VOLUME [ "/workspace/Reverse-proxy/config" ]