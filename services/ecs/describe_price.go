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

func (client *Client) DescribePrice(request *DescribePriceRequest) (response *DescribePriceResponse, err error) {
	response = CreateDescribePriceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribePriceWithChan(request *DescribePriceRequest) (<-chan *DescribePriceResponse, <-chan error) {
	responseChan := make(chan *DescribePriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePrice(request)
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

func (client *Client) DescribePriceWithCallback(request *DescribePriceRequest, callback func(response *DescribePriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePriceResponse
		var err error
		defer close(result)
		response, err = client.DescribePrice(request)
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

type DescribePriceRequest struct {
	*requests.RpcRequest
	DataDisk2Size           requests.Integer `position:"Query" name:"DataDisk.2.Size"`
	InternetMaxBandwidthOut requests.Integer `position:"Query" name:"InternetMaxBandwidthOut"`
	DataDisk3Size           requests.Integer `position:"Query" name:"DataDisk.3.Size"`
	SystemDiskCategory      string           `position:"Query" name:"SystemDisk.Category"`
	InternetChargeType      string           `position:"Query" name:"InternetChargeType"`
	DataDisk4Category       string           `position:"Query" name:"DataDisk.4.Category"`
	DataDisk4Size           requests.Integer `position:"Query" name:"DataDisk.4.Size"`
	ResourceType            string           `position:"Query" name:"ResourceType"`
	PriceUnit               string           `position:"Query" name:"PriceUnit"`
	Period                  requests.Integer `position:"Query" name:"Period"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	ImageId                 string           `position:"Query" name:"ImageId"`
	IoOptimized             string           `position:"Query" name:"IoOptimized"`
	InstanceType            string           `position:"Query" name:"InstanceType"`
	Amount                  requests.Integer `position:"Query" name:"Amount"`
	DataDisk1Category       string           `position:"Query" name:"DataDisk.1.Category"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	DataDisk2Category       string           `position:"Query" name:"DataDisk.2.Category"`
	DataDisk1Size           requests.Integer `position:"Query" name:"DataDisk.1.Size"`
	DataDisk3Category       string           `position:"Query" name:"DataDisk.3.Category"`
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
	SystemDiskSize          requests.Integer `position:"Query" name:"SystemDisk.Size"`
	InstanceNetworkType     string           `position:"Query" name:"InstanceNetworkType"`
}

type DescribePriceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	PriceInfo struct {
		Price struct {
			OriginalPrice requests.Float `json:"OriginalPrice" xml:"OriginalPrice"`
			DiscountPrice requests.Float `json:"DiscountPrice" xml:"DiscountPrice"`
			TradePrice    requests.Float `json:"TradePrice" xml:"TradePrice"`
			Currency      string         `json:"Currency" xml:"Currency"`
		} `json:"Price" xml:"Price"`
		Rules struct {
			Rule []struct {
				RuleId      requests.Integer `json:"RuleId" xml:"RuleId"`
				Description string           `json:"Description" xml:"Description"`
			} `json:"Rule" xml:"Rule"`
		} `json:"Rules" xml:"Rules"`
	} `json:"PriceInfo" xml:"PriceInfo"`
}

func CreateDescribePriceRequest() (request *DescribePriceRequest) {
	request = &DescribePriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribePrice", "", "")
	return
}

func CreateDescribePriceResponse() (response *DescribePriceResponse) {
	response = &DescribePriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
