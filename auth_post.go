package venuesuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewAuthPost() AuthPost {
	r := AuthPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AuthPost struct {
	client      *Client
	queryParams *AuthPostQueryParams
	pathParams  *AuthPostPathParams
	method      string
	headers     http.Header
	requestBody AuthPostBody
}

func (r AuthPost) NewQueryParams() *AuthPostQueryParams {
	return &AuthPostQueryParams{}
}

type AuthPostQueryParams struct{}

func (p AuthPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AuthPost) QueryParams() *AuthPostQueryParams {
	return r.queryParams
}

func (r AuthPost) NewPathParams() *AuthPostPathParams {
	return &AuthPostPathParams{}
}

type AuthPostPathParams struct {
}

func (p *AuthPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AuthPost) PathParams() *AuthPostPathParams {
	return r.pathParams
}

func (r *AuthPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AuthPost) SetMethod(method string) {
	r.method = method
}

func (r *AuthPost) Method() string {
	return r.method
}

func (r AuthPost) NewRequestBody() AuthPostBody {
	return AuthPostBody{}
}

type AuthPostBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// func (r *AuthPost) MarshalJSON() ([]byte, error) {
// 	return omitempty.MarshalJSON(r)
// }

func (r *AuthPost) RequestBody() *AuthPostBody {
	return &r.requestBody
}

func (r *AuthPost) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *AuthPost) SetRequestBody(body AuthPostBody) {
	r.requestBody = body
}

func (r *AuthPost) NewResponseBody() *AuthPostResponseBody {
	return &AuthPostResponseBody{}
}

type AuthPostResponseBody struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Status       int    `json:"status"`
	Msg          string `json:"msg"`
}

func (r *AuthPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/auth", r.PathParams())
	return &u
}

func (r *AuthPost) Do() (AuthPostResponseBody, error) {
	r.RequestBody().Username = r.client.User()
	r.RequestBody().Password = r.client.Password()

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
