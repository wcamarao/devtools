package core

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

const OffsetMaxLimit = 100

// APIResponse creates an APIGatewayProxyResponse with status code and body
func APIResponse(statusCode int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: body,
	}
}

// ParseOffsetOptions parses an APIGatewayProxyRequest into OffsetOptions
func ParseOffsetOptions(proxyReq events.APIGatewayProxyRequest) (*OffsetOptions, error) {
	limit, err := strconv.ParseUint(proxyReq.QueryStringParameters["limit"], 10, 64)
	if err != nil {
		return nil, errors.New("Invalid limit")
	}
	if limit > OffsetMaxLimit {
		return nil, fmt.Errorf("Limit must not exceed %d", OffsetMaxLimit)
	}

	offset, err := strconv.ParseUint(proxyReq.QueryStringParameters["offset"], 10, 64)
	if err != nil {
		return nil, errors.New("Invalid offset")
	}

	return &OffsetOptions{
		GroupBy: proxyReq.QueryStringParameters["groupBy"],
		OrderBy: proxyReq.QueryStringParameters["orderBy"],
		Limit:   limit,
		Offset:  offset,
	}, nil
}

// ParsePathID parses an APIGatewayProxyRequest into an entity ID
func ParsePathID(proxyReq events.APIGatewayProxyRequest) (string, error) {
	if len(proxyReq.PathParameters["id"]) != 53 {
		return "", errors.New("Invalid ID")
	}

	return proxyReq.PathParameters["id"], nil
}
