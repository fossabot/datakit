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

// CheckDBNameAvailable invokes the rds.CheckDBNameAvailable API synchronously
// api document: https://help.aliyun.com/api/rds/checkdbnameavailable.html
func (client *Client) CheckDBNameAvailable(request *CheckDBNameAvailableRequest) (response *CheckDBNameAvailableResponse, err error) {
	response = CreateCheckDBNameAvailableResponse()
	err = client.DoAction(request, response)
	return
}

// CheckDBNameAvailableWithChan invokes the rds.CheckDBNameAvailable API asynchronously
// api document: https://help.aliyun.com/api/rds/checkdbnameavailable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckDBNameAvailableWithChan(request *CheckDBNameAvailableRequest) (<-chan *CheckDBNameAvailableResponse, <-chan error) {
	responseChan := make(chan *CheckDBNameAvailableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckDBNameAvailable(request)
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

// CheckDBNameAvailableWithCallback invokes the rds.CheckDBNameAvailable API asynchronously
// api document: https://help.aliyun.com/api/rds/checkdbnameavailable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckDBNameAvailableWithCallback(request *CheckDBNameAvailableRequest, callback func(response *CheckDBNameAvailableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckDBNameAvailableResponse
		var err error
		defer close(result)
		response, err = client.CheckDBNameAvailable(request)
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

// CheckDBNameAvailableRequest is the request struct for api CheckDBNameAvailable
type CheckDBNameAvailableRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBName               string           `position:"Query" name:"DBName"`
}

// CheckDBNameAvailableResponse is the response struct for api CheckDBNameAvailable
type CheckDBNameAvailableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCheckDBNameAvailableRequest creates a request to invoke CheckDBNameAvailable API
func CreateCheckDBNameAvailableRequest() (request *CheckDBNameAvailableRequest) {
	request = &CheckDBNameAvailableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CheckDBNameAvailable", "rds", "openAPI")
	return
}

// CreateCheckDBNameAvailableResponse creates a response to parse from CheckDBNameAvailable response
func CreateCheckDBNameAvailableResponse() (response *CheckDBNameAvailableResponse) {
	response = &CheckDBNameAvailableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
