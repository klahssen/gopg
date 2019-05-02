export GOOS=linux
export GOARCH=amd64

.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get .

test: install
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o hellokube-go .

serve: build
	./hellokube-go

clean:
	rm ./hellokube-go

pack:
	GOOS=linux make build
	docker build -t twiggg/hellokube-go:$(TAG) .

upload:
	gcloud docker -- push twiggg/hellokube-go:$(TAG)

deploy:
	envsubst < k8s/deployment.yaml | kubectl apply -f -

ship: test pack upload deploy clean