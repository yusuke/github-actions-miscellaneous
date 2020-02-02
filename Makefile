.PHONY: test
test:
	if [ ! -d build ]; then mkdir build; fi
	go test -coverprofile=build/cover.out ./...

.PHONY: clean
clean:
	go clean
	if [ -d build ]; then rm -rf build ; fi

.PHONY: env
env:
	env | sort

.PHONY: cj
cj:
	go build -o build/cj cj/main.go
