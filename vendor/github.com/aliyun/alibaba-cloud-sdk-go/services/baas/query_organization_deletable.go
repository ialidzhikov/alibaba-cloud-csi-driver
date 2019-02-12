package baas

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

// QueryOrganizationDeletable invokes the baas.QueryOrganizationDeletable API synchronously
// api document: https://help.aliyun.com/api/baas/queryorganizationdeletable.html
func (client *Client) QueryOrganizationDeletable(request *QueryOrganizationDeletableRequest) (response *QueryOrganizationDeletableResponse, err error) {
	response = CreateQueryOrganizationDeletableResponse()
	err = client.DoAction(request, response)
	return
}

// QueryOrganizationDeletableWithChan invokes the baas.QueryOrganizationDeletable API asynchronously
// api document: https://help.aliyun.com/api/baas/queryorganizationdeletable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryOrganizationDeletableWithChan(request *QueryOrganizationDeletableRequest) (<-chan *QueryOrganizationDeletableResponse, <-chan error) {
	responseChan := make(chan *QueryOrganizationDeletableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryOrganizationDeletable(request)
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

// QueryOrganizationDeletableWithCallback invokes the baas.QueryOrganizationDeletable API asynchronously
// api document: https://help.aliyun.com/api/baas/queryorganizationdeletable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryOrganizationDeletableWithCallback(request *QueryOrganizationDeletableRequest, callback func(response *QueryOrganizationDeletableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryOrganizationDeletableResponse
		var err error
		defer close(result)
		response, err = client.QueryOrganizationDeletable(request)
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

// QueryOrganizationDeletableRequest is the request struct for api QueryOrganizationDeletable
type QueryOrganizationDeletableRequest struct {
	*requests.RpcRequest
	OrganizationId string `position:"Query" name:"OrganizationId"`
	Location       string `position:"Body" name:"Location"`
}

// QueryOrganizationDeletableResponse is the response struct for api QueryOrganizationDeletable
type QueryOrganizationDeletableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateQueryOrganizationDeletableRequest creates a request to invoke QueryOrganizationDeletable API
func CreateQueryOrganizationDeletableRequest() (request *QueryOrganizationDeletableRequest) {
	request = &QueryOrganizationDeletableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-07-31", "QueryOrganizationDeletable", "", "")
	return
}

// CreateQueryOrganizationDeletableResponse creates a response to parse from QueryOrganizationDeletable response
func CreateQueryOrganizationDeletableResponse() (response *QueryOrganizationDeletableResponse) {
	response = &QueryOrganizationDeletableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
