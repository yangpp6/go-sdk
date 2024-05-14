package test

import (
	"context"
	"github.com/yangpp6/go-sdk/sdk"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/vpc"
	"testing"
)

func TestVpcCreateApi_Do(t *testing.T) {
	credential, _ := common.NewCredential("cb65661223244d2fb9f99baac12478a3", "181fcdb898ae4ac9ac7af3563874d091")
	type fields struct {
		client *sdk.CtyunClient
	}
	type args struct {
		ctx        context.Context
		credential *common.Credential
		req        *vpc.VpcCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *vpc.VpcCreateResponse
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				client: sdk.EnvOf(sdk.EnvironmentTest, sdk.CtyunClientConfig{ApiHooks: []common.ApiHook{common.ConsoleLogApiHook}}),
			},
			args: args{
				ctx:        context.Background(),
				credential: credential,
				req: &vpc.VpcCreateRequest{
					RegionId:    "81f7728662dd11ec810800155d307d5b",
					Name:        "vpcTest",
					Cidr:        "10.0.0.0/8",
					Description: "我是测试",
					EnableIpv6:  false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.fields.client.Apis.VpcCreateApi.Do(tt.args.ctx, tt.args.credential, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Do() got = %v, want %v", got, tt.want)
			// }
		})
	}
}
