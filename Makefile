
all: proto mocks test

.PHONY: proto
proto:
	rm -rf pkg/apis
	make -C proto protoc

.PHONY: mocks
mocks:
	rm -rf test/mocks
	docker run -it --rm --user $$(id -u):$$(id -g) -w /work -v ${PWD}:/work vektra/mockery:v2.12.3 --keeptree --all --dir pkg/apis --output test/mocks

.PHONY: test
test:
	go test -v ./...