package rds

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

// DescribeAvailableClasses invokes the rds.DescribeAvailableClasses API synchronously
// api document: https://help.aliyun.com/api/rds/describeavailableclasses.html
func (client *Client) DescribeAvailableClasses(request *DescribeAvailableClassesRequest) (response *DescribeAvailableClassesResponse, err error) {
	response = CreateDescribeAvailableClassesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAvailableClassesWithChan invokes the rds.DescribeAvailableClasses API asynchronously
// api document: https://help.aliyun.com/api/rds/describeavailableclasses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAvailableClassesWithChan(request *DescribeAvailableClassesRequest) (<-chan *DescribeAvailableClassesResponse, <-chan error) {
	responseChan := make(chan *DescribeAvailableClassesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAvailableClasses(request)
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

// DescribeAvailableClassesWithCallback invokes the rds.DescribeAvailableClasses API asynchronously
// api document: https://help.aliyun.com/api/rds/describeavailableclasses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAvailableClassesWithCallback(request *DescribeAvailableClassesRequest, callback func(response *DescribeAvailableClassesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAvailableClassesResponse
		var err error
		defer close(result)
		response, err = client.DescribeAvailableClasses(request)
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

// DescribeAvailableClassesRequest is the request struct for api DescribeAvailableClasses
type DescribeAvailableClassesRequest struct {
	*requests.RpcRequest
	DBInstanceName        string           `position:"Query" name:"DBInstanceName"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EngineVersion         string           `position:"Query" name:"EngineVersion"`
	Engine                string           `position:"Query" name:"Engine"`
	DBInstanceId          string           `position:"Query" name:"DBInstanceId"`
	DBInstanceStorageType string           `position:"Query" name:"DBInstanceStorageType"`
	InstanceChargeType    string           `position:"Query" name:"InstanceChargeType"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	CommodityCode         string           `position:"Query" name:"CommodityCode"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	ZoneId                string           `position:"Query" name:"ZoneId"`
	Category              string           `position:"Query" name:"Category"`
	OrderType             string           `position:"Query" name:"OrderType"`
}

// DescribeAvailableClassesResponse is the response struct for api DescribeAvailableClasses
type DescribeAvailableClassesResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	DBInstanceClasses []DBInstanceClass `json:"DBInstanceClasses" xml:"DBInstanceClasses"`
}

// CreateDescribeAvailableClassesRequest creates a request to invoke DescribeAvailableClasses API
func CreateDescribeAvailableClassesRequest() (request *DescribeAvailableClassesRequest) {
	request = &DescribeAvailableClassesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeAvailableClasses", "rds", "openAPI")
	return
}

// CreateDescribeAvailableClassesResponse creates a response to parse from DescribeAvailableClasses response
func CreateDescribeAvailableClassesResponse() (response *DescribeAvailableClassesResponse) {
	response = &DescribeAvailableClassesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
