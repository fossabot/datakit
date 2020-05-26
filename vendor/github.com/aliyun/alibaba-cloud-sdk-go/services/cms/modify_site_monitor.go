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

// ModifySiteMonitor invokes the cms.ModifySiteMonitor API synchronously
// api document: https://help.aliyun.com/api/cms/modifysitemonitor.html
func (client *Client) ModifySiteMonitor(request *ModifySiteMonitorRequest) (response *ModifySiteMonitorResponse, err error) {
	response = CreateModifySiteMonitorResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySiteMonitorWithChan invokes the cms.ModifySiteMonitor API asynchronously
// api document: https://help.aliyun.com/api/cms/modifysitemonitor.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySiteMonitorWithChan(request *ModifySiteMonitorRequest) (<-chan *ModifySiteMonitorResponse, <-chan error) {
	responseChan := make(chan *ModifySiteMonitorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySiteMonitor(request)
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

// ModifySiteMonitorWithCallback invokes the cms.ModifySiteMonitor API asynchronously
// api document: https://help.aliyun.com/api/cms/modifysitemonitor.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySiteMonitorWithCallback(request *ModifySiteMonitorRequest, callback func(response *ModifySiteMonitorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySiteMonitorResponse
		var err error
		defer close(result)
		response, err = client.ModifySiteMonitor(request)
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

// ModifySiteMonitorRequest is the request struct for api ModifySiteMonitor
type ModifySiteMonitorRequest struct {
	*requests.RpcRequest
	Address     string `position:"Query" name:"Address"`
	TaskName    string `position:"Query" name:"TaskName"`
	IspCities   string `position:"Query" name:"IspCities"`
	OptionsJson string `position:"Query" name:"OptionsJson"`
	AlertIds    string `position:"Query" name:"AlertIds"`
	Interval    string `position:"Query" name:"Interval"`
	TaskId      string `position:"Query" name:"TaskId"`
}

// ModifySiteMonitorResponse is the response struct for api ModifySiteMonitor
type ModifySiteMonitorResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   string `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateModifySiteMonitorRequest creates a request to invoke ModifySiteMonitor API
func CreateModifySiteMonitorRequest() (request *ModifySiteMonitorRequest) {
	request = &ModifySiteMonitorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "ModifySiteMonitor", "cms", "openAPI")
	return
}

// CreateModifySiteMonitorResponse creates a response to parse from ModifySiteMonitor response
func CreateModifySiteMonitorResponse() (response *ModifySiteMonitorResponse) {
	response = &ModifySiteMonitorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
