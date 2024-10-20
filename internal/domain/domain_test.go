package domain

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func loadJsonFromFile(filename string) []byte {
	file, _ := os.Open(filename)
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	return byteValue
}

func TestNewApiCall(t *testing.T) {
	type args struct {
		payload UdpPayload
	}
	tests := []struct {
		name    string
		args    args
		want    *ApiCall
		wantErr bool
	}{
		{
			name: "s3 ApiCall Error",
			// load from json file
			args: args{payload: UdpPayload{Payload: loadJsonFromFile("testdata/s3.apicall.error.json")}},
			want: &ApiCall{
				Version:                  1,
				ClientId:                 "test-client",
				ApiBaseType:              ApiBaseType{Type: "ApiCall"},
				Service:                  "S3",
				Api:                      "ListBuckets",
				Timestamp:                1728846180515,
				AttemptCount:             1,
				Region:                   "eu-west-1",
				UserAgent:                "aws-cli/1.27.92 md/Botocore#1.31.2 ua/2.0 os/macos#21.6.0 md/arch#x86_64 lang/python#3.10.14 md/pyimpl#CPython cfg/retry-mode#legacy botocore/1.31.2",
				FinalHttpStatusCode:      400,
				FinalAwsException:        "ExpiredToken",
				FinalAwsExceptionMessage: "The provided token has expired.",
				Latency:                  1167,
				MaxRetriesExceeded:       0,
			},
			wantErr: false,
		},
		{
			name: "sqs ApiCall OK",
			// load from json file
			args: args{payload: UdpPayload{Payload: loadJsonFromFile("testdata/sqs.apicall.ok.json")}},
			want: &ApiCall{
				Version:             1,
				ClientId:            "test-client",
				ApiBaseType:         ApiBaseType{Type: "ApiCall"},
				Service:             "SQS",
				Api:                 "ListQueues",
				Timestamp:           1728846984653,
				AttemptCount:        1,
				Region:              "eu-west-1",
				UserAgent:           "aws-cli/1.27.92 md/Botocore#1.31.2 ua/2.0 os/macos#21.6.0 md/arch#x86_64 lang/python#3.10.14 md/pyimpl#CPython cfg/retry-mode#legacy botocore/1.31.2",
				FinalHttpStatusCode: 200,
				Latency:             2151,
				MaxRetriesExceeded:  0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run the actual test
			got, err := NewApiCall(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewApiCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApiCall() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewApiCallAttempt(t *testing.T) {
	type args struct {
		payload UdpPayload
	}
	tests := []struct {
		name    string
		args    args
		want    *ApiCallAttempt
		wantErr bool
	}{
		{
			name: "sqs ApiCallAttempt OK",
			// load from json file
			args: args{payload: UdpPayload{Payload: loadJsonFromFile("testdata/sqs.apicallattempt.ok.json")}},
			want: &ApiCallAttempt{
				Version:        1,
				ClientId:       "test-client",
				ApiBaseType:    ApiBaseType{Type: "ApiCallAttempt"},
				Service:        "SQS",
				Api:            "ListQueues",
				Timestamp:      1728849034025,
				AttemptLatency: 313,
				Fqdn:           "eu-west-1.queue.amazonaws.com",
				UserAgent:      "aws-cli/1.27.92 md/Botocore#1.31.2 ua/2.0 os/macos#21.6.0 md/arch#x86_64 lang/python#3.10.14 md/pyimpl#CPython cfg/retry-mode#legacy botocore/1.31.2",
				AccessKey:      "ASIATEST",
				Region:         "eu-west-1",
				SessionToken:   "IQTest=",
				HttpStatusCode: 200,
				XAmznRequestId: "b18f8c1c-aaea-5d12-a816-d39a8c209dd9",
			},
			wantErr: false,
		},
		{
			name: "sqs ApiCallAttempt Error",
			// load from json file
			args: args{payload: UdpPayload{Payload: loadJsonFromFile("testdata/sqs.apicallattempt.error.json")}},
			want: &ApiCallAttempt{
				Version:             1,
				ClientId:            "test-client",
				ApiBaseType:         ApiBaseType{Type: "ApiCallAttempt"},
				Service:             "SQS",
				Api:                 "ListQueues",
				Timestamp:           1728848411600,
				AttemptLatency:      200,
				Fqdn:                "eu-west-1.queue.amazonaws.com",
				UserAgent:           "aws-cli/1.27.92 md/Botocore#1.31.2 ua/2.0 os/macos#21.6.0 md/arch#x86_64 lang/python#3.10.14 md/pyimpl#CPython cfg/retry-mode#legacy botocore/1.31.2",
				AccessKey:           "ASIATEST",
				Region:              "eu-west-1",
				SessionToken:        "IQTest=",
				HttpStatusCode:      403,
				XAmznRequestId:      "93630470-f629-58a4-a2e8-5bf6d27281e2",
				AwsException:        "ExpiredToken",
				AwsExceptionMessage: "The security token included in the request is expired",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run the actual test
			got, err := NewApiCallAttempt(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewApiCallAttempt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApiCallAttempt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
