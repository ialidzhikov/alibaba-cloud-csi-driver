package dm

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

// Task is a nested struct in dm response
type Task struct {
	TaskId        string `json:"TaskId" xml:"TaskId"`
	ReceiversName string `json:"ReceiversName" xml:"ReceiversName"`
	TemplateName  string `json:"TemplateName" xml:"TemplateName"`
	TaskStatus    string `json:"TaskStatus" xml:"TaskStatus"`
	CreateTime    string `json:"CreateTime" xml:"CreateTime"`
	UtcCreateTime int    `json:"UtcCreateTime" xml:"UtcCreateTime"`
	AddressType   string `json:"AddressType" xml:"AddressType"`
	TagName       string `json:"TagName" xml:"TagName"`
	RequestCount  string `json:"RequestCount" xml:"RequestCount"`
}
