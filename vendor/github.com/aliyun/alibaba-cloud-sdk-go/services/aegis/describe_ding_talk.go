package aegis

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

// DescribeDingTalk invokes the aegis.DescribeDingTalk API synchronously
// api document: https://help.aliyun.com/api/aegis/describedingtalk.html
func (client *Client) DescribeDingTalk(request *DescribeDingTalkRequest) (response *DescribeDingTalkResponse, err error) {
	response = CreateDescribeDingTalkResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDingTalkWithChan invokes the aegis.DescribeDingTalk API asynchronously
// api document: https://help.aliyun.com/api/aegis/describedingtalk.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDingTalkWithChan(request *DescribeDingTalkRequest) (<-chan *DescribeDingTalkResponse, <-chan error) {
	responseChan := make(chan *DescribeDingTalkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDingTalk(request)
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

// DescribeDingTalkWithCallback invokes the aegis.DescribeDingTalk API asynchronously
// api document: https://help.aliyun.com/api/aegis/describedingtalk.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDingTalkWithCallback(request *DescribeDingTalkRequest, callback func(response *DescribeDingTalkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDingTalkResponse
		var err error
		defer close(result)
		response, err = client.DescribeDingTalk(request)
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

// DescribeDingTalkRequest is the request struct for api DescribeDingTalk
type DescribeDingTalkRequest struct {
	*requests.RpcRequest
	RuleActionName string `position:"Query" name:"RuleActionName"`
	SourceIp       string `position:"Query" name:"SourceIp"`
}

// DescribeDingTalkResponse is the response struct for api DescribeDingTalk
type DescribeDingTalkResponse struct {
	*responses.BaseResponse
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	PageInfo   PageInfo         `json:"PageInfo" xml:"PageInfo"`
	ActionList []ActionListItem `json:"ActionList" xml:"ActionList"`
}

// CreateDescribeDingTalkRequest creates a request to invoke DescribeDingTalk API
func CreateDescribeDingTalkRequest() (request *DescribeDingTalkRequest) {
	request = &DescribeDingTalkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeDingTalk", "vipaegis", "openAPI")
	return
}

// CreateDescribeDingTalkResponse creates a response to parse from DescribeDingTalk response
func CreateDescribeDingTalkResponse() (response *DescribeDingTalkResponse) {
	response = &DescribeDingTalkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
