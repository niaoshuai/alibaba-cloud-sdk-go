package live

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

func (client *Client) DescribeCasterConfig(request *DescribeCasterConfigRequest) (response *DescribeCasterConfigResponse, err error) {
	response = CreateDescribeCasterConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeCasterConfigWithChan(request *DescribeCasterConfigRequest) (<-chan *DescribeCasterConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeCasterConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCasterConfig(request)
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

func (client *Client) DescribeCasterConfigWithCallback(request *DescribeCasterConfigRequest, callback func(response *DescribeCasterConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCasterConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeCasterConfig(request)
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

type DescribeCasterConfigRequest struct {
	*requests.RpcRequest
	OwnerId       string `position:"Query" name:"OwnerId"`
	CasterId      string `position:"Query" name:"CasterId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Version       string `position:"Query" name:"Version"`
}

type DescribeCasterConfigResponse struct {
	*responses.BaseResponse
	RequestId        string  `json:"RequestId" xml:"RequestId"`
	CasterId         string  `json:"CasterId" xml:"CasterId"`
	CasterName       string  `json:"CasterName" xml:"CasterName"`
	DomainName       string  `json:"DomainName" xml:"DomainName"`
	Delay            float64 `json:"Delay" xml:"Delay"`
	UrgentMaterialId string  `json:"UrgentMaterialId" xml:"UrgentMaterialId"`
	TranscodeConfig  struct {
		CasterTemplate  string   `json:"CasterTemplate" xml:"CasterTemplate"`
		LiveTemplateIds []string `json:"LiveTemplateIds" xml:"LiveTemplateIds"`
	} `json:"TranscodeConfig" xml:"TranscodeConfig"`
	RecordConfig struct {
		OssEndpoint  string `json:"OssEndpoint" xml:"OssEndpoint"`
		OssBucket    string `json:"OssBucket" xml:"OssBucket"`
		RecordFormat []struct {
			Format               string `json:"Format" xml:"Format"`
			OssObjectPrefix      string `json:"OssObjectPrefix" xml:"OssObjectPrefix"`
			SliceOssObjectPrefix string `json:"SliceOssObjectPrefix" xml:"SliceOssObjectPrefix"`
			CycleDuration        int    `json:"CycleDuration" xml:"CycleDuration"`
		} `json:"RecordFormat" xml:"RecordFormat"`
	} `json:"RecordConfig" xml:"RecordConfig"`
}

func CreateDescribeCasterConfigRequest() (request *DescribeCasterConfigRequest) {
	request = &DescribeCasterConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeCasterConfig", "", "")
	return
}

func CreateDescribeCasterConfigResponse() (response *DescribeCasterConfigResponse) {
	response = &DescribeCasterConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}