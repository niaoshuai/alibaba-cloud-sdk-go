package cdn

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

func (client *Client) DescribeLiveRecordConfig(request *DescribeLiveRecordConfigRequest) (response *DescribeLiveRecordConfigResponse, err error) {
	response = CreateDescribeLiveRecordConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveRecordConfigWithChan(request *DescribeLiveRecordConfigRequest) (<-chan *DescribeLiveRecordConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveRecordConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveRecordConfig(request)
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

func (client *Client) DescribeLiveRecordConfigWithCallback(request *DescribeLiveRecordConfigRequest, callback func(response *DescribeLiveRecordConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveRecordConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveRecordConfig(request)
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

type DescribeLiveRecordConfigRequest struct {
	*requests.RpcRequest
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type DescribeLiveRecordConfigResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	LiveAppRecordList []struct {
		DomainName      string `json:"DomainName" xml:"DomainName"`
		AppName         string `json:"AppName" xml:"AppName"`
		OssEndpoint     string `json:"OssEndpoint" xml:"OssEndpoint"`
		OssBucket       string `json:"OssBucket" xml:"OssBucket"`
		OssObjectPrefix string `json:"OssObjectPrefix" xml:"OssObjectPrefix"`
		CreateTime      string `json:"CreateTime" xml:"CreateTime"`
	} `json:"LiveAppRecordList" xml:"LiveAppRecordList"`
}

func CreateDescribeLiveRecordConfigRequest() (request *DescribeLiveRecordConfigRequest) {
	request = &DescribeLiveRecordConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveRecordConfig", "", "")
	return
}

func CreateDescribeLiveRecordConfigResponse() (response *DescribeLiveRecordConfigResponse) {
	response = &DescribeLiveRecordConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}