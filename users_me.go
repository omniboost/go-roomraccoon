package venuesuite

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewUsersMe() UsersMe {
	r := UsersMe{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type UsersMe struct {
	client      *Client
	queryParams *UsersMeQueryParams
	pathParams  *UsersMePathParams
	method      string
	headers     http.Header
	requestBody UsersMeBody
}

func (r UsersMe) NewQueryParams() *UsersMeQueryParams {
	return &UsersMeQueryParams{}
}

type UsersMeQueryParams struct {
}

func (p UsersMeQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *UsersMe) QueryParams() QueryParams {
	return r.queryParams
}

func (r UsersMe) NewPathParams() *UsersMePathParams {
	return &UsersMePathParams{}
}

type UsersMePathParams struct {
}

func (p *UsersMePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *UsersMe) PathParams() *UsersMePathParams {
	return r.pathParams
}

func (r *UsersMe) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *UsersMe) SetMethod(method string) {
	r.method = method
}

func (r *UsersMe) Method() string {
	return r.method
}

func (r UsersMe) NewRequestBody() UsersMeBody {
	return UsersMeBody{}
}

type UsersMeBody struct {
}

func (r *UsersMe) RequestBody() *UsersMeBody {
	return nil
}

func (r *UsersMe) RequestBodyInterface() interface{} {
	return nil
}

func (r *UsersMe) SetRequestBody(body UsersMeBody) {
	r.requestBody = body
}

func (r *UsersMe) NewResponseBody() *UsersMeResponseBody {
	return &UsersMeResponseBody{}
}

type UsersMeResponseBody User

func (r *UsersMe) URL() *url.URL {
	u := r.client.GetEndpointURL("/users/me", r.PathParams())
	return &u
}

func (r *UsersMe) Do() (UsersMeResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	token, err := r.client.Token()
	if err != nil {
		return *r.NewResponseBody(), err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
