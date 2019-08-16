TARGET := jira-to-slack
VERSION := v0.0.0
LDFLAGS := -X main.version=$(VERSION)

all: $(TARGET)

$(TARGET): $(wildcard *.go)
	GOOS=linux GOARCH=amd64 go build -o $@ -ldflags "$(LDFLAGS)"

.PHONY: run
run: $(TARGET)
	sam local start-api

packaged.yaml: $(TARGET)
	sam package --template-file template.yaml --output-template-file $@ --s3-bucket $(SAM_S3_BUCKET_NAME)

.PHONY: deploy
deploy: packaged.yaml
	sam deploy --template-file $< --stack-name $(TARGET) --capabilities CAPABILITY_IAM
