.PHONY: test
test:
	mkdir build
	go test -coverprofile=build/cover.out ./...

.PHONY: clean
clean:
	go clean
	rm -rf build

.PHONY: env
env:
	env
