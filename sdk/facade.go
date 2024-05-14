package sdk

import (
	"github.com/yangpp6/go-sdk/sdk/bandwidth"
	"github.com/yangpp6/go-sdk/sdk/common"
	"github.com/yangpp6/go-sdk/sdk/ebs"
	"github.com/yangpp6/go-sdk/sdk/ecs"
	"github.com/yangpp6/go-sdk/sdk/eip"
	"github.com/yangpp6/go-sdk/sdk/idp"
	"github.com/yangpp6/go-sdk/sdk/image"
	"github.com/yangpp6/go-sdk/sdk/job"
	"github.com/yangpp6/go-sdk/sdk/keypair"
	"github.com/yangpp6/go-sdk/sdk/order"
	"github.com/yangpp6/go-sdk/sdk/region"
	"github.com/yangpp6/go-sdk/sdk/security_group"
	"github.com/yangpp6/go-sdk/sdk/security_group_rule"
	"github.com/yangpp6/go-sdk/sdk/subnet"
	"github.com/yangpp6/go-sdk/sdk/user"
	"github.com/yangpp6/go-sdk/sdk/vpc"
)

// Apis api的接口
type Apis struct {
	UserCreateApi                     common.ApiHandler[user.UserCreateRequest, user.UserCreateResponse]
	UserListApi                       common.ApiHandler[user.UserListRequest, user.UserListResponse]
	UserUpdateApi                     common.ApiHandler[user.UserUpdateRequest, user.UserUpdateResponse]
	UserInvalidApi                    common.ApiHandler[user.UserInvalidRequest, user.UserInvalidResponse]
	UserGroupCreateApi                common.ApiHandler[user.UserGroupCreateRequest, user.UserGroupCreateResponse]
	UserAssociationGroupApi           common.ApiHandler[user.UserAssociationGroupRequest, user.UserAssociationGroupResponse]
	QueryUserGroupApi                 common.ApiHandler[user.UserGroupQueryRequest, user.UserGroupQueryResponse]
	RegionListApi                     common.ApiHandler[region.RegionListRequest, region.RegionListResponse]
	IdpCreateApi                      common.ApiHandler[idp.IdpCreateRequest, idp.IdpCreateResponse]
	IdpDeleteApi                      common.ApiHandler[idp.IdpDeleteRequest, idp.IdpDeleteResponse]
	IdpUpdateApi                      common.ApiHandler[idp.IdpUpdateRequest, idp.IdpUpdateResponse]
	IdpListApi                        common.ApiHandler[idp.IdpListRequest, idp.IdpListResponse]
	VpcCreateApi                      common.ApiHandler[vpc.VpcCreateRequest, vpc.VpcCreateResponse]
	VpcUpdateApi                      common.ApiHandler[vpc.VpcUpdateRequest, vpc.VpcUpdateResponse]
	VpcDeleteApi                      common.ApiHandler[vpc.VpcDeleteRequest, vpc.VpcDeleteResponse]
	VpcListApi                        common.ApiHandler[vpc.VpcListRequest, vpc.VpcListResponse]
	SubnetCreateApi                   common.ApiHandler[subnet.SubnetCreateRequest, subnet.SubnetCreateResponse]
	SubnetUpdateApi                   common.ApiHandler[subnet.SubnetUpdateRequest, subnet.SubnetUpdateResponse]
	SubnetDeleteApi                   common.ApiHandler[subnet.SubnetDeleteRequest, subnet.SubnetDeleteResponse]
	SubnetListApi                     common.ApiHandler[subnet.SubnetListRequest, subnet.SubnetListResponse]
	EipCreateApi                      common.ApiHandler[eip.EipCreateRequest, eip.EipCreateResponse]
	EipDeleteApi                      common.ApiHandler[eip.EipDeleteRequest, eip.EipDeleteResponse]
	EipModifySpecApi                  common.ApiHandler[eip.EipModifySpecRequest, eip.EipModifySpecResponse]
	EipChangeNameApi                  common.ApiHandler[eip.EipChangeNameRequest, eip.EipChangeNameResponse]
	EipAssociateApi                   common.ApiHandler[eip.EipAssociateRequest, eip.EipAssociateResponse]
	EipDisassociateApi                common.ApiHandler[eip.EipDisassociateRequest, eip.EipDisassociateResponse]
	EipShowApi                        common.ApiHandler[eip.EipShowRequest, eip.EipShowResponse]
	SecurityGroupCreateApi            common.ApiHandler[security_group.SecurityGroupCreateRequest, security_group.SecurityGroupCreateResponse]
	SecurityGroupModifyAttributionApi common.ApiHandler[security_group.SecurityGroupModifyAttributionRequest, security_group.SecurityGroupModifyAttributionResponse]
	SecurityGroupDeleteApi            common.ApiHandler[security_group.SecurityGroupDeleteRequest, security_group.SecurityGroupDeleteResponse]
	SecurityGroupDescribeAttributeApi common.ApiHandler[security_group.SecurityGroupDescribeAttributeRequest, security_group.SecurityGroupDescribeAttributeResponse]
	SecurityGroupRuleEgressCreateApi  common.ApiHandler[security_group_rule.SecurityGroupRuleEgressCreateRequest, security_group_rule.SecurityGroupRuleEgressCreateResponse]
	SecurityGroupRuleEgressRevokeApi  common.ApiHandler[security_group_rule.SecurityGroupRuleEgressRevokeRequest, security_group_rule.SecurityGroupRuleEgressRevokeResponse]
	SecurityGroupRuleEgressModifyApi  common.ApiHandler[security_group_rule.SecurityGroupRuleEgressModifyRequest, security_group_rule.SecurityGroupRuleEgressModifyResponse]
	SecurityGroupRuleIngressCreateApi common.ApiHandler[security_group_rule.SecurityGroupRuleIngressCreateRequest, security_group_rule.SecurityGroupRuleIngressCreateResponse]
	SecurityGroupRuleIngressRevokeApi common.ApiHandler[security_group_rule.SecurityGroupRuleIngressRevokeRequest, security_group_rule.SecurityGroupRuleIngressRevokeResponse]
	SecurityGroupRuleIngressModifyApi common.ApiHandler[security_group_rule.SecurityGroupRuleIngressModifyRequest, security_group_rule.SecurityGroupRuleIngressModifyResponse]
	SecurityGroupRuleDescribeApi      common.ApiHandler[security_group_rule.SecurityGroupRuleDescribeRequest, security_group_rule.SecurityGroupRuleDescribeResponse]
	OrderPlaceNewPurchaseOrderApi     common.ApiHandler[order.OrderPlaceNewPurchaseOrderRequest, order.OrderPlaceNewPurchaseOrderResponse]
	OrderQueryOrderDetailApi          common.ApiHandler[order.OrderQueryOrderDetailRequest, order.OrderQueryOrderDetailResponse]
	OrderQueryResourceApi             common.ApiHandler[order.OrderQueryResourceRequest, order.OrderQueryResourceResponse]
	OrderPlaceRefundOrderApi          common.ApiHandler[order.OrderPlaceRefundOrderRequest, order.OrderPlaceRefundOrderResponse]
	OrderPlaceUpgradeOrderApi         common.ApiHandler[order.OrderPlaceUpgradeOrderRequest, order.OrderPlaceUpgradeOrderResponse]
	ImageListApi                      common.ApiHandler[image.ImageListRequest, image.ImageListResponse]
	ImageDetailApi                    common.ApiHandler[image.ImageDetailRequest, image.ImageDetailResponse]
	ImageImportApi                    common.ApiHandler[image.ImageImportRequest, image.ImageImportResponse]
	ImageListImportTasksApi           common.ApiHandler[image.ImageListImportTasksRequest, image.ImageListImportTasksResponse]
	ImageDeleteApi                    common.ApiHandler[image.ImageDeleteRequest, image.ImageDeleteResponse]
	ImageUpdateApi                    common.ApiHandler[image.ImageUpdateRequest, image.ImageUpdateResponse]
	EbsCreateApi                      common.ApiHandler[ebs.EbsCreateRequest, ebs.EbsCreateResponse]
	EbsDeleteApi                      common.ApiHandler[ebs.EbsDeleteRequest, ebs.EbsDeleteResponse]
	EbsChangeNameApi                  common.ApiHandler[ebs.EbsChangeNameRequest, ebs.EbsChangeNameResponse]
	EbsChangeSizeApi                  common.ApiHandler[ebs.EbsChangeSizeRequest, ebs.EbsChangeSizeResponse]
	EbsAssociateApi                   common.ApiHandler[ebs.EbsAssociateRequest, ebs.EbsAssociateResponse]
	EbsDisassociateApi                common.ApiHandler[ebs.EbsDisassociateRequest, ebs.EbsDisassociateResponse]
	EbsShowApi                        common.ApiHandler[ebs.EbsShowRequest, ebs.EbsShowResponse]
	JobShowApi                        common.ApiHandler[job.JobShowRequest, job.JobShowResponse]
	BandwidthShowApi                  common.ApiHandler[bandwidth.BandwidthShowRequest, bandwidth.BandwidthShowResponse]
	BandwidthCreateApi                common.ApiHandler[bandwidth.BandwidthCreateRequest, bandwidth.BandwidthCreateResponse]
	BandwidthChangeNameApi            common.ApiHandler[bandwidth.BandwidthChangeNameRequest, bandwidth.BandwidthChangeNameResponse]
	BandwidthChangeSpecApi            common.ApiHandler[bandwidth.BandwidthChangeSpecRequest, bandwidth.BandwidthChangeSpecResponse]
	BandwidthDeleteApi                common.ApiHandler[bandwidth.BandwidthDeleteRequest, bandwidth.BandwidthDeleteResponse]
	BandwidthAssociateEipApi          common.ApiHandler[bandwidth.BandwidthAssociationEipRequest, bandwidth.BandwidthAssociationEipResponse]
	BandwidthDisassociateEipApi       common.ApiHandler[bandwidth.BandwidthDisassociationEipRequest, bandwidth.BandwidthDisassociationEipResponse]
	KeypairAttachApi                  common.ApiHandler[keypair.KeypairAttachRequest, keypair.KeypairAttachResponse]
	KeypairDetachApi                  common.ApiHandler[keypair.KeypairDetachRequest, keypair.KeypairDetachResponse]
	KeypairCreateApi                  common.ApiHandler[keypair.KeypairCreateRequest, keypair.KeypairCreateResponse]
	KeypairDeleteApi                  common.ApiHandler[keypair.KeypairDeleteRequest, keypair.KeypairDeleteResponse]
	KeypairDetailApi                  common.ApiHandler[keypair.KeypairDetailRequest, keypair.KeypairDetailResponse]
	KeypairImportApi                  common.ApiHandler[keypair.KeypairImportRequest, keypair.KeypairImportResponse]
	EcsDescribeInstancesApi           common.ApiHandler[ecs.EcsDescribeInstancesRequest, ecs.EcsDescribeInstancesResponse]
	EcsFlavorListApi                  common.ApiHandler[ecs.EcsFlavorListRequest, ecs.EcsFlavorListResponse]
	EcsCreateInstanceApi              common.ApiHandler[ecs.EcsCreateInstanceRequest, ecs.EcsCreateInstanceResponse]
	EcsJoinSecurityGroupApi           common.ApiHandler[ecs.EcsJoinSecurityGroupRequest, ecs.EcsJoinSecurityGroupResponse]
	EcsLeaveSecurityGroupApi          common.ApiHandler[ecs.EcsLeaveSecurityGroupRequest, ecs.EcsLeaveSecurityGroupResponse]
	EcsVolumeListApi                  common.ApiHandler[ecs.EcsVolumeListRequest, ecs.EcsVolumeListResponse]
	EcsInstanceDetailsApi             common.ApiHandler[ecs.EcsInstanceDetailsRequest, ecs.EcsInstanceDetailsResponse]
	EcsUnsubscribeInstanceApi         common.ApiHandler[ecs.EcsUnsubscribeInstanceRequest, ecs.EcsUnsubscribeInstanceResponse]
	EcsUpdateFlavorSpecApi            common.ApiHandler[ecs.EcsUpdateFlavorSpecRequest, ecs.EcsUpdateFlavorSpecResponse]
	EcsQueryAsyncResultApi            common.ApiHandler[ecs.EcsQueryAsyncResultRequest, ecs.EcsQueryAsyncResultResponse]
	EcsStartInstanceApi               common.ApiHandler[ecs.EcsStartInstanceRequest, ecs.EcsStartInstanceResponse]
	EcsStopInstanceApi                common.ApiHandler[ecs.EcsStopInstanceRequest, ecs.EcsStopInstanceResponse]
	EcsInstanceStatusListApi          common.ApiHandler[ecs.EcsInstanceStatusListRequest, ecs.EcsInstanceStatusListResponse]
	EcsResetPasswordApi               common.ApiHandler[ecs.EcsResetPasswordRequest, ecs.EcsResetPasswordResponse]
	EcsRedeployApi                    common.ApiHandler[ecs.EcsRedeployRequest, ecs.EcsRedeployResponse]
}

// NewApis 注册钩子
func NewApis(client *common.CtyunSender, hooks ...common.ApiHook) *Apis {
	builder := common.NewApiHookBuilder()
	for _, hook := range hooks {
		builder.AddHooks(hook)
	}
	apis := &Apis{
		UserCreateApi:                     common.Wrap[user.UserCreateRequest, user.UserCreateResponse](user.NewUserCreateApi(client), *builder),
		UserListApi:                       common.Wrap[user.UserListRequest, user.UserListResponse](user.NewUserListApi(client), *builder),
		UserUpdateApi:                     common.Wrap[user.UserUpdateRequest, user.UserUpdateResponse](user.NewUserUpdateApi(client), *builder),
		UserInvalidApi:                    common.Wrap[user.UserInvalidRequest, user.UserInvalidResponse](user.NewUserInvalidApi(client), *builder),
		UserGroupCreateApi:                common.Wrap[user.UserGroupCreateRequest, user.UserGroupCreateResponse](user.NewUserGroupCreateApi(client), *builder),
		UserAssociationGroupApi:           common.Wrap[user.UserAssociationGroupRequest, user.UserAssociationGroupResponse](user.NewUserAssociationGroupApi(client), *builder),
		QueryUserGroupApi:                 common.Wrap[user.UserGroupQueryRequest, user.UserGroupQueryResponse](user.NewUserGroupQueryApi(client), *builder),
		RegionListApi:                     common.Wrap[region.RegionListRequest, region.RegionListResponse](region.NewRegionListApi(client), *builder),
		IdpCreateApi:                      common.Wrap[idp.IdpCreateRequest, idp.IdpCreateResponse](idp.NewIdpCreateApi(client), *builder),
		IdpDeleteApi:                      common.Wrap[idp.IdpDeleteRequest, idp.IdpDeleteResponse](idp.NewIdpDeleteApi(client), *builder),
		IdpUpdateApi:                      common.Wrap[idp.IdpUpdateRequest, idp.IdpUpdateResponse](idp.NewIdpUpdateApi(client), *builder),
		IdpListApi:                        common.Wrap[idp.IdpListRequest, idp.IdpListResponse](idp.NewIdpListApi(client), *builder),
		VpcCreateApi:                      common.Wrap[vpc.VpcCreateRequest, vpc.VpcCreateResponse](vpc.NewVpcCreateApi(client), *builder),
		VpcUpdateApi:                      common.Wrap[vpc.VpcUpdateRequest, vpc.VpcUpdateResponse](vpc.NewVpcUpdateApi(client), *builder),
		VpcDeleteApi:                      common.Wrap[vpc.VpcDeleteRequest, vpc.VpcDeleteResponse](vpc.NewVpcDeleteApi(client), *builder),
		VpcListApi:                        common.Wrap[vpc.VpcListRequest, vpc.VpcListResponse](vpc.NewVpcListApi(client), *builder),
		SubnetCreateApi:                   common.Wrap[subnet.SubnetCreateRequest, subnet.SubnetCreateResponse](subnet.NewSubnetCreateApi(client), *builder),
		SubnetUpdateApi:                   common.Wrap[subnet.SubnetUpdateRequest, subnet.SubnetUpdateResponse](subnet.NewSubnetUpdateApi(client), *builder),
		SubnetDeleteApi:                   common.Wrap[subnet.SubnetDeleteRequest, subnet.SubnetDeleteResponse](subnet.NewSubnetDeleteApi(client), *builder),
		SubnetListApi:                     common.Wrap[subnet.SubnetListRequest, subnet.SubnetListResponse](subnet.NewSubnetListApi(client), *builder),
		EipCreateApi:                      common.Wrap[eip.EipCreateRequest, eip.EipCreateResponse](eip.NewEipCreateApi(client), *builder),
		EipDeleteApi:                      common.Wrap[eip.EipDeleteRequest, eip.EipDeleteResponse](eip.NewEipDeleteApi(client), *builder),
		EipModifySpecApi:                  common.Wrap[eip.EipModifySpecRequest, eip.EipModifySpecResponse](eip.NewEipModifySpecApi(client), *builder),
		EipChangeNameApi:                  common.Wrap[eip.EipChangeNameRequest, eip.EipChangeNameResponse](eip.NewEipChangeNameApi(client), *builder),
		EipAssociateApi:                   common.Wrap[eip.EipAssociateRequest, eip.EipAssociateResponse](eip.NewEipAssociateApi(client), *builder),
		EipDisassociateApi:                common.Wrap[eip.EipDisassociateRequest, eip.EipDisassociateResponse](eip.NewEipDisassociateApi(client), *builder),
		EipShowApi:                        common.Wrap[eip.EipShowRequest, eip.EipShowResponse](eip.NewEipShowApi(client), *builder),
		SecurityGroupCreateApi:            common.Wrap[security_group.SecurityGroupCreateRequest, security_group.SecurityGroupCreateResponse](security_group.NewSecurityGroupCreateApi(client), *builder),
		SecurityGroupModifyAttributionApi: common.Wrap[security_group.SecurityGroupModifyAttributionRequest, security_group.SecurityGroupModifyAttributionResponse](security_group.NewSecurityGroupModifyAttributionApi(client), *builder),
		SecurityGroupDeleteApi:            common.Wrap[security_group.SecurityGroupDeleteRequest, security_group.SecurityGroupDeleteResponse](security_group.NewSecurityGroupDeleteApi(client), *builder),
		SecurityGroupDescribeAttributeApi: common.Wrap[security_group.SecurityGroupDescribeAttributeRequest, security_group.SecurityGroupDescribeAttributeResponse](security_group.NewSecurityGroupDescribeAttributeApi(client), *builder),
		SecurityGroupRuleEgressCreateApi:  common.Wrap[security_group_rule.SecurityGroupRuleEgressCreateRequest, security_group_rule.SecurityGroupRuleEgressCreateResponse](security_group_rule.NewSecurityGroupRuleEgressCreateApi(client), *builder),
		SecurityGroupRuleEgressRevokeApi:  common.Wrap[security_group_rule.SecurityGroupRuleEgressRevokeRequest, security_group_rule.SecurityGroupRuleEgressRevokeResponse](security_group_rule.NewSecurityGroupRuleEgressRevokeApi(client), *builder),
		SecurityGroupRuleEgressModifyApi:  common.Wrap[security_group_rule.SecurityGroupRuleEgressModifyRequest, security_group_rule.SecurityGroupRuleEgressModifyResponse](security_group_rule.NewSecurityGroupRuleEgressModifyApi(client), *builder),
		SecurityGroupRuleIngressCreateApi: common.Wrap[security_group_rule.SecurityGroupRuleIngressCreateRequest, security_group_rule.SecurityGroupRuleIngressCreateResponse](security_group_rule.NewSecurityGroupRuleIngressCreateApi(client), *builder),
		SecurityGroupRuleIngressRevokeApi: common.Wrap[security_group_rule.SecurityGroupRuleIngressRevokeRequest, security_group_rule.SecurityGroupRuleIngressRevokeResponse](security_group_rule.NewSecurityGroupRuleIngressRevokeApi(client), *builder),
		SecurityGroupRuleIngressModifyApi: common.Wrap[security_group_rule.SecurityGroupRuleIngressModifyRequest, security_group_rule.SecurityGroupRuleIngressModifyResponse](security_group_rule.NewSecurityGroupRuleIngressModifyApi(client), *builder),
		SecurityGroupRuleDescribeApi:      common.Wrap[security_group_rule.SecurityGroupRuleDescribeRequest, security_group_rule.SecurityGroupRuleDescribeResponse](security_group_rule.NewSecurityGroupRuleDescribeApi(client), *builder),
		OrderPlaceNewPurchaseOrderApi:     common.Wrap[order.OrderPlaceNewPurchaseOrderRequest, order.OrderPlaceNewPurchaseOrderResponse](order.NewOrderOrderApi(client), *builder),
		OrderQueryOrderDetailApi:          common.Wrap[order.OrderQueryOrderDetailRequest, order.OrderQueryOrderDetailResponse](order.NewOrderQueryOrderDetailApi(client), *builder),
		OrderQueryResourceApi:             common.Wrap[order.OrderQueryResourceRequest, order.OrderQueryResourceResponse](order.NewOrderQueryResourceApi(client), *builder),
		OrderPlaceRefundOrderApi:          common.Wrap[order.OrderPlaceRefundOrderRequest, order.OrderPlaceRefundOrderResponse](order.NewOrderPlaceRefundOrderApi(client), *builder),
		OrderPlaceUpgradeOrderApi:         common.Wrap[order.OrderPlaceUpgradeOrderRequest, order.OrderPlaceUpgradeOrderResponse](order.NewOrderPlaceUpgradeOrderApi(client), *builder),
		ImageListApi:                      common.Wrap[image.ImageListRequest, image.ImageListResponse](image.NewImageListApi(client), *builder),
		ImageDetailApi:                    common.Wrap[image.ImageDetailRequest, image.ImageDetailResponse](image.NewImageDetailApi(client), *builder),
		ImageImportApi:                    common.Wrap[image.ImageImportRequest, image.ImageImportResponse](image.NewImageImportApi(client), *builder),
		ImageListImportTasksApi:           common.Wrap[image.ImageListImportTasksRequest, image.ImageListImportTasksResponse](image.NewImageListImportTasksApi(client), *builder),
		ImageDeleteApi:                    common.Wrap[image.ImageDeleteRequest, image.ImageDeleteResponse](image.NewImageDeleteApi(client), *builder),
		ImageUpdateApi:                    common.Wrap[image.ImageUpdateRequest, image.ImageUpdateResponse](image.NewImageUpdateApi(client), *builder),
		EbsCreateApi:                      common.Wrap[ebs.EbsCreateRequest, ebs.EbsCreateResponse](ebs.NewEbsCreateApi(client), *builder),
		EbsDeleteApi:                      common.Wrap[ebs.EbsDeleteRequest, ebs.EbsDeleteResponse](ebs.NewEbsDeleteApi(client), *builder),
		EbsChangeNameApi:                  common.Wrap[ebs.EbsChangeNameRequest, ebs.EbsChangeNameResponse](ebs.NewEbsChangeNameApi(client), *builder),
		EbsChangeSizeApi:                  common.Wrap[ebs.EbsChangeSizeRequest, ebs.EbsChangeSizeResponse](ebs.NewEbsChangeSizeApi(client), *builder),
		EbsAssociateApi:                   common.Wrap[ebs.EbsAssociateRequest, ebs.EbsAssociateResponse](ebs.NewEbsAssociateApi(client), *builder),
		EbsDisassociateApi:                common.Wrap[ebs.EbsDisassociateRequest, ebs.EbsDisassociateResponse](ebs.NewEbsDisassociateApi(client), *builder),
		EbsShowApi:                        common.Wrap[ebs.EbsShowRequest, ebs.EbsShowResponse](ebs.NewEbsShowApi(client), *builder),
		JobShowApi:                        common.Wrap[job.JobShowRequest, job.JobShowResponse](job.NewJobShowApi(client), *builder),
		BandwidthShowApi:                  common.Wrap[bandwidth.BandwidthShowRequest, bandwidth.BandwidthShowResponse](bandwidth.NewBandwidthShowApi(client), *builder),
		BandwidthCreateApi:                common.Wrap[bandwidth.BandwidthCreateRequest, bandwidth.BandwidthCreateResponse](bandwidth.NewBandwidthCreateApi(client), *builder),
		BandwidthChangeNameApi:            common.Wrap[bandwidth.BandwidthChangeNameRequest, bandwidth.BandwidthChangeNameResponse](bandwidth.NewBandwidthChangeNameApi(client), *builder),
		BandwidthChangeSpecApi:            common.Wrap[bandwidth.BandwidthChangeSpecRequest, bandwidth.BandwidthChangeSpecResponse](bandwidth.NewBandwidthChangeSpecApi(client), *builder),
		BandwidthDeleteApi:                common.Wrap[bandwidth.BandwidthDeleteRequest, bandwidth.BandwidthDeleteResponse](bandwidth.NewBandwidthDeleteApi(client), *builder),
		BandwidthAssociateEipApi:          common.Wrap[bandwidth.BandwidthAssociationEipRequest, bandwidth.BandwidthAssociationEipResponse](bandwidth.NewBandwidthAssociationEipApi(client), *builder),
		BandwidthDisassociateEipApi:       common.Wrap[bandwidth.BandwidthDisassociationEipRequest, bandwidth.BandwidthDisassociationEipResponse](bandwidth.NewBandwidthDisassociationEipApi(client), *builder),
		KeypairAttachApi:                  common.Wrap[keypair.KeypairAttachRequest, keypair.KeypairAttachResponse](keypair.NewKeypairAttachApi(client), *builder),
		KeypairDetachApi:                  common.Wrap[keypair.KeypairDetachRequest, keypair.KeypairDetachResponse](keypair.NewKeypairDetachApi(client), *builder),
		KeypairCreateApi:                  common.Wrap[keypair.KeypairCreateRequest, keypair.KeypairCreateResponse](keypair.NewKeypairCreateApi(client), *builder),
		KeypairDeleteApi:                  common.Wrap[keypair.KeypairDeleteRequest, keypair.KeypairDeleteResponse](keypair.NewKeypairDeleteApi(client), *builder),
		KeypairDetailApi:                  common.Wrap[keypair.KeypairDetailRequest, keypair.KeypairDetailResponse](keypair.NewKeypairDetailApi(client), *builder),
		KeypairImportApi:                  common.Wrap[keypair.KeypairImportRequest, keypair.KeypairImportResponse](keypair.NewKeypairImportApi(client), *builder),
		EcsDescribeInstancesApi:           common.Wrap[ecs.EcsDescribeInstancesRequest, ecs.EcsDescribeInstancesResponse](ecs.NewEcsDescribeInstancesApi(client), *builder),
		EcsFlavorListApi:                  common.Wrap[ecs.EcsFlavorListRequest, ecs.EcsFlavorListResponse](ecs.NewEcsFlavorListApi(client), *builder),
		EcsCreateInstanceApi:              common.Wrap[ecs.EcsCreateInstanceRequest, ecs.EcsCreateInstanceResponse](ecs.NewEcsCreateInstanceApi(client), *builder),
		EcsJoinSecurityGroupApi:           common.Wrap[ecs.EcsJoinSecurityGroupRequest, ecs.EcsJoinSecurityGroupResponse](ecs.NewEcsJoinSecurityGroupApi(client), *builder),
		EcsLeaveSecurityGroupApi:          common.Wrap[ecs.EcsLeaveSecurityGroupRequest, ecs.EcsLeaveSecurityGroupResponse](ecs.NewEcsLeaveSecurityGroupApi(client), *builder),
		EcsVolumeListApi:                  common.Wrap[ecs.EcsVolumeListRequest, ecs.EcsVolumeListResponse](ecs.NewEcsVolumeListApi(client), *builder),
		EcsInstanceDetailsApi:             common.Wrap[ecs.EcsInstanceDetailsRequest, ecs.EcsInstanceDetailsResponse](ecs.NewEcsInstanceDetailsApi(client), *builder),
		EcsUnsubscribeInstanceApi:         common.Wrap[ecs.EcsUnsubscribeInstanceRequest, ecs.EcsUnsubscribeInstanceResponse](ecs.NewEcsUnsubscribeInstanceApi(client), *builder),
		EcsUpdateFlavorSpecApi:            common.Wrap[ecs.EcsUpdateFlavorSpecRequest, ecs.EcsUpdateFlavorSpecResponse](ecs.NewEcsUpdateFlavorSpecApi(client), *builder),
		EcsQueryAsyncResultApi:            common.Wrap[ecs.EcsQueryAsyncResultRequest, ecs.EcsQueryAsyncResultResponse](ecs.NewEcsQueryAsyncResultApi(client), *builder),
		EcsStartInstanceApi:               common.Wrap[ecs.EcsStartInstanceRequest, ecs.EcsStartInstanceResponse](ecs.NewEcsStartInstanceApi(client), *builder),
		EcsStopInstanceApi:                common.Wrap[ecs.EcsStopInstanceRequest, ecs.EcsStopInstanceResponse](ecs.NewEcsStopInstanceApi(client), *builder),
		EcsInstanceStatusListApi:          common.Wrap[ecs.EcsInstanceStatusListRequest, ecs.EcsInstanceStatusListResponse](ecs.NewEcsInstanceStatusListApi(client), *builder),
		EcsResetPasswordApi:               common.Wrap[ecs.EcsResetPasswordRequest, ecs.EcsResetPasswordResponse](ecs.NewEcsResetPasswordApi(client), *builder),
		EcsRedeployApi:                    common.Wrap[ecs.EcsRedeployRequest, ecs.EcsRedeployResponse](ecs.NewEcsRedeployApi(client), *builder),
	}
	return apis
}
