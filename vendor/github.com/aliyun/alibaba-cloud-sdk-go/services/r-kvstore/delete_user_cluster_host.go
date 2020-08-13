package r_kvstore

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

// DeleteUserClusterHost invokes the r_kvstore.DeleteUserClusterHost API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/deleteuserclusterhost.html
func (client *Client) DeleteUserClusterHost(request *DeleteUserClusterHostRequest) (response *DeleteUserClusterHostResponse, err error) {
	response = CreateDeleteUserClusterHostResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteUserClusterHostWithChan invokes the r_kvstore.DeleteUserClusterHost API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/deleteuserclusterhost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteUserClusterHostWithChan(request *DeleteUserClusterHostRequest) (<-chan *DeleteUserClusterHostResponse, <-chan error) {
	responseChan := make(chan *DeleteUserClusterHostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteUserClusterHost(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteUserClusterHostWithCallback invokes the r_kvstore.DeleteUserClusterHost API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/deleteuserclusterhost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteUserClusterHostWithCallback(request *DeleteUserClusterHostRequest, callback func(response *DeleteUserClusterHostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteUserClusterHostResponse
		var err error
		defer close(result)
		response, err = client.DeleteUserClusterHost(request)
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

// DeleteUserClusterHostRequest is the request struct for api DeleteUserClusterHost
type DeleteUserClusterHostRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	HostId               string           `position:"Query" name:"HostId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	Engine               string           `position:"Query" name:"Engine"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ClusterId            string           `position:"Query" name:"ClusterId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// DeleteUserClusterHostResponse is the response struct for api DeleteUserClusterHost
type DeleteUserClusterHostResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteUserClusterHostRequest creates a request to invoke DeleteUserClusterHost API
func CreateDeleteUserClusterHostRequest() (request *DeleteUserClusterHostRequest) {
	request = &DeleteUserClusterHostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DeleteUserClusterHost", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteUserClusterHostResponse creates a response to parse from DeleteUserClusterHost response
func CreateDeleteUserClusterHostResponse() (response *DeleteUserClusterHostResponse) {
	response = &DeleteUserClusterHostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
