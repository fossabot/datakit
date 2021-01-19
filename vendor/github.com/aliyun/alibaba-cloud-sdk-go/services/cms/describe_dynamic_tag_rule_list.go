package cms

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

// DescribeDynamicTagRuleList invokes the cms.DescribeDynamicTagRuleList API synchronously
// api document: https://help.aliyun.com/api/cms/describedynamictagrulelist.html
func (client *Client) DescribeDynamicTagRuleList(request *DescribeDynamicTagRuleListRequest) (response *DescribeDynamicTagRuleListResponse, err error) {
	response = CreateDescribeDynamicTagRuleListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDynamicTagRuleListWithChan invokes the cms.DescribeDynamicTagRuleList API asynchronously
// api document: https://help.aliyun.com/api/cms/describedynamictagrulelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDynamicTagRuleListWithChan(request *DescribeDynamicTagRuleListRequest) (<-chan *DescribeDynamicTagRuleListResponse, <-chan error) {
	responseChan := make(chan *DescribeDynamicTagRuleListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDynamicTagRuleList(request)
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

// DescribeDynamicTagRuleListWithCallback invokes the cms.DescribeDynamicTagRuleList API asynchronously
// api document: https://help.aliyun.com/api/cms/describedynamictagrulelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDynamicTagRuleListWithCallback(request *DescribeDynamicTagRuleListRequest, callback func(response *DescribeDynamicTagRuleListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDynamicTagRuleListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDynamicTagRuleList(request)
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

// DescribeDynamicTagRuleListRequest is the request struct for api DescribeDynamicTagRuleList
type DescribeDynamicTagRuleListRequest struct {
	*requests.RpcRequest
	PageNumber string `position:"Query" name:"PageNumber"`
	PageSize   string `position:"Query" name:"PageSize"`
	TagKey     string `position:"Query" name:"TagKey"`
}

// DescribeDynamicTagRuleListResponse is the response struct for api DescribeDynamicTagRuleList
type DescribeDynamicTagRuleListResponse struct {
	*responses.BaseResponse
	Success      bool         `json:"Success" xml:"Success"`
	Code         string       `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Total        int          `json:"Total" xml:"Total"`
	PageNumber   string       `json:"PageNumber" xml:"PageNumber"`
	PageSize     string       `json:"PageSize" xml:"PageSize"`
	TagGroupList TagGroupList `json:"TagGroupList" xml:"TagGroupList"`
}

// CreateDescribeDynamicTagRuleListRequest creates a request to invoke DescribeDynamicTagRuleList API
func CreateDescribeDynamicTagRuleListRequest() (request *DescribeDynamicTagRuleListRequest) {
	request = &DescribeDynamicTagRuleListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeDynamicTagRuleList", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDynamicTagRuleListResponse creates a response to parse from DescribeDynamicTagRuleList response
func CreateDescribeDynamicTagRuleListResponse() (response *DescribeDynamicTagRuleListResponse) {
	response = &DescribeDynamicTagRuleListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
