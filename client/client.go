// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetKpmEncryptedNodeTuplesByOrderIdRequest struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s GetKpmEncryptedNodeTuplesByOrderIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetKpmEncryptedNodeTuplesByOrderIdRequest) GoString() string {
	return s.String()
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdRequest) SetOrderId(v int64) *GetKpmEncryptedNodeTuplesByOrderIdRequest {
	s.OrderId = &v
	return s
}

type GetKpmEncryptedNodeTuplesByOrderIdResponseBody struct {
	Code                *string                                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode         *string                                                            `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage      *string                                                            `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	EncryptedNodeTuples *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuples `json:"EncryptedNodeTuples,omitempty" xml:"EncryptedNodeTuples,omitempty" type:"Struct"`
	EncryptedSessionZmk *string                                                            `json:"EncryptedSessionZmk,omitempty" xml:"EncryptedSessionZmk,omitempty"`
	Message             *string                                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId           *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success             *bool                                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKpmEncryptedNodeTuplesByOrderIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetKpmEncryptedNodeTuplesByOrderIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBody) SetCode(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBody) SetDynamicCode(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBody) SetDynamicMessage(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBody) SetEncryptedNodeTuples(v *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuples) *GetKpmEncryptedNodeTuplesByOrderIdResponseBody {
	s.EncryptedNodeTuples = v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBody) SetEncryptedSessionZmk(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBody {
	s.EncryptedSessionZmk = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBody) SetMessage(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBody) SetRequestId(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBody) SetSuccess(v bool) *GetKpmEncryptedNodeTuplesByOrderIdResponseBody {
	s.Success = &v
	return s
}

type GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuples struct {
	EncryptedNodeTuple []*GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple `json:"EncryptedNodeTuple,omitempty" xml:"EncryptedNodeTuple,omitempty" type:"Repeated"`
}

func (s GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuples) String() string {
	return tea.Prettify(s)
}

func (s GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuples) GoString() string {
	return s.String()
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuples) SetEncryptedNodeTuple(v []*GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple) *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuples {
	s.EncryptedNodeTuple = v
	return s
}

type GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple struct {
	AppKeyKcv          *string `json:"AppKeyKcv,omitempty" xml:"AppKeyKcv,omitempty"`
	DevEui             *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	EncryptedAppKey    *string `json:"EncryptedAppKey,omitempty" xml:"EncryptedAppKey,omitempty"`
	EncryptedGenAppKey *string `json:"EncryptedGenAppKey,omitempty" xml:"EncryptedGenAppKey,omitempty"`
	EncryptedNwkKey    *string `json:"EncryptedNwkKey,omitempty" xml:"EncryptedNwkKey,omitempty"`
	GenAppKeyKcv       *string `json:"GenAppKeyKcv,omitempty" xml:"GenAppKeyKcv,omitempty"`
	LoraVersion        *string `json:"LoraVersion,omitempty" xml:"LoraVersion,omitempty"`
	NwkKeyKcv          *string `json:"NwkKeyKcv,omitempty" xml:"NwkKeyKcv,omitempty"`
	PinCode            *string `json:"PinCode,omitempty" xml:"PinCode,omitempty"`
}

func (s GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple) String() string {
	return tea.Prettify(s)
}

func (s GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple) GoString() string {
	return s.String()
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple) SetAppKeyKcv(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple {
	s.AppKeyKcv = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple) SetDevEui(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple {
	s.DevEui = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple) SetEncryptedAppKey(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple {
	s.EncryptedAppKey = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple) SetEncryptedGenAppKey(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple {
	s.EncryptedGenAppKey = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple) SetEncryptedNwkKey(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple {
	s.EncryptedNwkKey = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple) SetGenAppKeyKcv(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple {
	s.GenAppKeyKcv = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple) SetLoraVersion(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple {
	s.LoraVersion = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple) SetNwkKeyKcv(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple {
	s.NwkKeyKcv = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple) SetPinCode(v string) *GetKpmEncryptedNodeTuplesByOrderIdResponseBodyEncryptedNodeTuplesEncryptedNodeTuple {
	s.PinCode = &v
	return s
}

type GetKpmEncryptedNodeTuplesByOrderIdResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetKpmEncryptedNodeTuplesByOrderIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetKpmEncryptedNodeTuplesByOrderIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetKpmEncryptedNodeTuplesByOrderIdResponse) GoString() string {
	return s.String()
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponse) SetHeaders(v map[string]*string) *GetKpmEncryptedNodeTuplesByOrderIdResponse {
	s.Headers = v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponse) SetStatusCode(v int32) *GetKpmEncryptedNodeTuplesByOrderIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKpmEncryptedNodeTuplesByOrderIdResponse) SetBody(v *GetKpmEncryptedNodeTuplesByOrderIdResponseBody) *GetKpmEncryptedNodeTuplesByOrderIdResponse {
	s.Body = v
	return s
}

type SubmitKpmEncryptedNodeTupleOrderRequest struct {
	LoraVersion   *string `json:"LoraVersion,omitempty" xml:"LoraVersion,omitempty"`
	RequiredCount *int64  `json:"RequiredCount,omitempty" xml:"RequiredCount,omitempty"`
}

func (s SubmitKpmEncryptedNodeTupleOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitKpmEncryptedNodeTupleOrderRequest) GoString() string {
	return s.String()
}

func (s *SubmitKpmEncryptedNodeTupleOrderRequest) SetLoraVersion(v string) *SubmitKpmEncryptedNodeTupleOrderRequest {
	s.LoraVersion = &v
	return s
}

func (s *SubmitKpmEncryptedNodeTupleOrderRequest) SetRequiredCount(v int64) *SubmitKpmEncryptedNodeTupleOrderRequest {
	s.RequiredCount = &v
	return s
}

type SubmitKpmEncryptedNodeTupleOrderResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId        *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitKpmEncryptedNodeTupleOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitKpmEncryptedNodeTupleOrderResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitKpmEncryptedNodeTupleOrderResponseBody) SetCode(v string) *SubmitKpmEncryptedNodeTupleOrderResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitKpmEncryptedNodeTupleOrderResponseBody) SetDynamicCode(v string) *SubmitKpmEncryptedNodeTupleOrderResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SubmitKpmEncryptedNodeTupleOrderResponseBody) SetDynamicMessage(v string) *SubmitKpmEncryptedNodeTupleOrderResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SubmitKpmEncryptedNodeTupleOrderResponseBody) SetMessage(v string) *SubmitKpmEncryptedNodeTupleOrderResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitKpmEncryptedNodeTupleOrderResponseBody) SetOrderId(v int64) *SubmitKpmEncryptedNodeTupleOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *SubmitKpmEncryptedNodeTupleOrderResponseBody) SetRequestId(v string) *SubmitKpmEncryptedNodeTupleOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitKpmEncryptedNodeTupleOrderResponseBody) SetSuccess(v bool) *SubmitKpmEncryptedNodeTupleOrderResponseBody {
	s.Success = &v
	return s
}

type SubmitKpmEncryptedNodeTupleOrderResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitKpmEncryptedNodeTupleOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitKpmEncryptedNodeTupleOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitKpmEncryptedNodeTupleOrderResponse) GoString() string {
	return s.String()
}

func (s *SubmitKpmEncryptedNodeTupleOrderResponse) SetHeaders(v map[string]*string) *SubmitKpmEncryptedNodeTupleOrderResponse {
	s.Headers = v
	return s
}

func (s *SubmitKpmEncryptedNodeTupleOrderResponse) SetStatusCode(v int32) *SubmitKpmEncryptedNodeTupleOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitKpmEncryptedNodeTupleOrderResponse) SetBody(v *SubmitKpmEncryptedNodeTupleOrderResponseBody) *SubmitKpmEncryptedNodeTupleOrderResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("linkwan"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetKpmEncryptedNodeTuplesByOrderIdWithOptions(request *GetKpmEncryptedNodeTuplesByOrderIdRequest, runtime *util.RuntimeOptions) (_result *GetKpmEncryptedNodeTuplesByOrderIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetKpmEncryptedNodeTuplesByOrderId"),
		Version:     tea.String("2018-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetKpmEncryptedNodeTuplesByOrderIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetKpmEncryptedNodeTuplesByOrderId(request *GetKpmEncryptedNodeTuplesByOrderIdRequest) (_result *GetKpmEncryptedNodeTuplesByOrderIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetKpmEncryptedNodeTuplesByOrderIdResponse{}
	_body, _err := client.GetKpmEncryptedNodeTuplesByOrderIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitKpmEncryptedNodeTupleOrderWithOptions(request *SubmitKpmEncryptedNodeTupleOrderRequest, runtime *util.RuntimeOptions) (_result *SubmitKpmEncryptedNodeTupleOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoraVersion)) {
		body["LoraVersion"] = request.LoraVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RequiredCount)) {
		body["RequiredCount"] = request.RequiredCount
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitKpmEncryptedNodeTupleOrder"),
		Version:     tea.String("2018-10-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitKpmEncryptedNodeTupleOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitKpmEncryptedNodeTupleOrder(request *SubmitKpmEncryptedNodeTupleOrderRequest) (_result *SubmitKpmEncryptedNodeTupleOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitKpmEncryptedNodeTupleOrderResponse{}
	_body, _err := client.SubmitKpmEncryptedNodeTupleOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
