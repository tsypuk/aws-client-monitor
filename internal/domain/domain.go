package domain

type ApiCallAttempt struct {
	Version        int    `json:"Version"`
	ClientId       string `json:"ClientId"`
	Type           string `json:"Type"`
	Service        string `json:"Service"`
	Api            string `json:"Api"`
	Timestamp      int64  `json:"Timestamp"`
	AttemptLatency int    `json:"AttemptLatency"`
	Fqdn           string `json:"Fqdn"`
	UserAgent      string `json:"UserAgent"`
	AccessKey      string `json:"AccessKey"`
	Region         string `json:"Region"`
	SessionToken   string `json:"SessionToken"`
	HttpStatusCode int    `json:"HttpStatusCode"`
	XAmzRequestId  string `json:"XAmzRequestId"`
	XAmzId2        string `json:"XAmzId2"`
}

// Struct for the ApiCall message
type ApiCall struct {
	Version             int    `json:"Version"`
	ClientId            string `json:"ClientId"`
	Type                string `json:"Type"`
	Service             string `json:"Service"`
	Api                 string `json:"Api"`
	Timestamp           int64  `json:"Timestamp"`
	AttemptCount        int    `json:"AttemptCount"`
	Region              string `json:"Region"`
	UserAgent           string `json:"UserAgent"`
	FinalHttpStatusCode int    `json:"FinalHttpStatusCode"`
	Latency             int    `json:"Latency"`
	MaxRetriesExceeded  int    `json:"MaxRetriesExceeded"`
}
