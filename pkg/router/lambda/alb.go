package lambda

import (
	"context"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/int128/jira-to-slack/pkg/handlers"
)

func ALBTargetGroup(ctx context.Context, r events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
	switch {
	case r.HTTPMethod == "GET":
		var h handlers.Index
		code, body, err := h.Serve(r.MultiValueQueryStringParameters)
		if err != nil {
			return events.ALBTargetGroupResponse{
				StatusCode: http.StatusOK,
				Headers:    map[string]string{"content-type": "text/plain"},
				Body:       err.Error(),
			}, nil
		}
		return events.ALBTargetGroupResponse{
			StatusCode: code,
			Headers:    map[string]string{"content-type": "text/plain"},
			Body:       body,
		}, nil

	case r.HTTPMethod == "POST":
		var h handlers.Webhook
		code, err := h.Serve(ctx, r.MultiValueQueryStringParameters, strings.NewReader(r.Body))
		if err != nil {
			return events.ALBTargetGroupResponse{
				StatusCode: code,
				Headers:    map[string]string{"content-type": "text/plain"},
				Body:       err.Error(),
			}, nil
		}
		return events.ALBTargetGroupResponse{
			StatusCode: code,
			Headers:    map[string]string{"content-type": "text/plain"},
			Body:       "OK",
		}, nil

	default:
		return events.ALBTargetGroupResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Headers:    map[string]string{"content-type": "text/plain"},
			Body:       "Method Not Allowed",
		}, nil
	}
}
