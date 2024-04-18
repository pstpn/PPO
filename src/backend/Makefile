all: build-service build-storage build-ui

build-service:
	cd internal/service && go build -o service && mv service ../..

build-storage:
	cd internal/storage && go tool compile -S -I /opt/homebrew/opt/go company.go && mv storage ../..

build-ui:
	cd console && go build -o ui && mv ui ..

go-tests_template:
	gotests -all -w <path>

mockery_template:
	mockery --all --inpackage

go-test_command:
	go test -v ./...

.PHONY: go-tests_template mockery_template go-test_command all build-service build-storage build-ui