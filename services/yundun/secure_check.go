package yundun

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

func (client *Client) SecureCheck(request *SecureCheckRequest) (response *SecureCheckResponse, err error) {
	response = CreateSecureCheckResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SecureCheckWithChan(request *SecureCheckRequest) (<-chan *SecureCheckResponse, <-chan error) {
	responseChan := make(chan *SecureCheckResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SecureCheck(request)
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

func (client *Client) SecureCheckWithCallback(request *SecureCheckRequest, callback func(response *SecureCheckResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SecureCheckResponse
		var err error
		defer close(result)
		response, err = client.SecureCheck(request)
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

type SecureCheckRequest struct {
	*requests.RpcRequest
	InstanceIds string `position:"Query" name:"InstanceIds"`
	JstOwnerId  string `position:"Query" name:"JstOwnerId"`
}

type SecureCheckResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	RecentInstanceId string `json:"RecentInstanceId" xml:"RecentInstanceId"`
	ProblemList      []struct {
		Ip         string `json:"Ip" xml:"Ip"`
		Status     string `json:"Status" xml:"Status"`
		VulNum     string `json:"VulNum" xml:"VulNum"`
		InstanceId string `json:"InstanceId" xml:"InstanceId"`
	} `json:"ProblemList" xml:"ProblemList"`
	NoProblemList []struct {
		Ip         string `json:"Ip" xml:"Ip"`
		Status     string `json:"Status" xml:"Status"`
		VulNum     string `json:"VulNum" xml:"VulNum"`
		InstanceId string `json:"InstanceId" xml:"InstanceId"`
	} `json:"NoProblemList" xml:"NoProblemList"`
	NoScanList []struct {
		Ip         string `json:"Ip" xml:"Ip"`
		Status     string `json:"Status" xml:"Status"`
		VulNum     string `json:"VulNum" xml:"VulNum"`
		InstanceId string `json:"InstanceId" xml:"InstanceId"`
	} `json:"NoScanList" xml:"NoScanList"`
	ScanningList []struct {
		Ip         string `json:"Ip" xml:"Ip"`
		Status     string `json:"Status" xml:"Status"`
		VulNum     string `json:"VulNum" xml:"VulNum"`
		InstanceId string `json:"InstanceId" xml:"InstanceId"`
	} `json:"ScanningList" xml:"ScanningList"`
	InnerIpList []struct {
		Ip         string `json:"Ip" xml:"Ip"`
		Status     string `json:"Status" xml:"Status"`
		VulNum     string `json:"VulNum" xml:"VulNum"`
		InstanceId string `json:"InstanceId" xml:"InstanceId"`
	} `json:"InnerIpList" xml:"InnerIpList"`
}

func CreateSecureCheckRequest() (request *SecureCheckRequest) {
	request = &SecureCheckRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun", "2015-04-16", "SecureCheck", "", "")
	return
}

func CreateSecureCheckResponse() (response *SecureCheckResponse) {
	response = &SecureCheckResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}