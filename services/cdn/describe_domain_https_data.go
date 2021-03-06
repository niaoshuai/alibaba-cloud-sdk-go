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

func (client *Client) DescribeDomainHttpsData(request *DescribeDomainHttpsDataRequest) (response *DescribeDomainHttpsDataResponse, err error) {
	response = CreateDescribeDomainHttpsDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainHttpsDataWithChan(request *DescribeDomainHttpsDataRequest) (<-chan *DescribeDomainHttpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainHttpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainHttpsData(request)
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

func (client *Client) DescribeDomainHttpsDataWithCallback(request *DescribeDomainHttpsDataRequest, callback func(response *DescribeDomainHttpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainHttpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainHttpsData(request)
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

type DescribeDomainHttpsDataRequest struct {
	*requests.RpcRequest
	EndTime       string           `position:"Query" name:"EndTime"`
	Cls           string           `position:"Query" name:"Cls"`
	StartTime     string           `position:"Query" name:"StartTime"`
	Interval      string           `position:"Query" name:"Interval"`
	FixTimeGap    string           `position:"Query" name:"FixTimeGap"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainType    string           `position:"Query" name:"DomainType"`
	DomainNames   string           `position:"Query" name:"DomainNames"`
	TimeMerge     string           `position:"Query" name:"TimeMerge"`
}

type DescribeDomainHttpsDataResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	DomainNames          string `json:"DomainNames" xml:"DomainNames"`
	DataInterval         string `json:"DataInterval" xml:"DataInterval"`
	HttpsStatisticsInfos struct {
		HttpsStatisticsInfo []struct {
			Time               string           `json:"Time" xml:"Time"`
			L1HttpsBps         requests.Float   `json:"L1HttpsBps" xml:"L1HttpsBps"`
			L1HttpsInnerBps    requests.Float   `json:"L1HttpsInnerBps" xml:"L1HttpsInnerBps"`
			L1HttpsOutBps      requests.Float   `json:"L1HttpsOutBps" xml:"L1HttpsOutBps"`
			L1HttpsQps         requests.Integer `json:"L1HttpsQps" xml:"L1HttpsQps"`
			L1HttpsInnerQps    requests.Integer `json:"L1HttpsInnerQps" xml:"L1HttpsInnerQps"`
			L1HttpsOutQps      requests.Integer `json:"L1HttpsOutQps" xml:"L1HttpsOutQps"`
			L1HttpsTtraf       requests.Integer `json:"L1HttpsTtraf" xml:"L1HttpsTtraf"`
			HttpsSrcBps        requests.Integer `json:"HttpsSrcBps" xml:"HttpsSrcBps"`
			HttpsSrcTraf       requests.Integer `json:"HttpsSrcTraf" xml:"HttpsSrcTraf"`
			L1HttpsInnerTraf   requests.Integer `json:"L1HttpsInnerTraf" xml:"L1HttpsInnerTraf"`
			L1HttpsOutTraf     requests.Integer `json:"L1HttpsOutTraf" xml:"L1HttpsOutTraf"`
			HttpsByteHitRate   requests.Float   `json:"HttpsByteHitRate" xml:"HttpsByteHitRate"`
			HttpsReqHitRate    requests.Float   `json:"HttpsReqHitRate" xml:"HttpsReqHitRate"`
			L1HttpsHitRate     requests.Float   `json:"L1HttpsHitRate" xml:"L1HttpsHitRate"`
			L1HttpsInnerAcc    requests.Float   `json:"L1HttpsInner_acc" xml:"L1HttpsInner_acc"`
			L1HttpsOutAcc      requests.Float   `json:"L1HttpsOut_acc" xml:"L1HttpsOut_acc"`
			L1HttpsTacc        requests.Float   `json:"L1HttpsTacc" xml:"L1HttpsTacc"`
			L1DyHttpsBps       requests.Float   `json:"L1DyHttpsBps" xml:"L1DyHttpsBps"`
			L1DyHttpsInnerBps  requests.Float   `json:"L1DyHttpsInnerBps" xml:"L1DyHttpsInnerBps"`
			L1DyHttpsOutBps    requests.Float   `json:"L1DyHttpsOutBps" xml:"L1DyHttpsOutBps"`
			L1StHttpsBps       requests.Float   `json:"L1StHttpsBps" xml:"L1StHttpsBps"`
			L1StHttpsInnerBps  requests.Float   `json:"L1StHttpsInnerBps" xml:"L1StHttpsInnerBps"`
			L1StHttpsOutBps    requests.Float   `json:"L1StHttpsOutBps" xml:"L1StHttpsOutBps"`
			L1DyHttpsTraf      requests.Float   `json:"l1DyHttpsTraf" xml:"l1DyHttpsTraf"`
			L1DyHttpsInnerTraf requests.Float   `json:"L1DyHttpsInnerTraf" xml:"L1DyHttpsInnerTraf"`
			L1DyHttpsOutTraf   requests.Float   `json:"L1DyHttpsOutTraf" xml:"L1DyHttpsOutTraf"`
			L1StHttpsTraf      requests.Float   `json:"L1StHttpsTraf" xml:"L1StHttpsTraf"`
			L1StHttpsInnerTraf requests.Float   `json:"L1StHttpsInnerTraf" xml:"L1StHttpsInnerTraf"`
			L1StHttpsOutTraf   requests.Float   `json:"L1StHttpsOutTraf" xml:"L1StHttpsOutTraf"`
			L1DyHttpsQps       requests.Float   `json:"L1DyHttpsQps" xml:"L1DyHttpsQps"`
			L1DyHttpsInnerQps  requests.Float   `json:"L1DyHttpsInnerQps" xml:"L1DyHttpsInnerQps"`
			L1DyHttpsOutQps    requests.Float   `json:"L1DyHttpsOutQps" xml:"L1DyHttpsOutQps"`
			L1StHttpsQps       requests.Float   `json:"L1StHttpsQps" xml:"L1StHttpsQps"`
			L1StHttpsInnerQps  requests.Float   `json:"L1StHttpsInnerQps" xml:"L1StHttpsInnerQps"`
			L1StHttpsOutQps    requests.Float   `json:"L1StHttpsOutQps" xml:"L1StHttpsOutQps"`
			L1DyHttpsAcc       requests.Float   `json:"L1DyHttpsAcc" xml:"L1DyHttpsAcc"`
			L1DyHttpsInnerAcc  requests.Float   `json:"L1DyHttpsInnerAcc" xml:"L1DyHttpsInnerAcc"`
			L1DyHttpsOutAcc    requests.Float   `json:"L1DyHttpsOutAcc" xml:"L1DyHttpsOutAcc"`
			L1StHttpsAcc       requests.Float   `json:"L1StHttpsAcc" xml:"L1StHttpsAcc"`
			L1StHttpsInnerAcc  requests.Float   `json:"L1StHttpsInnerAcc" xml:"L1StHttpsInnerAcc"`
			L1StHttpsOutAcc    requests.Float   `json:"L1StHttpsOutAcc" xml:"L1StHttpsOutAcc"`
		} `json:"HttpsStatisticsInfo" xml:"HttpsStatisticsInfo"`
	} `json:"HttpsStatisticsInfos" xml:"HttpsStatisticsInfos"`
}

func CreateDescribeDomainHttpsDataRequest() (request *DescribeDomainHttpsDataRequest) {
	request = &DescribeDomainHttpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainHttpsData", "", "")
	return
}

func CreateDescribeDomainHttpsDataResponse() (response *DescribeDomainHttpsDataResponse) {
	response = &DescribeDomainHttpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
