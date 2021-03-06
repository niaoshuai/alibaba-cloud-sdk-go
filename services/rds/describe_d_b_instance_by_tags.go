package rds

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

func (client *Client) DescribeDBInstanceByTags(request *DescribeDBInstanceByTagsRequest) (response *DescribeDBInstanceByTagsResponse, err error) {
	response = CreateDescribeDBInstanceByTagsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDBInstanceByTagsWithChan(request *DescribeDBInstanceByTagsRequest) (<-chan *DescribeDBInstanceByTagsResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceByTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceByTags(request)
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

func (client *Client) DescribeDBInstanceByTagsWithCallback(request *DescribeDBInstanceByTagsRequest, callback func(response *DescribeDBInstanceByTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceByTagsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceByTags(request)
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

type DescribeDBInstanceByTagsRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ProxyId              string           `position:"Query" name:"proxyId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeDBInstanceByTagsResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	PageNumber       requests.Integer `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  requests.Integer `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalRecordCount requests.Integer `json:"TotalRecordCount" xml:"TotalRecordCount"`
	Items            struct {
		DBInstanceTag []struct {
			DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
			Tags         struct {
				Tag []struct {
					TagKey   string `json:"TagKey" xml:"TagKey"`
					TagValue string `json:"TagValue" xml:"TagValue"`
				} `json:"Tag" xml:"Tag"`
			} `json:"Tags" xml:"Tags"`
		} `json:"DBInstanceTag" xml:"DBInstanceTag"`
	} `json:"Items" xml:"Items"`
}

func CreateDescribeDBInstanceByTagsRequest() (request *DescribeDBInstanceByTagsRequest) {
	request = &DescribeDBInstanceByTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceByTags", "", "")
	return
}

func CreateDescribeDBInstanceByTagsResponse() (response *DescribeDBInstanceByTagsResponse) {
	response = &DescribeDBInstanceByTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
