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

// DescribeSQLCollectorRetention invokes the rds.DescribeSQLCollectorRetention API synchronously
// api document: https://help.aliyun.com/api/rds/describesqlcollectorretention.html
func (client *Client) DescribeSQLCollectorRetention(request *DescribeSQLCollectorRetentionRequest) (response *DescribeSQLCollectorRetentionResponse, err error) {
	response = CreateDescribeSQLCollectorRetentionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSQLCollectorRetentionWithChan invokes the rds.DescribeSQLCollectorRetention API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqlcollectorretention.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSQLCollectorRetentionWithChan(request *DescribeSQLCollectorRetentionRequest) (<-chan *DescribeSQLCollectorRetentionResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLCollectorRetentionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLCollectorRetention(request)
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

// DescribeSQLCollectorRetentionWithCallback invokes the rds.DescribeSQLCollectorRetention API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqlcollectorretention.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSQLCollectorRetentionWithCallback(request *DescribeSQLCollectorRetentionRequest, callback func(response *DescribeSQLCollectorRetentionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLCollectorRetentionResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLCollectorRetention(request)
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

// DescribeSQLCollectorRetentionRequest is the request struct for api DescribeSQLCollectorRetention
type DescribeSQLCollectorRetentionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeSQLCollectorRetentionResponse is the response struct for api DescribeSQLCollectorRetention
type DescribeSQLCollectorRetentionResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	ConfigValue string `json:"ConfigValue" xml:"ConfigValue"`
}

// CreateDescribeSQLCollectorRetentionRequest creates a request to invoke DescribeSQLCollectorRetention API
func CreateDescribeSQLCollectorRetentionRequest() (request *DescribeSQLCollectorRetentionRequest) {
	request = &DescribeSQLCollectorRetentionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLCollectorRetention", "rds", "openAPI")
	return
}

// CreateDescribeSQLCollectorRetentionResponse creates a response to parse from DescribeSQLCollectorRetention response
func CreateDescribeSQLCollectorRetentionResponse() (response *DescribeSQLCollectorRetentionResponse) {
	response = &DescribeSQLCollectorRetentionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
