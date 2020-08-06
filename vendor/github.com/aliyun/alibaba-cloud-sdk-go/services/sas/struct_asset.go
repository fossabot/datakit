package sas

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

// Asset is a nested struct in sas response
type Asset struct {
	InternetIp    string `json:"InternetIp" xml:"InternetIp"`
	IntranetIp    string `json:"IntranetIp" xml:"IntranetIp"`
	InstanceName  int    `json:"InstanceName" xml:"InstanceName"`
	InstanceId    int    `json:"InstanceId" xml:"InstanceId"`
	Ip            string `json:"Ip" xml:"Ip"`
	Uuid          string `json:"Uuid" xml:"Uuid"`
	AssetType     string `json:"AssetType" xml:"AssetType"`
	Os            string `json:"Os" xml:"Os"`
	ClientStatus  string `json:"ClientStatus" xml:"ClientStatus"`
	Region        string `json:"Region" xml:"Region"`
	RegionName    string `json:"RegionName" xml:"RegionName"`
	VpcInstanceId string `json:"VpcInstanceId" xml:"VpcInstanceId"`
	OsName        string `json:"OsName" xml:"OsName"`
}
