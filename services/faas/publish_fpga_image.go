package faas

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

func (client *Client) PublishFpgaImage(request *PublishFpgaImageRequest) (response *PublishFpgaImageResponse, err error) {
	response = CreatePublishFpgaImageResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) PublishFpgaImageWithChan(request *PublishFpgaImageRequest) (<-chan *PublishFpgaImageResponse, <-chan error) {
	responseChan := make(chan *PublishFpgaImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PublishFpgaImage(request)
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

func (client *Client) PublishFpgaImageWithCallback(request *PublishFpgaImageRequest, callback func(response *PublishFpgaImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PublishFpgaImageResponse
		var err error
		defer close(result)
		response, err = client.PublishFpgaImage(request)
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

type PublishFpgaImageRequest struct {
	*requests.RpcRequest
	CallerUid     string `position:"Query" name:"callerUid"`
	FpgaImageUUID string `position:"Query" name:"FpgaImageUUID"`
	ImageID       string `position:"Query" name:"ImageID"`
}

type PublishFpgaImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Status    string `json:"Status" xml:"Status"`
	Message   string `json:"Message" xml:"Message"`
}

func CreatePublishFpgaImageRequest() (request *PublishFpgaImageRequest) {
	request = &PublishFpgaImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("faas", "2017-08-24", "PublishFpgaImage", "", "")
	return
}

func CreatePublishFpgaImageResponse() (response *PublishFpgaImageResponse) {
	response = &PublishFpgaImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}