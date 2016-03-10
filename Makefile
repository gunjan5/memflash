default: image

get:
	goimports -w .
	go get -t -d -v ./...

src:
	docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app golang:1.5.1 sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o app'

fmt:
	gofmt -w .
	#TODO: go lint, go vet
test:
	go test -v -race ./...
	go test -cover -v ./...

image:
	docker build -t gunjan5/memflash .

run:
	docker run --rm -it gunjan5/memflash

build:
	GOOS=linux GOARCH=amd64 go build -o app

depsave:
	godep save
	
depupdate:
	go get -t -v ./...
	godep update

mongo:
	docker run -p 27017:27017 -d mongo

mem:
	docker run -p 11211:11211 -d memcached

stats:
	docker run -d -p 80:80 -p 2003-2004:2003-2004 -p 2023-2024:2023-2024 -p 8125:8125/udp -p 8126:8126 hopsoft/graphite-statsd
