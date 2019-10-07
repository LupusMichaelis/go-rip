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

run: host $(SRCS)
	go $@ .

test: host $(TESTS) $(SRCS)
	go test -v ./...

format: $(SRCS) $(TESTS)
	go fmt ./...

cert:
	openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
	chmod 400 server.key server.crt

host:
	go get -t ./...
	test -h $(GOROOT)/src/$(VENDOR)/$(notdir $(PWD)) \
		|| ln -s $(PWD) $(GOROOT)/src/$(VENDOR)/
