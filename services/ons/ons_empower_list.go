package ons

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

func (client *Client) OnsEmpowerList(request *OnsEmpowerListRequest) (response *OnsEmpowerListResponse, err error) {
	response = CreateOnsEmpowerListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) OnsEmpowerListWithChan(request *OnsEmpowerListRequest) (<-chan *OnsEmpowerListResponse, <-chan error) {
	responseChan := make(chan *OnsEmpowerListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsEmpowerList(request)
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

func (client *Client) OnsEmpowerListWithCallback(request *OnsEmpowerListRequest, callback func(response *OnsEmpowerListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsEmpowerListResponse
		var err error
		defer close(result)
		response, err = client.OnsEmpowerList(request)
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

type OnsEmpowerListRequest struct {
	*requests.RpcRequest
	Topic        string `position:"Query" name:"Topic"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	EmpowerUser  string `position:"Query" name:"EmpowerUser"`
	PreventCache string `position:"Query" name:"PreventCache"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

type OnsEmpowerListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	Data      []struct {
		Topic    string `json:"Topic" xml:"Topic"`
		Owner    int64  `json:"Owner" xml:"Owner"`
		Relation int    `json:"Relation" xml:"Relation"`
	} `json:"Data" xml:"Data"`
}

func CreateOnsEmpowerListRequest() (request *OnsEmpowerListRequest) {
	request = &OnsEmpowerListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2017-09-18", "OnsEmpowerList", "", "")
	return
}

func CreateOnsEmpowerListResponse() (response *OnsEmpowerListResponse) {
	response = &OnsEmpowerListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}