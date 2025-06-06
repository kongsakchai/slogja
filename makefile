.PHONY: example
example:
	go run ./example

.PHONY: bench-uber
bench-uber:
	go test -bench=. -benchmem ./benchmark/uber/

.PHONY: bench-golog
bench-golog:
	go test -bench=. -benchmem ./benchmark/golog/
