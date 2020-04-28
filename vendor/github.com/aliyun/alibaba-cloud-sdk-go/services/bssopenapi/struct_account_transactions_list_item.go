package bssopenapi

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

// AccountTransactionsListItem is a nested struct in bssopenapi response
type AccountTransactionsListItem struct {
	TransactionNumber    string `json:"TransactionNumber" xml:"TransactionNumber"`
	TransactionTime      string `json:"TransactionTime" xml:"TransactionTime"`
	TransactionFlow      string `json:"TransactionFlow" xml:"TransactionFlow"`
	TransactionType      string `json:"TransactionType" xml:"TransactionType"`
	TransactionChannel   string `json:"TransactionChannel" xml:"TransactionChannel"`
	TransactionChannelSN string `json:"TransactionChannelSN" xml:"TransactionChannelSN"`
	FundType             string `json:"FundType" xml:"FundType"`
	RecordID             string `json:"RecordID" xml:"RecordID"`
	Remarks              string `json:"Remarks" xml:"Remarks"`
	BillingCycle         string `json:"BillingCycle" xml:"BillingCycle"`
	Amount               string `json:"Amount" xml:"Amount"`
	Balance              string `json:"Balance" xml:"Balance"`
	TransactionAccount   string `json:"TransactionAccount" xml:"TransactionAccount"`
}
