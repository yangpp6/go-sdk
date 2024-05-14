package job

import (
	"context"
	"net/http"
	"openapi-sdk-go/sdk/common"
)

type jobShowApi struct {
	common.CtyunRequestBuilder
	client *common.CtyunSender
}

func NewJobShowApi(client *common.CtyunSender) common.ApiHandler[JobShowRequest, JobShowResponse] {
	return &jobShowApi{
		CtyunRequestBuilder: common.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/job/info",
		},
		client: client,
	}
}

type JobShowRequest struct {
	RegionID string
	JobID    string
}

type JobShowResponse struct {
	Status     int               `json:"status"`
	ResourceId string            `json:"resourceId"`
	Fields     map[string]string `json:"fields"`
	JobStatus  string            `json:"jobStatus"`
	JobID      string            `json:"jobID"`
}

func (s *jobShowApi) Do(ctx context.Context, credential *common.Credential, t *JobShowRequest) (*JobShowResponse, common.CtyunRequestError) {
	builder := s.WithCredential(credential)
	builder.AddParam("jobID", t.JobID)
	builder.AddParam("regionID", t.RegionID)
	resp, requestError := s.client.SendCtEcs(ctx, builder)
	if requestError != nil {
		return nil, requestError
	}
	response := &JobShowResponse{}
	err := resp.ParseWithCheck(response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
