.PHONY: \
	cert \
	format \
	host
	run \
	test \

VENDOR=lupusmic.org

SRCS=\
	main.go \

TESTS=\

build: host $(SRCS)
	go $@ .

run: build
	go $@ .

cert:
	openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
	chmod 400 server.key server.crt

host:
	go get -t ./...
	test -h $(GOPATH)/src/$(VENDOR)/$(notdir $(PWD)) \
		|| ln -s $(PWD) $(GOPATH)/src/$(VENDOR)/

format: $(SRCS) $(TESTS)
	go fmt ./...

test: host $(TESTS) $(SRCS)
	go test -v ./...
