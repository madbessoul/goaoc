init:
	@if [ -z "$(DAY)" ]; then \
		echo "Error: DAY parameter is required. Usage: make init DAY=02"; \
		exit 1; \
	fi; \
	PADDED_DAY=$$(printf "%02d" $(DAY)); \
	DIR=day$$PADDED_DAY; \
	if [ -d "$$DIR" ]; then \
		echo "Error: $$DIR already exists"; \
		exit 1; \
	fi; \
	mkdir -p $$DIR; \
	sed "s/dayXX/day$$PADDED_DAY/g" tmpl/solve.go > $$DIR/solve.go; \
	sed "s/dayXX/day$$PADDED_DAY/g" tmpl/solve_test.go > $$DIR/solve_test.go; \
	touch $$DIR/input.txt; \
	echo "Created $$DIR with solve.go, solve_test.go, and input.txt"

day-%:
	go test ./day$* -v
