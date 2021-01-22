git:
	git add .
	git commit -m "$(msg)"
	git push origin master

gen: 
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

clean: 
	rm pb/*.go

run:
	go run main.go

test:
	go test -cover -race ./...

.PHONY: git gen clean run test