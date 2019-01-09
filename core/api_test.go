package core_test

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/wcamarao/devtools/core"
)

func TestAPIResponse(t *testing.T) {
	expected := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "Okay",
	}

	assert.Equal(t, expected, core.APIResponse(200, "Okay"))
}
