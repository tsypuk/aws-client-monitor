package validator

import (
	"aws-client-monitor/internal/domain"
	"fmt"
)

func ValidateApiCall(apiCall *domain.ApiCall) error {
	if apiCall.Type != "ApiCall" {
		return fmt.Errorf("invalid api call type: %s", apiCall.Type)
	}
	return nil
}
