package venuesuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewAvailabilityEventPost() AvailabilityEventPost {
	r := AvailabilityEventPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AvailabilityEventPost struct {
	client      *Client
	queryParams *AvailabilityEventPostQueryParams
	pathParams  *AvailabilityEventPostPathParams
	method      string
	headers     http.Header
	requestBody AvailabilityEventPostBody
}

func (r AvailabilityEventPost) NewQueryParams() *AvailabilityEventPostQueryParams {
	return &AvailabilityEventPostQueryParams{}
}

type AvailabilityEventPostQueryParams struct{}

func (p AvailabilityEventPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AvailabilityEventPost) QueryParams() *AvailabilityEventPostQueryParams {
	return r.queryParams
}

func (r AvailabilityEventPost) NewPathParams() *AvailabilityEventPostPathParams {
	return &AvailabilityEventPostPathParams{}
}

type AvailabilityEventPostPathParams struct {
	VenueID int `schema:"venue_id"`
}

func (p *AvailabilityEventPostPathParams) Params() map[string]string {
	return map[string]string{
		"venue_id": strconv.Itoa(p.VenueID),
	}
}

func (r *AvailabilityEventPost) PathParams() *AvailabilityEventPostPathParams {
	return r.pathParams
}

func (r *AvailabilityEventPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AvailabilityEventPost) SetMethod(method string) {
	r.method = method
}

func (r *AvailabilityEventPost) Method() string {
	return r.method
}

func (r AvailabilityEventPost) NewRequestBody() AvailabilityEventPostBody {
	return AvailabilityEventPostBody{}
}

type AvailabilityEventPostBody struct {
	Start   DateTime `json:"start"`
	End     DateTime `json:"end"`
	Title   string   `json:"title"`
	SpaceID int      `json:"space_id"`
}

// func (r *AvailabilityEventPost) MarshalJSON() ([]byte, error) {
// 	return omitempty.MarshalJSON(r)
// }

func (r *AvailabilityEventPost) RequestBody() *AvailabilityEventPostBody {
	return &r.requestBody
}

func (r *AvailabilityEventPost) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *AvailabilityEventPost) SetRequestBody(body AvailabilityEventPostBody) {
	r.requestBody = body
}

func (r *AvailabilityEventPost) NewResponseBody() *AvailabilityEventPostResponseBody {
	return &AvailabilityEventPostResponseBody{}
}

type AvailabilityEventPostResponseBody AvailabilityEvent

func (r *AvailabilityEventPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/venues/{{.venue_id}}/availability_events", r.PathParams())
	return &u
}

func (r *AvailabilityEventPost) Do() (AvailabilityEventPostResponseBody, error) {
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
