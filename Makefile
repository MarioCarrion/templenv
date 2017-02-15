NAMESPACE=mariocarrion/templenv

default: build

builddocker:
	docker build --tag ${NAMESPACE}:build --file ./Dockerfile .

build: builddocker
	docker run --tty ${NAMESPACE}:build /bin/true && \
		docker cp `docker ps -q -n=1`:/templenv . && \
		docker rm `docker ps -q -n=1` && \
		docker build --rm --tag ${NAMESPACE}:latest --file ./Dockerfile.static .

gobuild:
	CGO_ENABLED=0 GOOS=linux go build --ldflags="-s" -a -installsuffix cgo \
							-o templenv ./go/src/github.com/MarioCarrion/templenv

test:
	go test github.com/MarioCarrion/templenv/parser -coverprofile=coverage.out && \
		go tool cover -html=coverage.out
