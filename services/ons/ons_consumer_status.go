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

func (client *Client) OnsConsumerStatus(request *OnsConsumerStatusRequest) (response *OnsConsumerStatusResponse, err error) {
	response = CreateOnsConsumerStatusResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) OnsConsumerStatusWithChan(request *OnsConsumerStatusRequest) (<-chan *OnsConsumerStatusResponse, <-chan error) {
	responseChan := make(chan *OnsConsumerStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsConsumerStatus(request)
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

func (client *Client) OnsConsumerStatusWithCallback(request *OnsConsumerStatusRequest, callback func(response *OnsConsumerStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsConsumerStatusResponse
		var err error
		defer close(result)
		response, err = client.OnsConsumerStatus(request)
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

type OnsConsumerStatusRequest struct {
	*requests.RpcRequest
	Detail       string `position:"Query" name:"Detail"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	PreventCache string `position:"Query" name:"PreventCache"`
	NeedJstack   string `position:"Query" name:"NeedJstack"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

type OnsConsumerStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	Data      struct {
		Online           bool    `json:"Online" xml:"Online"`
		TotalDiff        int64   `json:"TotalDiff" xml:"TotalDiff"`
		ConsumeTps       float64 `json:"ConsumeTps" xml:"ConsumeTps"`
		LastTimestamp    int64   `json:"LastTimestamp" xml:"LastTimestamp"`
		DelayTime        int64   `json:"DelayTime" xml:"DelayTime"`
		ConsumeModel     string  `json:"ConsumeModel" xml:"ConsumeModel"`
		SubscriptionSame bool    `json:"SubscriptionSame" xml:"SubscriptionSame"`
		RebalanceOK      bool    `json:"RebalanceOK" xml:"RebalanceOK"`
		ConnectionSet    []struct {
			ClientId   string `json:"ClientId" xml:"ClientId"`
			ClientAddr string `json:"ClientAddr" xml:"ClientAddr"`
			Language   string `json:"Language" xml:"Language"`
			Version    string `json:"Version" xml:"Version"`
		} `json:"ConnectionSet" xml:"ConnectionSet"`
		DetailInTopicList []struct {
			Topic         string `json:"Topic" xml:"Topic"`
			TotalDiff     int64  `json:"TotalDiff" xml:"TotalDiff"`
			LastTimestamp int64  `json:"LastTimestamp" xml:"LastTimestamp"`
			DelayTime     int64  `json:"DelayTime" xml:"DelayTime"`
		} `json:"DetailInTopicList" xml:"DetailInTopicList"`
		ConsumerConnectionInfoList []struct {
			ClientId        string `json:"ClientId" xml:"ClientId"`
			Connection      string `json:"Connection" xml:"Connection"`
			Language        string `json:"Language" xml:"Language"`
			Version         string `json:"Version" xml:"Version"`
			ConsumeModel    string `json:"ConsumeModel" xml:"ConsumeModel"`
			ConsumeType     string `json:"ConsumeType" xml:"ConsumeType"`
			ThreadCount     int    `json:"ThreadCount" xml:"ThreadCount"`
			StartTimeStamp  int64  `json:"StartTimeStamp" xml:"StartTimeStamp"`
			LastTimeStamp   int64  `json:"LastTimeStamp" xml:"LastTimeStamp"`
			SubscriptionSet []struct {
				Topic      string   `json:"Topic" xml:"Topic"`
				SubString  string   `json:"SubString" xml:"SubString"`
				SubVersion int64    `json:"SubVersion" xml:"SubVersion"`
				TagsSet    []string `json:"TagsSet" xml:"TagsSet"`
			} `json:"SubscriptionSet" xml:"SubscriptionSet"`
			RunningDataList []struct {
				ConsumerId         string  `json:"ConsumerId" xml:"ConsumerId"`
				Topic              string  `json:"Topic" xml:"Topic"`
				Rt                 float64 `json:"Rt" xml:"Rt"`
				OkTps              float64 `json:"OkTps" xml:"OkTps"`
				FailedTps          float64 `json:"FailedTps" xml:"FailedTps"`
				FailedCountPerHour int64   `json:"FailedCountPerHour" xml:"FailedCountPerHour"`
			} `json:"RunningDataList" xml:"RunningDataList"`
			Jstack []struct {
				Thread    string   `json:"Thread" xml:"Thread"`
				TrackList []string `json:"TrackList" xml:"TrackList"`
			} `json:"Jstack" xml:"Jstack"`
		} `json:"ConsumerConnectionInfoList" xml:"ConsumerConnectionInfoList"`
	} `json:"Data" xml:"Data"`
}

func CreateOnsConsumerStatusRequest() (request *OnsConsumerStatusRequest) {
	request = &OnsConsumerStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2017-09-18", "OnsConsumerStatus", "", "")
	return
}

func CreateOnsConsumerStatusResponse() (response *OnsConsumerStatusResponse) {
	response = &OnsConsumerStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}