package test

import (
	"context"
	"github.com/google/uuid"
	"github.com/yangpp6/go-sdk/sdk/security_group_rule"
	"testing"
)

func TestCreateSecurityGroupRule(t *testing.T) {
	_, requestError := DevClient.Apis.SecurityGroupRuleEgressCreateApi.Do(context.Background(), TestCredential, &security_group_rule.SecurityGroupRuleEgressCreateRequest{
		RegionId:        TestNeiMeng8RegionId,
		SecurityGroupId: "sg-vkwmcyz5m5",
		ClientToken:     uuid.NewString(),
		SecurityGroupRules: []security_group_rule.SecurityGroupRuleEgressCreateSecurityGroupRulesRequest{
			{
				Direction:   "egress",
				Action:      "accept",
				Priority:    1,
				Protocol:    "ANY",
				Ethertype:   "IPv4",
				DestCidrIp:  "0.0.0.0/0",
				Description: "开放22端口",
				Range:       "22",
			},
		},
	})
	if requestError != nil {
		t.Error(requestError)
	}
}
