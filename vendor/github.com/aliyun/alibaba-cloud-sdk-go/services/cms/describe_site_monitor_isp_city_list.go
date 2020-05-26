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

// DescribeSiteMonitorISPCityList invokes the cms.DescribeSiteMonitorISPCityList API synchronously
// api document: https://help.aliyun.com/api/cms/describesitemonitorispcitylist.html
func (client *Client) DescribeSiteMonitorISPCityList(request *DescribeSiteMonitorISPCityListRequest) (response *DescribeSiteMonitorISPCityListResponse, err error) {
	response = CreateDescribeSiteMonitorISPCityListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSiteMonitorISPCityListWithChan invokes the cms.DescribeSiteMonitorISPCityList API asynchronously
// api document: https://help.aliyun.com/api/cms/describesitemonitorispcitylist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSiteMonitorISPCityListWithChan(request *DescribeSiteMonitorISPCityListRequest) (<-chan *DescribeSiteMonitorISPCityListResponse, <-chan error) {
	responseChan := make(chan *DescribeSiteMonitorISPCityListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSiteMonitorISPCityList(request)
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

// DescribeSiteMonitorISPCityListWithCallback invokes the cms.DescribeSiteMonitorISPCityList API asynchronously
// api document: https://help.aliyun.com/api/cms/describesitemonitorispcitylist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSiteMonitorISPCityListWithCallback(request *DescribeSiteMonitorISPCityListRequest, callback func(response *DescribeSiteMonitorISPCityListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSiteMonitorISPCityListResponse
		var err error
		defer close(result)
		response, err = client.DescribeSiteMonitorISPCityList(request)
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

// DescribeSiteMonitorISPCityListRequest is the request struct for api DescribeSiteMonitorISPCityList
type DescribeSiteMonitorISPCityListRequest struct {
	*requests.RpcRequest
	City string `position:"Query" name:"City"`
	Isp  string `position:"Query" name:"Isp"`
}

// DescribeSiteMonitorISPCityListResponse is the response struct for api DescribeSiteMonitorISPCityList
type DescribeSiteMonitorISPCityListResponse struct {
	*responses.BaseResponse
	Code        string      `json:"Code" xml:"Code"`
	Message     string      `json:"Message" xml:"Message"`
	Success     string      `json:"Success" xml:"Success"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	IspCityList IspCityList `json:"IspCityList" xml:"IspCityList"`
}

// CreateDescribeSiteMonitorISPCityListRequest creates a request to invoke DescribeSiteMonitorISPCityList API
func CreateDescribeSiteMonitorISPCityListRequest() (request *DescribeSiteMonitorISPCityListRequest) {
	request = &DescribeSiteMonitorISPCityListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeSiteMonitorISPCityList", "cms", "openAPI")
	return
}

// CreateDescribeSiteMonitorISPCityListResponse creates a response to parse from DescribeSiteMonitorISPCityList response
func CreateDescribeSiteMonitorISPCityListResponse() (response *DescribeSiteMonitorISPCityListResponse) {
	response = &DescribeSiteMonitorISPCityListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
