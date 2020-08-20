package domain

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

// DeleteRegistrantProfile invokes the domain.DeleteRegistrantProfile API synchronously
// api document: https://help.aliyun.com/api/domain/deleteregistrantprofile.html
func (client *Client) DeleteRegistrantProfile(request *DeleteRegistrantProfileRequest) (response *DeleteRegistrantProfileResponse, err error) {
	response = CreateDeleteRegistrantProfileResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRegistrantProfileWithChan invokes the domain.DeleteRegistrantProfile API asynchronously
// api document: https://help.aliyun.com/api/domain/deleteregistrantprofile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRegistrantProfileWithChan(request *DeleteRegistrantProfileRequest) (<-chan *DeleteRegistrantProfileResponse, <-chan error) {
	responseChan := make(chan *DeleteRegistrantProfileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRegistrantProfile(request)
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

// DeleteRegistrantProfileWithCallback invokes the domain.DeleteRegistrantProfile API asynchronously
// api document: https://help.aliyun.com/api/domain/deleteregistrantprofile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRegistrantProfileWithCallback(request *DeleteRegistrantProfileRequest, callback func(response *DeleteRegistrantProfileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRegistrantProfileResponse
		var err error
		defer close(result)
		response, err = client.DeleteRegistrantProfile(request)
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

// DeleteRegistrantProfileRequest is the request struct for api DeleteRegistrantProfile
type DeleteRegistrantProfileRequest struct {
	*requests.RpcRequest
	RegistrantProfileId requests.Integer `position:"Query" name:"RegistrantProfileId"`
	UserClientIp        string           `position:"Query" name:"UserClientIp"`
	Lang                string           `position:"Query" name:"Lang"`
}

// DeleteRegistrantProfileResponse is the response struct for api DeleteRegistrantProfile
type DeleteRegistrantProfileResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteRegistrantProfileRequest creates a request to invoke DeleteRegistrantProfile API
func CreateDeleteRegistrantProfileRequest() (request *DeleteRegistrantProfileRequest) {
	request = &DeleteRegistrantProfileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "DeleteRegistrantProfile", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteRegistrantProfileResponse creates a response to parse from DeleteRegistrantProfile response
func CreateDeleteRegistrantProfileResponse() (response *DeleteRegistrantProfileResponse) {
	response = &DeleteRegistrantProfileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
