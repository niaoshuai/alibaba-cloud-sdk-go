package cms

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

func (client *Client) NodeProcessCreate(request *NodeProcessCreateRequest) (response *NodeProcessCreateResponse, err error) {
	response = CreateNodeProcessCreateResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) NodeProcessCreateWithChan(request *NodeProcessCreateRequest) (<-chan *NodeProcessCreateResponse, <-chan error) {
	responseChan := make(chan *NodeProcessCreateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.NodeProcessCreate(request)
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

func (client *Client) NodeProcessCreateWithCallback(request *NodeProcessCreateRequest, callback func(response *NodeProcessCreateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *NodeProcessCreateResponse
		var err error
		defer close(result)
		response, err = client.NodeProcessCreate(request)
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

type NodeProcessCreateRequest struct {
	*requests.RpcRequest
	ProcessName string `position:"Query" name:"ProcessName"`
	Name        string `position:"Query" name:"Name"`
	Command     string `position:"Query" name:"Command"`
	ProcessUser string `position:"Query" name:"ProcessUser"`
	InstanceId  string `position:"Query" name:"InstanceId"`
}

type NodeProcessCreateResponse struct {
	*responses.BaseResponse
	ErrorCode    requests.Integer `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string           `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      requests.Boolean `json:"Success" xml:"Success"`
	RequestId    string           `json:"RequestId" xml:"RequestId"`
}

func CreateNodeProcessCreateRequest() (request *NodeProcessCreateRequest) {
	request = &NodeProcessCreateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "NodeProcessCreate", "", "")
	return
}

func CreateNodeProcessCreateResponse() (response *NodeProcessCreateResponse) {
	response = &NodeProcessCreateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
