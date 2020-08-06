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

// SetIgnoreQueryStringConfig invokes the cdn.SetIgnoreQueryStringConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/setignorequerystringconfig.html
func (client *Client) SetIgnoreQueryStringConfig(request *SetIgnoreQueryStringConfigRequest) (response *SetIgnoreQueryStringConfigResponse, err error) {
	response = CreateSetIgnoreQueryStringConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetIgnoreQueryStringConfigWithChan invokes the cdn.SetIgnoreQueryStringConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setignorequerystringconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetIgnoreQueryStringConfigWithChan(request *SetIgnoreQueryStringConfigRequest) (<-chan *SetIgnoreQueryStringConfigResponse, <-chan error) {
	responseChan := make(chan *SetIgnoreQueryStringConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetIgnoreQueryStringConfig(request)
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

// SetIgnoreQueryStringConfigWithCallback invokes the cdn.SetIgnoreQueryStringConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setignorequerystringconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetIgnoreQueryStringConfigWithCallback(request *SetIgnoreQueryStringConfigRequest, callback func(response *SetIgnoreQueryStringConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetIgnoreQueryStringConfigResponse
		var err error
		defer close(result)
		response, err = client.SetIgnoreQueryStringConfig(request)
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

// SetIgnoreQueryStringConfigRequest is the request struct for api SetIgnoreQueryStringConfig
type SetIgnoreQueryStringConfigRequest struct {
	*requests.RpcRequest
	Enable      string           `position:"Query" name:"Enable"`
	KeepOssArgs string           `position:"Query" name:"KeepOssArgs"`
	DomainName  string           `position:"Query" name:"DomainName"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	HashKeyArgs string           `position:"Query" name:"HashKeyArgs"`
	ConfigId    requests.Integer `position:"Query" name:"ConfigId"`
}

// SetIgnoreQueryStringConfigResponse is the response struct for api SetIgnoreQueryStringConfig
type SetIgnoreQueryStringConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetIgnoreQueryStringConfigRequest creates a request to invoke SetIgnoreQueryStringConfig API
func CreateSetIgnoreQueryStringConfigRequest() (request *SetIgnoreQueryStringConfigRequest) {
	request = &SetIgnoreQueryStringConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "SetIgnoreQueryStringConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateSetIgnoreQueryStringConfigResponse creates a response to parse from SetIgnoreQueryStringConfig response
func CreateSetIgnoreQueryStringConfigResponse() (response *SetIgnoreQueryStringConfigResponse) {
	response = &SetIgnoreQueryStringConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
