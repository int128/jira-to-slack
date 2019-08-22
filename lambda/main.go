package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/int128/jira-to-slack/pkg/handlers"
	"github.com/int128/jira-to-slack/pkg/jira"
	"github.com/int128/jira-to-slack/pkg/usecases"
)

func handleIndex(_ context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	params, err := handlers.ParseWebhookParams(r.MultiValueQueryStringParameters)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers:    map[string]string{"content-type": "text/plain"},
			Body:       fmt.Sprintf("OK\n%s", err.Error()),
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprintf("OK\nreceived the parameters: %+v", params),
	}, nil
}

func handleWebhook(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	params, err := handlers.ParseWebhookParams(r.MultiValueQueryStringParameters)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers:    map[string]string{"content-type": "text/plain"},
			Body:       err.Error(),
		}, nil
	}
	var event jira.Event
	if err := json.Unmarshal([]byte(r.Body), &event); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers:    map[string]string{"content-type": "text/plain"},
			Body:       fmt.Sprintf("could not decode json of response body: %s", err),
		}, nil
	}
	in := usecases.WebhookIn{
		JiraEvent:       &event,
		SlackWebhookURL: params.Webhook,
		SlackUsername:   params.Username,
		SlackIcon:       params.Icon,
		SlackDialect:    params.Dialect,
	}
	var u usecases.Webhook
	if err := u.Do(ctx, in); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    map[string]string{"content-type": "text/plain"},
			Body:       err.Error(),
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"content-type": "text/plain"},
		Body:       "OK",
	}, nil
}

func handler(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if r.HTTPMethod == "GET" {
		return handleIndex(ctx, r)
	}
	if r.HTTPMethod == "POST" {
		return handleWebhook(ctx, r)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       "Method Not Allowed",
	}, nil
}

func main() {
	lambda.Start(handler)
}
