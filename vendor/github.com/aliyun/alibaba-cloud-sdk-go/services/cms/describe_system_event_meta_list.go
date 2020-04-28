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

// DescribeSystemEventMetaList invokes the cms.DescribeSystemEventMetaList API synchronously
// api document: https://help.aliyun.com/api/cms/describesystemeventmetalist.html
func (client *Client) DescribeSystemEventMetaList(request *DescribeSystemEventMetaListRequest) (response *DescribeSystemEventMetaListResponse, err error) {
	response = CreateDescribeSystemEventMetaListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSystemEventMetaListWithChan invokes the cms.DescribeSystemEventMetaList API asynchronously
// api document: https://help.aliyun.com/api/cms/describesystemeventmetalist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSystemEventMetaListWithChan(request *DescribeSystemEventMetaListRequest) (<-chan *DescribeSystemEventMetaListResponse, <-chan error) {
	responseChan := make(chan *DescribeSystemEventMetaListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSystemEventMetaList(request)
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

// DescribeSystemEventMetaListWithCallback invokes the cms.DescribeSystemEventMetaList API asynchronously
// api document: https://help.aliyun.com/api/cms/describesystemeventmetalist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSystemEventMetaListWithCallback(request *DescribeSystemEventMetaListRequest, callback func(response *DescribeSystemEventMetaListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSystemEventMetaListResponse
		var err error
		defer close(result)
		response, err = client.DescribeSystemEventMetaList(request)
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

// DescribeSystemEventMetaListRequest is the request struct for api DescribeSystemEventMetaList
type DescribeSystemEventMetaListRequest struct {
	*requests.RpcRequest
}

// DescribeSystemEventMetaListResponse is the response struct for api DescribeSystemEventMetaList
type DescribeSystemEventMetaListResponse struct {
	*responses.BaseResponse
	RequestId string                            `json:"RequestId" xml:"RequestId"`
	Success   bool                              `json:"Success" xml:"Success"`
	Code      int                               `json:"Code" xml:"Code"`
	Message   string                            `json:"Message" xml:"Message"`
	Data      DataInDescribeSystemEventMetaList `json:"Data" xml:"Data"`
}

// CreateDescribeSystemEventMetaListRequest creates a request to invoke DescribeSystemEventMetaList API
func CreateDescribeSystemEventMetaListRequest() (request *DescribeSystemEventMetaListRequest) {
	request = &DescribeSystemEventMetaListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeSystemEventMetaList", "cms", "openAPI")
	return
}

// CreateDescribeSystemEventMetaListResponse creates a response to parse from DescribeSystemEventMetaList response
func CreateDescribeSystemEventMetaListResponse() (response *DescribeSystemEventMetaListResponse) {
	response = &DescribeSystemEventMetaListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
