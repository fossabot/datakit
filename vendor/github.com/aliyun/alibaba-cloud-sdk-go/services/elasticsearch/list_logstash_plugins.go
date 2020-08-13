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

// ListLogstashPlugins invokes the elasticsearch.ListLogstashPlugins API synchronously
// api document: https://help.aliyun.com/api/elasticsearch/listlogstashplugins.html
func (client *Client) ListLogstashPlugins(request *ListLogstashPluginsRequest) (response *ListLogstashPluginsResponse, err error) {
	response = CreateListLogstashPluginsResponse()
	err = client.DoAction(request, response)
	return
}

// ListLogstashPluginsWithChan invokes the elasticsearch.ListLogstashPlugins API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/listlogstashplugins.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListLogstashPluginsWithChan(request *ListLogstashPluginsRequest) (<-chan *ListLogstashPluginsResponse, <-chan error) {
	responseChan := make(chan *ListLogstashPluginsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLogstashPlugins(request)
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

// ListLogstashPluginsWithCallback invokes the elasticsearch.ListLogstashPlugins API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/listlogstashplugins.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListLogstashPluginsWithCallback(request *ListLogstashPluginsRequest, callback func(response *ListLogstashPluginsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLogstashPluginsResponse
		var err error
		defer close(result)
		response, err = client.ListLogstashPlugins(request)
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

// ListLogstashPluginsRequest is the request struct for api ListLogstashPlugins
type ListLogstashPluginsRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Path" name:"InstanceId"`
	Size       requests.Integer `position:"Query" name:"size"`
	Name       string           `position:"Query" name:"name"`
	Page       string           `position:"Query" name:"page"`
	Source     string           `position:"Query" name:"source"`
}

// ListLogstashPluginsResponse is the response struct for api ListLogstashPlugins
type ListLogstashPluginsResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateListLogstashPluginsRequest creates a request to invoke ListLogstashPlugins API
func CreateListLogstashPluginsRequest() (request *ListLogstashPluginsRequest) {
	request = &ListLogstashPluginsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ListLogstashPlugins", "/openapi/logstashes/[InstanceId]/plugins", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListLogstashPluginsResponse creates a response to parse from ListLogstashPlugins response
func CreateListLogstashPluginsResponse() (response *ListLogstashPluginsResponse) {
	response = &ListLogstashPluginsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
