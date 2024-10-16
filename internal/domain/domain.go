package domain

import (
	"fmt"
	"github.com/goccy/go-json"
	"net"
)

type UdpPayload struct {
	UDPAddr *net.UDPAddr
	Payload []byte
}

type ApiCallAttempt struct {
	Version             int    `json:"Version"`
	ClientId            string `json:"ClientId"`
	Type                string `json:"Type"`
	Service             string `json:"Service"`
	Api                 string `json:"Api"`
	Timestamp           int64  `json:"Timestamp"`
	AttemptLatency      int    `json:"AttemptLatency"`
	Fqdn                string `json:"Fqdn"`
	UserAgent           string `json:"UserAgent"`
	AccessKey           string `json:"AccessKey"`
	Region              string `json:"Region"`
	SessionToken        string `json:"SessionToken"`
	HttpStatusCode      int    `json:"HttpStatusCode"`
	XAmznRequestId      string `json:"XAmznRequestId"`
	AwsException        string `json:"AwsException,omitempty"`
	AwsExceptionMessage string `json:"AwsExceptionMessage,omitempty"`
}

func NewApiCallAttempt(payload UdpPayload) (*ApiCallAttempt, error) {
	var apiCallAttempt ApiCallAttempt
	if err := json.Unmarshal(payload.Payload, &apiCallAttempt); err != nil {
		return nil, fmt.Errorf("unmarshal payload: %w", err)
	}
	return &apiCallAttempt, nil
}

func (apiCallAttempt *ApiCallAttempt) Validate() error {
	if apiCallAttempt.Type != "ApiCallAttempt" {
		return fmt.Errorf("invalid api call type: %s", apiCallAttempt.Type)
	}
	return nil
}

// Struct for the ApiCall message
type ApiCall struct {
	Version                  int    `json:"Version"`
	ClientId                 string `json:"ClientId"`
	Type                     string `json:"Type"`
	Service                  string `json:"Service"`
	Api                      string `json:"Api"`
	Timestamp                int64  `json:"Timestamp"`
	AttemptCount             int    `json:"AttemptCount"`
	Region                   string `json:"Region"`
	UserAgent                string `json:"UserAgent"`
	FinalHttpStatusCode      int    `json:"FinalHttpStatusCode"`
	Latency                  int    `json:"Latency"`
	MaxRetriesExceeded       int    `json:"MaxRetriesExceeded"`
	FinalAwsException        string `json:"FinalAwsException,omitempty"`
	FinalAwsExceptionMessage string `json:"FinalAwsExceptionMessage,omitempty"`
}

func NewApiCall(payload UdpPayload) (*ApiCall, error) {
	var apiCall ApiCall
	if err := json.Unmarshal(payload.Payload, &apiCall); err != nil {
		return nil, fmt.Errorf("unmarshal payload: %w", err)
	}
	return &apiCall, nil
}

func (apiCall *ApiCall) Validate() error {
	if apiCall.Type != "ApiCall" {
		return fmt.Errorf("invalid api call type: %s", apiCall.Type)
	}
	return nil
}
