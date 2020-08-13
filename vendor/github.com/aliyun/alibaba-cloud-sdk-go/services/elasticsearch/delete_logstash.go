package elasticsearch

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

// DeleteLogstash invokes the elasticsearch.DeleteLogstash API synchronously
// api document: https://help.aliyun.com/api/elasticsearch/deletelogstash.html
func (client *Client) DeleteLogstash(request *DeleteLogstashRequest) (response *DeleteLogstashResponse, err error) {
	response = CreateDeleteLogstashResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLogstashWithChan invokes the elasticsearch.DeleteLogstash API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/deletelogstash.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLogstashWithChan(request *DeleteLogstashRequest) (<-chan *DeleteLogstashResponse, <-chan error) {
	responseChan := make(chan *DeleteLogstashResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLogstash(request)
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

// DeleteLogstashWithCallback invokes the elasticsearch.DeleteLogstash API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/deletelogstash.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLogstashWithCallback(request *DeleteLogstashRequest, callback func(response *DeleteLogstashResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLogstashResponse
		var err error
		defer close(result)
		response, err = client.DeleteLogstash(request)
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

// DeleteLogstashRequest is the request struct for api DeleteLogstash
type DeleteLogstashRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
}

// DeleteLogstashResponse is the response struct for api DeleteLogstash
type DeleteLogstashResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLogstashRequest creates a request to invoke DeleteLogstash API
func CreateDeleteLogstashRequest() (request *DeleteLogstashRequest) {
	request = &DeleteLogstashRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "DeleteLogstash", "/openapi/logstashes/[InstanceId]", "elasticsearch", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteLogstashResponse creates a response to parse from DeleteLogstash response
func CreateDeleteLogstashResponse() (response *DeleteLogstashResponse) {
	response = &DeleteLogstashResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
