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

// DescribeMetricData invokes the cms.DescribeMetricData API synchronously
// api document: https://help.aliyun.com/api/cms/describemetricdata.html
func (client *Client) DescribeMetricData(request *DescribeMetricDataRequest) (response *DescribeMetricDataResponse, err error) {
	response = CreateDescribeMetricDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMetricDataWithChan invokes the cms.DescribeMetricData API asynchronously
// api document: https://help.aliyun.com/api/cms/describemetricdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMetricDataWithChan(request *DescribeMetricDataRequest) (<-chan *DescribeMetricDataResponse, <-chan error) {
	responseChan := make(chan *DescribeMetricDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMetricData(request)
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

// DescribeMetricDataWithCallback invokes the cms.DescribeMetricData API asynchronously
// api document: https://help.aliyun.com/api/cms/describemetricdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMetricDataWithCallback(request *DescribeMetricDataRequest, callback func(response *DescribeMetricDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMetricDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeMetricData(request)
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

// DescribeMetricDataRequest is the request struct for api DescribeMetricData
type DescribeMetricDataRequest struct {
	*requests.RpcRequest
	Period     string `position:"Query" name:"Period"`
	Length     string `position:"Query" name:"Length"`
	EndTime    string `position:"Query" name:"EndTime"`
	Express    string `position:"Query" name:"Express"`
	StartTime  string `position:"Query" name:"StartTime"`
	Namespace  string `position:"Query" name:"Namespace"`
	MetricName string `position:"Query" name:"MetricName"`
	Dimensions string `position:"Query" name:"Dimensions"`
}

// DescribeMetricDataResponse is the response struct for api DescribeMetricData
type DescribeMetricDataResponse struct {
	*responses.BaseResponse
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Datapoints string `json:"Datapoints" xml:"Datapoints"`
	Period     string `json:"Period" xml:"Period"`
}

// CreateDescribeMetricDataRequest creates a request to invoke DescribeMetricData API
func CreateDescribeMetricDataRequest() (request *DescribeMetricDataRequest) {
	request = &DescribeMetricDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeMetricData", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMetricDataResponse creates a response to parse from DescribeMetricData response
func CreateDescribeMetricDataResponse() (response *DescribeMetricDataResponse) {
	response = &DescribeMetricDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
