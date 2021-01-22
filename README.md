# go-grpc

## Install Protobuf compiler [protoc]
```bash
# mac
brew install protobuf

# linux(usage with docker)
PROTOC_ZIP=protoc-3.14.0.-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP
```

## Installing gRPC code generation libs for go
```bash
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go

# generating codes
protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
```

## IDE extension set for proto files
```bash
# configure proto path with pasting following settings
"protoc": {
  "path": "/usr/local/bin/protoc",
  "compile_on_save": false,
  "options": [
    "--proto_path=protos"
  ]
}

# instal clang-format for proto formatting(debian)
apt install clang-format

# ADD FOLLOWING SNIPPET FOR AUTO FORMATTING
"clang-format.executable": "/usr/bin/clang-format",
"[proto3]": {
  "editor.defaultFormatter": "xaver.clang-format",
  "editor.formatOnSave": true
}
```

## Makefile specs
- **git** - git add - commit - push commands
- **gen** - generate go code from proto file
- **clean** - clean all generated code
- **server** - run server main
- **run** - run client main
- **test** - run all package test with coverage

## References
[protobuf](https://github.com/protocolbuffers/protobuf)<br>
[protoc-linux-install](http://google.github.io/proto-lens/installing-protoc.html)<br>
[proto-package-option](https://developers.google.com/protocol-buffers/docs/reference/go-generated#package)<br>
[grpc-io-go](https://grpc.io/docs/languages/go/quickstart/)<br>
[protobuf-famous-types](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf)<br>
[google-uuid](https://github.com/google/uuid)<br>
[golang-copier](https://github.com/jinzhu/copier)<br>

## Author
**Akshit Sadana <akshitsadana@gmail.com>**

- Github: [@Akshit8](https://github.com/Akshit8)
- LinkedIn: [@akshitsadana](https://www.linkedin.com/in/akshit-sadana-b051ab121/)

## License
Licensed under the MIT License