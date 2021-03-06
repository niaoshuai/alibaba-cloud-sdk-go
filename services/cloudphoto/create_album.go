package cloudphoto

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

func (client *Client) CreateAlbum(request *CreateAlbumRequest) (response *CreateAlbumResponse, err error) {
	response = CreateCreateAlbumResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateAlbumWithChan(request *CreateAlbumRequest) (<-chan *CreateAlbumResponse, <-chan error) {
	responseChan := make(chan *CreateAlbumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAlbum(request)
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

func (client *Client) CreateAlbumWithCallback(request *CreateAlbumRequest, callback func(response *CreateAlbumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAlbumResponse
		var err error
		defer close(result)
		response, err = client.CreateAlbum(request)
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

type CreateAlbumRequest struct {
	*requests.RpcRequest
	Remark    string `position:"Query" name:"Remark"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	AlbumName string `position:"Query" name:"AlbumName"`
}

type CreateAlbumResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Album     struct {
		Id          requests.Integer `json:"Id" xml:"Id"`
		Name        string           `json:"Name" xml:"Name"`
		State       string           `json:"State" xml:"State"`
		Remark      string           `json:"Remark" xml:"Remark"`
		PhotosCount requests.Integer `json:"PhotosCount" xml:"PhotosCount"`
		Ctime       requests.Integer `json:"Ctime" xml:"Ctime"`
		Mtime       requests.Integer `json:"Mtime" xml:"Mtime"`
		Cover       struct {
			Id      requests.Integer `json:"Id" xml:"Id"`
			Title   string           `json:"Title" xml:"Title"`
			FileId  string           `json:"FileId" xml:"FileId"`
			State   string           `json:"State" xml:"State"`
			Md5     string           `json:"Md5" xml:"Md5"`
			IsVideo requests.Boolean `json:"IsVideo" xml:"IsVideo"`
			Width   requests.Integer `json:"Width" xml:"Width"`
			Height  requests.Integer `json:"Height" xml:"Height"`
			Ctime   requests.Integer `json:"Ctime" xml:"Ctime"`
			Mtime   requests.Integer `json:"Mtime" xml:"Mtime"`
			Remark  string           `json:"Remark" xml:"Remark"`
		} `json:"Cover" xml:"Cover"`
	} `json:"Album" xml:"Album"`
}

func CreateCreateAlbumRequest() (request *CreateAlbumRequest) {
	request = &CreateAlbumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "CreateAlbum", "", "")
	return
}

func CreateCreateAlbumResponse() (response *CreateAlbumResponse) {
	response = &CreateAlbumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
