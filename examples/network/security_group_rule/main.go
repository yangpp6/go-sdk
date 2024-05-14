package main

import (
	"context"
	"flag"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/security_group_rule"
	"log"
	"net/http"
)

func securityGroupRuleOperation(ak, sk, regionID, securityGroupID string) {
	client := &common.CtyunSender{
		Client:        &http.Client{},
		CtVpcEndPoint: "ctvpc-global.ctapi.ctyun.cn",
	}

	credential, _ := common.NewCredential(ak, sk)
	ctx := context.TODO()
	ingressCreateHandler := security_group_rule.NewSecurityGroupRuleIngressCreateApi(client)

	createRes, err := ingressCreateHandler.Do(ctx, credential, &security_group_rule.SecurityGroupRuleIngressCreateRequest{
		RegionId:        regionID,
		ClientToken:     "xyz",
		SecurityGroupId: securityGroupID,
		SecurityGroupRules: []security_group_rule.SecurityGroupRuleIngressCreateSecurityGroupRulesRequest{{
			Direction:   "ingress",
			Action:      "accept",
			Priority:    1,
			Protocol:    "TCP",
			Ethertype:   "IPv4",
			DestCidrIp:  "0.0.0.0/0",
			Description: "dafgsdfd",
			Range:       "1-200",
		}},
	})

	if err != nil {
		panic(err)
	}

	ingressSgRuleID := createRes.SgRuleIds[0]

	egressCreateHandler := security_group_rule.NewSecurityGroupRuleEgressCreateApi(client)

	egressCreateRes, err := egressCreateHandler.Do(ctx, credential, &security_group_rule.SecurityGroupRuleEgressCreateRequest{
		RegionId:        regionID,
		ClientToken:     "xyz",
		SecurityGroupId: securityGroupID,
		SecurityGroupRules: []security_group_rule.SecurityGroupRuleEgressCreateSecurityGroupRulesRequest{{
			Direction:   "gress",
			Action:      "accept",
			Priority:    1,
			Protocol:    "TCP",
			Ethertype:   "IPv4",
			DestCidrIp:  "0.0.0.0/0",
			Description: "dafgsdfd",
			Range:       "1-200",
		}},
	})

	if err != nil {
		panic(err)
	}

	egressSgRuleID := egressCreateRes.SgRuleIds[0]

	descHandler := security_group_rule.NewSecurityGroupRuleDescribeApi(client)
	res, err := descHandler.Do(ctx, credential, &security_group_rule.SecurityGroupRuleDescribeRequest{
		RegionId:            regionID,
		SecurityGroupId:     securityGroupID,
		SecurityGroupRuleId: ingressSgRuleID,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("%+v\n", res)

	res, err = descHandler.Do(ctx, credential, &security_group_rule.SecurityGroupRuleDescribeRequest{
		RegionId:            regionID,
		SecurityGroupId:     securityGroupID,
		SecurityGroupRuleId: egressSgRuleID,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("%+v\n", res)

	ingressModifyHandler := security_group_rule.NewSecurityGroupRuleIngressModifyApi(client)
	_, err = ingressModifyHandler.Do(ctx, credential, &security_group_rule.SecurityGroupRuleIngressModifyRequest{
		RegionId:            regionID,
		SecurityGroupId:     securityGroupID,
		SecurityGroupRuleId: ingressSgRuleID,
		ClientToken:         "xyz",
		Description:         "xxxxx",
	})
	if err != nil {
		panic(err)
	}

	egressModifyHandler := security_group_rule.NewSecurityGroupRuleEgressModifyApi(client)
	_, err = egressModifyHandler.Do(ctx, credential, &security_group_rule.SecurityGroupRuleEgressModifyRequest{
		RegionId:            regionID,
		SecurityGroupId:     securityGroupID,
		SecurityGroupRuleId: egressSgRuleID,
		ClientToken:         "xyz",
		Description:         "xxxxx",
	})
	if err != nil {
		panic(err)
	}

	ingressRevokeHandler := security_group_rule.NewSecurityGroupRuleIngressRevokeApi(client)
	_, err = ingressRevokeHandler.Do(ctx, credential, &security_group_rule.SecurityGroupRuleIngressRevokeRequest{
		RegionId:            regionID,
		SecurityGroupId:     securityGroupID,
		SecurityGroupRuleId: ingressSgRuleID,
		ClientToken:         "xyz",
	})
	if err != nil {
		panic(err)
	}

	egressRevokeHandler := security_group_rule.NewSecurityGroupRuleEgressRevokeApi(client)
	_, err = egressRevokeHandler.Do(ctx, credential, &security_group_rule.SecurityGroupRuleEgressRevokeRequest{
		RegionId:            regionID,
		SecurityGroupId:     securityGroupID,
		SecurityGroupRuleId: ingressSgRuleID,
		ClientToken:         "xyz",
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	var action string
	var ak string
	var sk string
	var regionID string
	var securityGroupID string
	flag.StringVar(&action, "action", "list", "example action: list / operation")
	flag.StringVar(&ak, "ak", "", "access key")
	flag.StringVar(&sk, "sk", "", "secret key")
	flag.StringVar(&regionID, "region-id", "", "region id")
	flag.StringVar(&securityGroupID, "security-group-id", "", "security group id")
	flag.Parse()

	if len(ak) == 0 || len(sk) == 0 || len(regionID) == 0 || len(securityGroupID) == 0 {
		log.Print("ak or sk or region-id or security-group-id is required")
		return
	}

	securityGroupRuleOperation(ak, sk, regionID, securityGroupID)
}
