FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

WORKDIR /home/vscode


RUN mkdir .local && \
PB_REL="https://github.com/protocolbuffers/protobuf/releases" && \
curl -LO $PB_REL/download/v30.2/protoc-30.2-linux-x86_64.zip && \
unzip protoc-30.2-linux-x86_64.zip -d .local && \
rm protoc-30.2-linux-x86_64.zip && \
export PATH="$PATH:$HOME/.local/bin" && \
go install github.com/google/wire/cmd/wire@latest && \
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
sudo chown -R vscode:golang /go/pkg