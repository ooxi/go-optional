.PHONY:	all
all:	clean		\
	download	\
	tidy		\
	build		\
	test


.PHONY:	clean
clean:
	@if [ -f 'go.sum' ]; then	\
		rm 'go.sum';		\
	fi

	go clean -testcache


.PHONY:	download
download:
	go mod download


.PHONY:	tidy
tidy:
	go mod tidy
	go fmt ./...


.PHONY:	build
build:
	go build ./...


.PHONY:	test
test:
	go test ./...

.PHONY:	doc
doc:
	godoc -http=:6060

