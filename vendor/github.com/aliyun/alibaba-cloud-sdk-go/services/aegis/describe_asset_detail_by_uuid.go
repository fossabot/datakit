package aegis

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

// DescribeAssetDetailByUuid invokes the aegis.DescribeAssetDetailByUuid API synchronously
// api document: https://help.aliyun.com/api/aegis/describeassetdetailbyuuid.html
func (client *Client) DescribeAssetDetailByUuid(request *DescribeAssetDetailByUuidRequest) (response *DescribeAssetDetailByUuidResponse, err error) {
	response = CreateDescribeAssetDetailByUuidResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAssetDetailByUuidWithChan invokes the aegis.DescribeAssetDetailByUuid API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeassetdetailbyuuid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAssetDetailByUuidWithChan(request *DescribeAssetDetailByUuidRequest) (<-chan *DescribeAssetDetailByUuidResponse, <-chan error) {
	responseChan := make(chan *DescribeAssetDetailByUuidResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAssetDetailByUuid(request)
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

// DescribeAssetDetailByUuidWithCallback invokes the aegis.DescribeAssetDetailByUuid API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeassetdetailbyuuid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAssetDetailByUuidWithCallback(request *DescribeAssetDetailByUuidRequest, callback func(response *DescribeAssetDetailByUuidResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAssetDetailByUuidResponse
		var err error
		defer close(result)
		response, err = client.DescribeAssetDetailByUuid(request)
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

// DescribeAssetDetailByUuidRequest is the request struct for api DescribeAssetDetailByUuid
type DescribeAssetDetailByUuidRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
	Uuid     string `position:"Query" name:"Uuid"`
}

// DescribeAssetDetailByUuidResponse is the response struct for api DescribeAssetDetailByUuid
type DescribeAssetDetailByUuidResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	AssetDetail AssetDetail `json:"AssetDetail" xml:"AssetDetail"`
}

// CreateDescribeAssetDetailByUuidRequest creates a request to invoke DescribeAssetDetailByUuid API
func CreateDescribeAssetDetailByUuidRequest() (request *DescribeAssetDetailByUuidRequest) {
	request = &DescribeAssetDetailByUuidRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeAssetDetailByUuid", "vipaegis", "openAPI")
	return
}

// CreateDescribeAssetDetailByUuidResponse creates a response to parse from DescribeAssetDetailByUuid response
func CreateDescribeAssetDetailByUuidResponse() (response *DescribeAssetDetailByUuidResponse) {
	response = &DescribeAssetDetailByUuidResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
