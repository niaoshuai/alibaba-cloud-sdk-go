package ecs

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

func (client *Client) DescribeTaskAttribute(request *DescribeTaskAttributeRequest) (response *DescribeTaskAttributeResponse, err error) {
	response = CreateDescribeTaskAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeTaskAttributeWithChan(request *DescribeTaskAttributeRequest) (<-chan *DescribeTaskAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeTaskAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTaskAttribute(request)
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

func (client *Client) DescribeTaskAttributeWithCallback(request *DescribeTaskAttributeRequest, callback func(response *DescribeTaskAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTaskAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeTaskAttribute(request)
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

type DescribeTaskAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	TaskId               string           `position:"Query" name:"TaskId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeTaskAttributeResponse struct {
	*responses.BaseResponse
	RequestId            string           `json:"RequestId" xml:"RequestId"`
	TaskId               string           `json:"TaskId" xml:"TaskId"`
	RegionId             string           `json:"RegionId" xml:"RegionId"`
	TaskAction           string           `json:"TaskAction" xml:"TaskAction"`
	TaskStatus           string           `json:"TaskStatus" xml:"TaskStatus"`
	TaskProcess          string           `json:"TaskProcess" xml:"TaskProcess"`
	SupportCancel        string           `json:"SupportCancel" xml:"SupportCancel"`
	TotalCount           requests.Integer `json:"TotalCount" xml:"TotalCount"`
	SuccessCount         requests.Integer `json:"SuccessCount" xml:"SuccessCount"`
	FailedCount          requests.Integer `json:"FailedCount" xml:"FailedCount"`
	CreationTime         string           `json:"CreationTime" xml:"CreationTime"`
	FinishedTime         string           `json:"FinishedTime" xml:"FinishedTime"`
	OperationProgressSet struct {
		OperationProgress []struct {
			OperationStatus string `json:"OperationStatus" xml:"OperationStatus"`
			ErrorCode       string `json:"ErrorCode" xml:"ErrorCode"`
			ErrorMsg        string `json:"ErrorMsg" xml:"ErrorMsg"`
			RelatedItemSet  struct {
				RelatedItem []struct {
					Name  string `json:"Name" xml:"Name"`
					Value string `json:"Value" xml:"Value"`
				} `json:"RelatedItem" xml:"RelatedItem"`
			} `json:"RelatedItemSet" xml:"RelatedItemSet"`
		} `json:"OperationProgress" xml:"OperationProgress"`
	} `json:"OperationProgressSet" xml:"OperationProgressSet"`
}

func CreateDescribeTaskAttributeRequest() (request *DescribeTaskAttributeRequest) {
	request = &DescribeTaskAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeTaskAttribute", "", "")
	return
}

func CreateDescribeTaskAttributeResponse() (response *DescribeTaskAttributeResponse) {
	response = &DescribeTaskAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
