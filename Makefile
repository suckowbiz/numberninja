all: tests tidy binary

.PHONY: test
test: tests

.PHONY: tests
tests:
	go test ./...

.PHONY: tidy
tidy:
	go mod tidy
	go mod verify

# Build with:
# - a  						to force build
# - ldflags
#	-w 						do not include debug information to keep file size low
#	-extldflags "-static"	static binding
.PHONY: binary
binary:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-a \
		-ldflags '-w -extldflags "-static"' \
		-o NumberNinja \
		.
	tar -cf NumberNinja.tar NumberNinja