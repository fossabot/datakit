package bssopenapi

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

// QueryEvaluateList invokes the bssopenapi.QueryEvaluateList API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryevaluatelist.html
func (client *Client) QueryEvaluateList(request *QueryEvaluateListRequest) (response *QueryEvaluateListResponse, err error) {
	response = CreateQueryEvaluateListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEvaluateListWithChan invokes the bssopenapi.QueryEvaluateList API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryevaluatelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryEvaluateListWithChan(request *QueryEvaluateListRequest) (<-chan *QueryEvaluateListResponse, <-chan error) {
	responseChan := make(chan *QueryEvaluateListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEvaluateList(request)
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

// QueryEvaluateListWithCallback invokes the bssopenapi.QueryEvaluateList API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryevaluatelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryEvaluateListWithCallback(request *QueryEvaluateListRequest, callback func(response *QueryEvaluateListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEvaluateListResponse
		var err error
		defer close(result)
		response, err = client.QueryEvaluateList(request)
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

// QueryEvaluateListRequest is the request struct for api QueryEvaluateList
type QueryEvaluateListRequest struct {
	*requests.RpcRequest
	EndSearchTime   string           `position:"Query" name:"EndSearchTime"`
	OutBizId        string           `position:"Query" name:"OutBizId"`
	SortType        requests.Integer `position:"Query" name:"SortType"`
	Type            requests.Integer `position:"Query" name:"Type"`
	PageNum         requests.Integer `position:"Query" name:"PageNum"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	EndAmount       requests.Integer `position:"Query" name:"EndAmount"`
	BillCycle       string           `position:"Query" name:"BillCycle"`
	BizTypeList     *[]string        `position:"Query" name:"BizTypeList"  type:"Repeated"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	StartSearchTime string           `position:"Query" name:"StartSearchTime"`
	EndBizTime      string           `position:"Query" name:"EndBizTime"`
	StartAmount     requests.Integer `position:"Query" name:"StartAmount"`
	StartBizTime    string           `position:"Query" name:"StartBizTime"`
}

// QueryEvaluateListResponse is the response struct for api QueryEvaluateList
type QueryEvaluateListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryEvaluateListRequest creates a request to invoke QueryEvaluateList API
func CreateQueryEvaluateListRequest() (request *QueryEvaluateListRequest) {
	request = &QueryEvaluateListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryEvaluateList", "", "")
	return
}

// CreateQueryEvaluateListResponse creates a response to parse from QueryEvaluateList response
func CreateQueryEvaluateListResponse() (response *QueryEvaluateListResponse) {
	response = &QueryEvaluateListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
