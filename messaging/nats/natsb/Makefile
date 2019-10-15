run:
	go run . -np=5 -ns=5 hello
protogo:
	protoc ./proto/hello/*.proto -I. -I=/usr/local/include \
	-I=~/go/src --go_out=plugins=grpc:~/go/src

gogoproto:
	protoc ./proto/hello/*.proto -I. \
	-I=${GOPATH}/src -I=$(GOPATH)/src/github.com/gogo/protobuf/protobuf --gogofaster_out=plugins=grpc:${GOPATH}/src;