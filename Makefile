# phony rules
.PHONY: build
build:
	@docker run --rm -it -v F:\Pitech-bootcamp\Data-Structures-Algorithms/:/go/src test-go bash -c "go build -o go-test main.go"

.PHONY: build-image
build-image:
	@docker build -t=test-go . 

.PHONY: run
run:
	@docker run --rm -it -v F:\Pitech-bootcamp\Data-Structures-Algorithms/:/go/src test-go bash -c "go run main.go"
