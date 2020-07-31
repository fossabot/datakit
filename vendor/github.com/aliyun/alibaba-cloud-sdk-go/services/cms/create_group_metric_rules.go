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

// CreateGroupMetricRules invokes the cms.CreateGroupMetricRules API synchronously
// api document: https://help.aliyun.com/api/cms/creategroupmetricrules.html
func (client *Client) CreateGroupMetricRules(request *CreateGroupMetricRulesRequest) (response *CreateGroupMetricRulesResponse, err error) {
	response = CreateCreateGroupMetricRulesResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGroupMetricRulesWithChan invokes the cms.CreateGroupMetricRules API asynchronously
// api document: https://help.aliyun.com/api/cms/creategroupmetricrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGroupMetricRulesWithChan(request *CreateGroupMetricRulesRequest) (<-chan *CreateGroupMetricRulesResponse, <-chan error) {
	responseChan := make(chan *CreateGroupMetricRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGroupMetricRules(request)
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

// CreateGroupMetricRulesWithCallback invokes the cms.CreateGroupMetricRules API asynchronously
// api document: https://help.aliyun.com/api/cms/creategroupmetricrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGroupMetricRulesWithCallback(request *CreateGroupMetricRulesRequest, callback func(response *CreateGroupMetricRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGroupMetricRulesResponse
		var err error
		defer close(result)
		response, err = client.CreateGroupMetricRules(request)
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

// CreateGroupMetricRulesRequest is the request struct for api CreateGroupMetricRules
type CreateGroupMetricRulesRequest struct {
	*requests.RpcRequest
	GroupId          requests.Integer                          `position:"Query" name:"GroupId"`
	GroupMetricRules *[]CreateGroupMetricRulesGroupMetricRules `position:"Query" name:"GroupMetricRules"  type:"Repeated"`
}

// CreateGroupMetricRulesGroupMetricRules is a repeated param struct in CreateGroupMetricRulesRequest
type CreateGroupMetricRulesGroupMetricRules struct {
	Webhook                               string `name:"Webhook"`
	EscalationsWarnComparisonOperator     string `name:"Escalations.Warn.ComparisonOperator"`
	RuleName                              string `name:"RuleName"`
	EscalationsInfoStatistics             string `name:"Escalations.Info.Statistics"`
	EffectiveInterval                     string `name:"EffectiveInterval"`
	EscalationsInfoComparisonOperator     string `name:"Escalations.Info.ComparisonOperator"`
	NoEffectiveInterval                   string `name:"NoEffectiveInterval"`
	EmailSubject                          string `name:"EmailSubject"`
	SilenceTime                           string `name:"SilenceTime"`
	MetricName                            string `name:"MetricName"`
	EscalationsWarnTimes                  string `name:"Escalations.Warn.Times"`
	Period                                string `name:"Period"`
	EscalationsWarnThreshold              string `name:"Escalations.Warn.Threshold"`
	EscalationsCriticalStatistics         string `name:"Escalations.Critical.Statistics"`
	EscalationsInfoTimes                  string `name:"Escalations.Info.Times"`
	EscalationsCriticalTimes              string `name:"Escalations.Critical.Times"`
	EscalationsWarnStatistics             string `name:"Escalations.Warn.Statistics"`
	EscalationsInfoThreshold              string `name:"Escalations.Info.Threshold"`
	Namespace                             string `name:"Namespace"`
	Interval                              string `name:"Interval"`
	Category                              string `name:"Category"`
	RuleId                                string `name:"RuleId"`
	EscalationsCriticalComparisonOperator string `name:"Escalations.Critical.ComparisonOperator"`
	EscalationsCriticalThreshold          string `name:"Escalations.Critical.Threshold"`
	Dimensions                            string `name:"Dimensions"`
}

// CreateGroupMetricRulesResponse is the response struct for api CreateGroupMetricRules
type CreateGroupMetricRulesResponse struct {
	*responses.BaseResponse
	RequestId string                            `json:"RequestId" xml:"RequestId"`
	Success   bool                              `json:"Success" xml:"Success"`
	Code      int                               `json:"Code" xml:"Code"`
	Message   string                            `json:"Message" xml:"Message"`
	Resources ResourcesInCreateGroupMetricRules `json:"Resources" xml:"Resources"`
}

// CreateCreateGroupMetricRulesRequest creates a request to invoke CreateGroupMetricRules API
func CreateCreateGroupMetricRulesRequest() (request *CreateGroupMetricRulesRequest) {
	request = &CreateGroupMetricRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "CreateGroupMetricRules", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateGroupMetricRulesResponse creates a response to parse from CreateGroupMetricRules response
func CreateCreateGroupMetricRulesResponse() (response *CreateGroupMetricRulesResponse) {
	response = &CreateGroupMetricRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
