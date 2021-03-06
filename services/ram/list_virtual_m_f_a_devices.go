package ram

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

func (client *Client) ListVirtualMFADevices(request *ListVirtualMFADevicesRequest) (response *ListVirtualMFADevicesResponse, err error) {
	response = CreateListVirtualMFADevicesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListVirtualMFADevicesWithChan(request *ListVirtualMFADevicesRequest) (<-chan *ListVirtualMFADevicesResponse, <-chan error) {
	responseChan := make(chan *ListVirtualMFADevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVirtualMFADevices(request)
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

func (client *Client) ListVirtualMFADevicesWithCallback(request *ListVirtualMFADevicesRequest, callback func(response *ListVirtualMFADevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVirtualMFADevicesResponse
		var err error
		defer close(result)
		response, err = client.ListVirtualMFADevices(request)
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

type ListVirtualMFADevicesRequest struct {
	*requests.RpcRequest
}

type ListVirtualMFADevicesResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	VirtualMFADevices struct {
		VirtualMFADevice []struct {
			SerialNumber string `json:"SerialNumber" xml:"SerialNumber"`
			ActivateDate string `json:"ActivateDate" xml:"ActivateDate"`
			User         struct {
				UserId      string `json:"UserId" xml:"UserId"`
				UserName    string `json:"UserName" xml:"UserName"`
				DisplayName string `json:"DisplayName" xml:"DisplayName"`
			} `json:"User" xml:"User"`
		} `json:"VirtualMFADevice" xml:"VirtualMFADevice"`
	} `json:"VirtualMFADevices" xml:"VirtualMFADevices"`
}

func CreateListVirtualMFADevicesRequest() (request *ListVirtualMFADevicesRequest) {
	request = &ListVirtualMFADevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "ListVirtualMFADevices", "", "")
	return
}

func CreateListVirtualMFADevicesResponse() (response *ListVirtualMFADevicesResponse) {
	response = &ListVirtualMFADevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
