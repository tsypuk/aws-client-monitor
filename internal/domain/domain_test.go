package domain

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func loadJsonFromFile(filename string) []byte {
	file, _ := os.Open(filename)
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
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
				Version:             1,
				ClientId:            "test-client",
				Type:                "ApiCall",
				Service:             "S3",
				Api:                 "ListBuckets",
				Timestamp:           1728846180515,
				AttemptCount:        1,
				Region:              "eu-west-1",
				UserAgent:           "aws-cli/1.27.92 md/Botocore#1.31.2 ua/2.0 os/macos#21.6.0 md/arch#x86_64 lang/python#3.10.14 md/pyimpl#CPython cfg/retry-mode#legacy botocore/1.31.2",
				FinalHttpStatusCode: 400,
				//FinalAwsException:        "ExpiredToken",
				//FinalAwsExceptionMessage: "The provided token has expired.",
				Latency:            1167,
				MaxRetriesExceeded: 0,
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
