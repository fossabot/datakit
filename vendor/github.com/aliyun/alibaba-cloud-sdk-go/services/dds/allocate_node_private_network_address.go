package dds

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

// AllocateNodePrivateNetworkAddress invokes the dds.AllocateNodePrivateNetworkAddress API synchronously
// api document: https://help.aliyun.com/api/dds/allocatenodeprivatenetworkaddress.html
func (client *Client) AllocateNodePrivateNetworkAddress(request *AllocateNodePrivateNetworkAddressRequest) (response *AllocateNodePrivateNetworkAddressResponse, err error) {
	response = CreateAllocateNodePrivateNetworkAddressResponse()
	err = client.DoAction(request, response)
	return
}

// AllocateNodePrivateNetworkAddressWithChan invokes the dds.AllocateNodePrivateNetworkAddress API asynchronously
// api document: https://help.aliyun.com/api/dds/allocatenodeprivatenetworkaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AllocateNodePrivateNetworkAddressWithChan(request *AllocateNodePrivateNetworkAddressRequest) (<-chan *AllocateNodePrivateNetworkAddressResponse, <-chan error) {
	responseChan := make(chan *AllocateNodePrivateNetworkAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllocateNodePrivateNetworkAddress(request)
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

// AllocateNodePrivateNetworkAddressWithCallback invokes the dds.AllocateNodePrivateNetworkAddress API asynchronously
// api document: https://help.aliyun.com/api/dds/allocatenodeprivatenetworkaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AllocateNodePrivateNetworkAddressWithCallback(request *AllocateNodePrivateNetworkAddressRequest, callback func(response *AllocateNodePrivateNetworkAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllocateNodePrivateNetworkAddressResponse
		var err error
		defer close(result)
		response, err = client.AllocateNodePrivateNetworkAddress(request)
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

// AllocateNodePrivateNetworkAddressRequest is the request struct for api AllocateNodePrivateNetworkAddress
type AllocateNodePrivateNetworkAddressRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountName          string           `position:"Query" name:"AccountName"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	NodeId               string           `position:"Query" name:"NodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AccountPassword      string           `position:"Query" name:"AccountPassword"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// AllocateNodePrivateNetworkAddressResponse is the response struct for api AllocateNodePrivateNetworkAddress
type AllocateNodePrivateNetworkAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAllocateNodePrivateNetworkAddressRequest creates a request to invoke AllocateNodePrivateNetworkAddress API
func CreateAllocateNodePrivateNetworkAddressRequest() (request *AllocateNodePrivateNetworkAddressRequest) {
	request = &AllocateNodePrivateNetworkAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "AllocateNodePrivateNetworkAddress", "Dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAllocateNodePrivateNetworkAddressResponse creates a response to parse from AllocateNodePrivateNetworkAddress response
func CreateAllocateNodePrivateNetworkAddressResponse() (response *AllocateNodePrivateNetworkAddressResponse) {
	response = &AllocateNodePrivateNetworkAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
