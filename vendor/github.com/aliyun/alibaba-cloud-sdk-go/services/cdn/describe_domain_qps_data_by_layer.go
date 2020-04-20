package cdn

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

// DescribeDomainQpsDataByLayer invokes the cdn.DescribeDomainQpsDataByLayer API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainqpsdatabylayer.html
func (client *Client) DescribeDomainQpsDataByLayer(request *DescribeDomainQpsDataByLayerRequest) (response *DescribeDomainQpsDataByLayerResponse, err error) {
	response = CreateDescribeDomainQpsDataByLayerResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainQpsDataByLayerWithChan invokes the cdn.DescribeDomainQpsDataByLayer API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainqpsdatabylayer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainQpsDataByLayerWithChan(request *DescribeDomainQpsDataByLayerRequest) (<-chan *DescribeDomainQpsDataByLayerResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainQpsDataByLayerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainQpsDataByLayer(request)
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

// DescribeDomainQpsDataByLayerWithCallback invokes the cdn.DescribeDomainQpsDataByLayer API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainqpsdatabylayer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainQpsDataByLayerWithCallback(request *DescribeDomainQpsDataByLayerRequest, callback func(response *DescribeDomainQpsDataByLayerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainQpsDataByLayerResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainQpsDataByLayer(request)
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

// DescribeDomainQpsDataByLayerRequest is the request struct for api DescribeDomainQpsDataByLayer
type DescribeDomainQpsDataByLayerRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	Layer          string           `position:"Query" name:"Layer"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Interval       string           `position:"Query" name:"Interval"`
}

// DescribeDomainQpsDataByLayerResponse is the response struct for api DescribeDomainQpsDataByLayer
type DescribeDomainQpsDataByLayerResponse struct {
	*responses.BaseResponse
	RequestId       string                                        `json:"RequestId" xml:"RequestId"`
	DomainName      string                                        `json:"DomainName" xml:"DomainName"`
	StartTime       string                                        `json:"StartTime" xml:"StartTime"`
	EndTime         string                                        `json:"EndTime" xml:"EndTime"`
	DataInterval    string                                        `json:"DataInterval" xml:"DataInterval"`
	Layer           string                                        `json:"Layer" xml:"Layer"`
	QpsDataInterval QpsDataIntervalInDescribeDomainQpsDataByLayer `json:"QpsDataInterval" xml:"QpsDataInterval"`
}

// CreateDescribeDomainQpsDataByLayerRequest creates a request to invoke DescribeDomainQpsDataByLayer API
func CreateDescribeDomainQpsDataByLayerRequest() (request *DescribeDomainQpsDataByLayerRequest) {
	request = &DescribeDomainQpsDataByLayerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainQpsDataByLayer", "", "")
	return
}

// CreateDescribeDomainQpsDataByLayerResponse creates a response to parse from DescribeDomainQpsDataByLayer response
func CreateDescribeDomainQpsDataByLayerResponse() (response *DescribeDomainQpsDataByLayerResponse) {
	response = &DescribeDomainQpsDataByLayerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
