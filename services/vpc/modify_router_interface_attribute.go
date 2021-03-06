package vpc

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

func (client *Client) ModifyRouterInterfaceAttribute(request *ModifyRouterInterfaceAttributeRequest) (response *ModifyRouterInterfaceAttributeResponse, err error) {
	response = CreateModifyRouterInterfaceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyRouterInterfaceAttributeWithChan(request *ModifyRouterInterfaceAttributeRequest) (<-chan *ModifyRouterInterfaceAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyRouterInterfaceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyRouterInterfaceAttribute(request)
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

func (client *Client) ModifyRouterInterfaceAttributeWithCallback(request *ModifyRouterInterfaceAttributeRequest, callback func(response *ModifyRouterInterfaceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyRouterInterfaceAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyRouterInterfaceAttribute(request)
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

type ModifyRouterInterfaceAttributeRequest struct {
	*requests.RpcRequest
	HealthCheckSourceIp      string           `position:"Query" name:"HealthCheckSourceIp"`
	OppositeRouterType       string           `position:"Query" name:"OppositeRouterType"`
	HealthCheckTargetIp      string           `position:"Query" name:"HealthCheckTargetIp"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	OppositeInterfaceOwnerId requests.Integer `position:"Query" name:"OppositeInterfaceOwnerId"`
	OppositeRouterId         string           `position:"Query" name:"OppositeRouterId"`
	Description              string           `position:"Query" name:"Description"`
	Name                     string           `position:"Query" name:"Name"`
	OppositeInterfaceId      string           `position:"Query" name:"OppositeInterfaceId"`
	RouterInterfaceId        string           `position:"Query" name:"RouterInterfaceId"`
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	DeleteHealthCheckIp      requests.Boolean `position:"Query" name:"DeleteHealthCheckIp"`
}

type ModifyRouterInterfaceAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyRouterInterfaceAttributeRequest() (request *ModifyRouterInterfaceAttributeRequest) {
	request = &ModifyRouterInterfaceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyRouterInterfaceAttribute", "", "")
	return
}

func CreateModifyRouterInterfaceAttributeResponse() (response *ModifyRouterInterfaceAttributeResponse) {
	response = &ModifyRouterInterfaceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
