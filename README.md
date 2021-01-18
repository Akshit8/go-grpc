# go-grpc

## Install Protobuf compiler [protoc]
```bash
# mac
brew install prtobuf

# linux(usage with docker)
PROTOC_ZIP=protoc-3.14.0.-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP
```

## Installing gRPC code generation libs for go
```bash
go get google.golang.org/grpc
go get 
```

## Makefile specs
- **git** - git add - commit - push commands


## References
[protobuf](https://github.com/protocolbuffers/protobuf)

## Author
**Akshit Sadana <akshitsadana@gmail.com>**

- Github: [@Akshit8](https://github.com/Akshit8)
- LinkedIn: [@akshitsadana](https://www.linkedin.com/in/akshit-sadana-b051ab121/)

## License
Licensed under the MIT License