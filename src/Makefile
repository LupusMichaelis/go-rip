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
	openssl \
		req \
		-subj '/CN=lupusmic.org/O=None/C=FR' \
		-newkey rsa:2048 \
		-nodes -keyout server.key \
		-x509 -days 365 \
		-out server.crt
	chmod 400 server.key server.crt

host:
	test -h $(GOPATH)/src/$(VENDOR)/rip \
		|| ln -s $(PWD) $(GOPATH)/src/$(VENDOR)/rip
	go get -t ./...

format: $(SRCS) $(TESTS)
	go fmt ./...

test: host $(TESTS) $(SRCS)
	go test -v ./...
