TARGET := jira-to-slack
VERSION := v0.0.0
LDFLAGS := -X main.version=$(VERSION)

all: $(TARGET)

.PHONY: check
check:
	golangci-lint
	go test -v ./...

$(TARGET): $(wildcard *.go)
	go build -o $@ -ldflags "$(LDFLAGS)"

.PHONY: run-appengine
run-appengine:
	dev_appserver.py --port 3000 appengine/app.yaml

dist:
	goxzst -d dist -o "$(TARGET)" -- -ldflags "$(LDFLAGS)"

.PHONY: release
release: dist
	ghr -u "$(CIRCLE_PROJECT_USERNAME)" -r "$(CIRCLE_PROJECT_REPONAME)" "$(CIRCLE_TAG)" dist
