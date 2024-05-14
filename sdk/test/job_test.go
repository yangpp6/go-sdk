package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/job"
	"testing"
)

func TestJobShowApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("", "")
	do, _ := BuildProdClient().Apis.JobShowApi.Do(context.Background(), credential, &job.JobShowRequest{
		RegionID: ProdHuaDong1RegionId,
		JobID:    "b85dee4c-eb7a-4345-aab3-6920b5ba04d2",
	})
	marshal, _ := json.Marshal(do)
	fmt.Printf(string(marshal))
}
