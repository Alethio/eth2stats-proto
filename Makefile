proto:
	protoc -I ./ --go_out=plugins=grpc:./ ./*.proto