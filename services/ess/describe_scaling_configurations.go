package ess

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DescribeScalingConfigurations(request *DescribeScalingConfigurationsRequest) (response *DescribeScalingConfigurationsResponse, err error) {
	response = CreateDescribeScalingConfigurationsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeScalingConfigurationsWithChan(request *DescribeScalingConfigurationsRequest) (<-chan *DescribeScalingConfigurationsResponse, <-chan error) {
	responseChan := make(chan *DescribeScalingConfigurationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScalingConfigurations(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeScalingConfigurationsWithCallback(request *DescribeScalingConfigurationsRequest, callback func(response *DescribeScalingConfigurationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScalingConfigurationsResponse
		var err error
		defer close(result)
		response, err = client.DescribeScalingConfigurations(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DescribeScalingConfigurationsRequest struct {
	*requests.RpcRequest
	PageSize                   requests.Integer `position:"Query" name:"PageSize"`
	ScalingConfigurationId2    string           `position:"Query" name:"ScalingConfigurationId.2"`
	ScalingConfigurationId3    string           `position:"Query" name:"ScalingConfigurationId.3"`
	ScalingConfigurationId1    string           `position:"Query" name:"ScalingConfigurationId.1"`
	ScalingConfigurationId9    string           `position:"Query" name:"ScalingConfigurationId.9"`
	ScalingConfigurationId8    string           `position:"Query" name:"ScalingConfigurationId.8"`
	ResourceOwnerAccount       string           `position:"Query" name:"ResourceOwnerAccount"`
	ScalingConfigurationId7    string           `position:"Query" name:"ScalingConfigurationId.7"`
	ScalingConfigurationName3  string           `position:"Query" name:"ScalingConfigurationName.3"`
	ScalingConfigurationId6    string           `position:"Query" name:"ScalingConfigurationId.6"`
	ScalingConfigurationName2  string           `position:"Query" name:"ScalingConfigurationName.2"`
	ScalingConfigurationId5    string           `position:"Query" name:"ScalingConfigurationId.5"`
	ScalingConfigurationName1  string           `position:"Query" name:"ScalingConfigurationName.1"`
	ScalingConfigurationId4    string           `position:"Query" name:"ScalingConfigurationId.4"`
	ResourceOwnerId            requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount               string           `position:"Query" name:"OwnerAccount"`
	ScalingConfigurationName10 string           `position:"Query" name:"ScalingConfigurationName.10"`
	PageNumber                 requests.Integer `position:"Query" name:"PageNumber"`
	ScalingConfigurationName4  string           `position:"Query" name:"ScalingConfigurationName.4"`
	ScalingConfigurationName5  string           `position:"Query" name:"ScalingConfigurationName.5"`
	ScalingConfigurationName6  string           `position:"Query" name:"ScalingConfigurationName.6"`
	ScalingConfigurationName7  string           `position:"Query" name:"ScalingConfigurationName.7"`
	ScalingConfigurationName8  string           `position:"Query" name:"ScalingConfigurationName.8"`
	OwnerId                    requests.Integer `position:"Query" name:"OwnerId"`
	ScalingConfigurationName9  string           `position:"Query" name:"ScalingConfigurationName.9"`
	ScalingGroupId             string           `position:"Query" name:"ScalingGroupId"`
	ScalingConfigurationId10   string           `position:"Query" name:"ScalingConfigurationId.10"`
}

type DescribeScalingConfigurationsResponse struct {
	*responses.BaseResponse
	TotalCount            requests.Integer `json:"TotalCount" xml:"TotalCount"`
	PageNumber            requests.Integer `json:"PageNumber" xml:"PageNumber"`
	PageSize              requests.Integer `json:"PageSize" xml:"PageSize"`
	RequestId             string           `json:"RequestId" xml:"RequestId"`
	ScalingConfigurations struct {
		ScalingConfiguration []struct {
			ScalingConfigurationId      string           `json:"ScalingConfigurationId" xml:"ScalingConfigurationId"`
			ScalingConfigurationName    string           `json:"ScalingConfigurationName" xml:"ScalingConfigurationName"`
			ScalingGroupId              string           `json:"ScalingGroupId" xml:"ScalingGroupId"`
			InstanceName                string           `json:"InstanceName" xml:"InstanceName"`
			ImageId                     string           `json:"ImageId" xml:"ImageId"`
			InstanceType                string           `json:"InstanceType" xml:"InstanceType"`
			InstanceGeneration          string           `json:"InstanceGeneration" xml:"InstanceGeneration"`
			SecurityGroupId             string           `json:"SecurityGroupId" xml:"SecurityGroupId"`
			IoOptimized                 string           `json:"IoOptimized" xml:"IoOptimized"`
			InternetChargeType          string           `json:"InternetChargeType" xml:"InternetChargeType"`
			InternetMaxBandwidthIn      requests.Integer `json:"InternetMaxBandwidthIn" xml:"InternetMaxBandwidthIn"`
			InternetMaxBandwidthOut     requests.Integer `json:"InternetMaxBandwidthOut" xml:"InternetMaxBandwidthOut"`
			SystemDiskCategory          string           `json:"SystemDiskCategory" xml:"SystemDiskCategory"`
			SystemDiskSize              requests.Integer `json:"SystemDiskSize" xml:"SystemDiskSize"`
			LifecycleState              string           `json:"LifecycleState" xml:"LifecycleState"`
			CreationTime                string           `json:"CreationTime" xml:"CreationTime"`
			LoadBalancerWeight          requests.Integer `json:"LoadBalancerWeight" xml:"LoadBalancerWeight"`
			UserData                    string           `json:"UserData" xml:"UserData"`
			KeyPairName                 string           `json:"KeyPairName" xml:"KeyPairName"`
			RamRoleName                 string           `json:"RamRoleName" xml:"RamRoleName"`
			DeploymentSetId             string           `json:"DeploymentSetId" xml:"DeploymentSetId"`
			SecurityEnhancementStrategy string           `json:"SecurityEnhancementStrategy" xml:"SecurityEnhancementStrategy"`
			InstanceTypes               struct {
				InstanceType []string `json:"InstanceType" xml:"InstanceType"`
			} `json:"InstanceTypes" xml:"InstanceTypes"`
			DataDisks struct {
				DataDisk []struct {
					Size       requests.Integer `json:"Size" xml:"Size"`
					Category   string           `json:"Category" xml:"Category"`
					SnapshotId string           `json:"SnapshotId" xml:"SnapshotId"`
					Device     string           `json:"Device" xml:"Device"`
				} `json:"DataDisk" xml:"DataDisk"`
			} `json:"DataDisks" xml:"DataDisks"`
			Tags struct {
				Tag []struct {
					Key   string `json:"Key" xml:"Key"`
					Value string `json:"Value" xml:"Value"`
				} `json:"Tag" xml:"Tag"`
			} `json:"Tags" xml:"Tags"`
		} `json:"ScalingConfiguration" xml:"ScalingConfiguration"`
	} `json:"ScalingConfigurations" xml:"ScalingConfigurations"`
}

func CreateDescribeScalingConfigurationsRequest() (request *DescribeScalingConfigurationsRequest) {
	request = &DescribeScalingConfigurationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DescribeScalingConfigurations", "", "")
	return
}

func CreateDescribeScalingConfigurationsResponse() (response *DescribeScalingConfigurationsResponse) {
	response = &DescribeScalingConfigurationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
