package cr

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

func (client *Client) CreateRepo(request *CreateRepoRequest) (response *CreateRepoResponse, err error) {
	response = CreateCreateRepoResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateRepoWithChan(request *CreateRepoRequest) (<-chan *CreateRepoResponse, <-chan error) {
	responseChan := make(chan *CreateRepoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRepo(request)
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

func (client *Client) CreateRepoWithCallback(request *CreateRepoRequest, callback func(response *CreateRepoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRepoResponse
		var err error
		defer close(result)
		response, err = client.CreateRepo(request)
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

type CreateRepoRequest struct {
	*requests.RoaRequest
}

type CreateRepoResponse struct {
	*responses.BaseResponse
}

func CreateCreateRepoRequest() (request *CreateRepoRequest) {
	request = &CreateRepoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "CreateRepo", "/repos", "", "")
	return
}

func CreateCreateRepoResponse() (response *CreateRepoResponse) {
	response = &CreateRepoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}