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

// QueryPrepaidCards invokes the bssopenapi.QueryPrepaidCards API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryprepaidcards.html
func (client *Client) QueryPrepaidCards(request *QueryPrepaidCardsRequest) (response *QueryPrepaidCardsResponse, err error) {
	response = CreateQueryPrepaidCardsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryPrepaidCardsWithChan invokes the bssopenapi.QueryPrepaidCards API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryprepaidcards.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPrepaidCardsWithChan(request *QueryPrepaidCardsRequest) (<-chan *QueryPrepaidCardsResponse, <-chan error) {
	responseChan := make(chan *QueryPrepaidCardsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPrepaidCards(request)
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

// QueryPrepaidCardsWithCallback invokes the bssopenapi.QueryPrepaidCards API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryprepaidcards.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPrepaidCardsWithCallback(request *QueryPrepaidCardsRequest, callback func(response *QueryPrepaidCardsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPrepaidCardsResponse
		var err error
		defer close(result)
		response, err = client.QueryPrepaidCards(request)
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

// QueryPrepaidCardsRequest is the request struct for api QueryPrepaidCards
type QueryPrepaidCardsRequest struct {
	*requests.RpcRequest
	ExpiryTimeEnd   string           `position:"Query" name:"ExpiryTimeEnd"`
	ExpiryTimeStart string           `position:"Query" name:"ExpiryTimeStart"`
	EffectiveOrNot  requests.Boolean `position:"Query" name:"EffectiveOrNot"`
}

// QueryPrepaidCardsResponse is the response struct for api QueryPrepaidCards
type QueryPrepaidCardsResponse struct {
	*responses.BaseResponse
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	Success   bool                    `json:"Success" xml:"Success"`
	Code      string                  `json:"Code" xml:"Code"`
	Message   string                  `json:"Message" xml:"Message"`
	Data      DataInQueryPrepaidCards `json:"Data" xml:"Data"`
}

// CreateQueryPrepaidCardsRequest creates a request to invoke QueryPrepaidCards API
func CreateQueryPrepaidCardsRequest() (request *QueryPrepaidCardsRequest) {
	request = &QueryPrepaidCardsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryPrepaidCards", "", "")
	return
}

// CreateQueryPrepaidCardsResponse creates a response to parse from QueryPrepaidCards response
func CreateQueryPrepaidCardsResponse() (response *QueryPrepaidCardsResponse) {
	response = &QueryPrepaidCardsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
