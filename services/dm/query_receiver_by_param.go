package dm

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

func (client *Client) QueryReceiverByParam(request *QueryReceiverByParamRequest) (response *QueryReceiverByParamResponse, err error) {
	response = CreateQueryReceiverByParamResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryReceiverByParamWithChan(request *QueryReceiverByParamRequest) (<-chan *QueryReceiverByParamResponse, <-chan error) {
	responseChan := make(chan *QueryReceiverByParamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryReceiverByParam(request)
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

func (client *Client) QueryReceiverByParamWithCallback(request *QueryReceiverByParamRequest, callback func(response *QueryReceiverByParamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryReceiverByParamResponse
		var err error
		defer close(result)
		response, err = client.QueryReceiverByParam(request)
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

type QueryReceiverByParamRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Status               requests.Integer `position:"Query" name:"Status"`
	KeyWord              string           `position:"Query" name:"KeyWord"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNo               requests.Integer `position:"Query" name:"PageNo"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type QueryReceiverByParamResponse struct {
	*responses.BaseResponse
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	TotalCount requests.Integer `json:"TotalCount" xml:"TotalCount"`
	NextStart  string           `json:"NextStart" xml:"NextStart"`
	PageSize   requests.Integer `json:"PageSize" xml:"PageSize"`
	Data       struct {
		Receiver []struct {
			ReceiverId      string           `json:"ReceiverId" xml:"ReceiverId"`
			ReceiversName   string           `json:"ReceiversName" xml:"ReceiversName"`
			Count           string           `json:"Count" xml:"Count"`
			ReceiversAlias  string           `json:"ReceiversAlias" xml:"ReceiversAlias"`
			Desc            string           `json:"Desc" xml:"Desc"`
			ReceiversStatus string           `json:"ReceiversStatus" xml:"ReceiversStatus"`
			CreateTime      string           `json:"CreateTime" xml:"CreateTime"`
			UtcCreateTime   requests.Integer `json:"UtcCreateTime" xml:"UtcCreateTime"`
		} `json:"receiver" xml:"receiver"`
	} `json:"data" xml:"data"`
}

func CreateQueryReceiverByParamRequest() (request *QueryReceiverByParamRequest) {
	request = &QueryReceiverByParamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "QueryReceiverByParam", "", "")
	return
}

func CreateQueryReceiverByParamResponse() (response *QueryReceiverByParamResponse) {
	response = &QueryReceiverByParamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
