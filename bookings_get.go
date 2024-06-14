package venuesuite

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewBookingsGet() BookingsGet {
	r := BookingsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BookingsGet struct {
	client      *Client
	queryParams *BookingsGetQueryParams
	pathParams  *BookingsGetPathParams
	method      string
	headers     http.Header
	requestBody BookingsGetBody
}

func (r BookingsGet) NewQueryParams() *BookingsGetQueryParams {
	return &BookingsGetQueryParams{}
}

type BookingsGetQueryParams struct {
	Status       []string `schema:"status,omitempty"`
	PeriodStart  Date     `schema:"period[start],omitempty"`
	PeriodEnd    Date     `schema:"period[end],omitempty"`
	CreatedFrom  Date     `schema:"created_from,omitempty"`
	CreatedUntil Date     `schema:"created_until,omitempty"`
	Limit        int      `schema:"limit,omitempty"`
	Page         int      `schema:"page,omitempty"`
}

func (p BookingsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BookingsGet) QueryParams() *BookingsGetQueryParams {
	return r.queryParams
}

func (r BookingsGet) NewPathParams() *BookingsGetPathParams {
	return &BookingsGetPathParams{}
}

type BookingsGetPathParams struct{}

func (p *BookingsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BookingsGet) PathParams() *BookingsGetPathParams {
	return r.pathParams
}

func (r *BookingsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BookingsGet) SetMethod(method string) {
	r.method = method
}

func (r *BookingsGet) Method() string {
	return r.method
}

func (r BookingsGet) NewRequestBody() BookingsGetBody {
	return BookingsGetBody{}
}

type BookingsGetBody struct {
}

func (r *BookingsGet) RequestBody() *BookingsGetBody {
	return nil
}

func (r *BookingsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *BookingsGet) SetRequestBody(body BookingsGetBody) {
	r.requestBody = body
}

func (r *BookingsGet) NewResponseBody() *BookingsGetResponseBody {
	return &BookingsGetResponseBody{}
}

type BookingsGetResponseBody struct {
	Data Bookings `json:"data"`
	Meta struct {
		Pagination Pagination `json:"pagination"`
	} `json:"meta"`
}

func (r *BookingsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/bookings", r.PathParams())
	return &u
}

func (r *BookingsGet) Do() (BookingsGetResponseBody, error) {
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

func (r *BookingsGet) All() (Bookings, error) {
	bookings := Bookings{}
	response := *r.NewResponseBody()
	for {
		resp, err := r.Do()
		if err != nil {
			return bookings, err
		}

		// Break out of loop when no bookings are found
		if len(resp.Data) == 0 {
			break
		}

		// Add bookings to list
		bookings = append(bookings, resp.Data...)

		if response.Meta.Pagination.CurrentPage == response.Meta.Pagination.TotalPages {
			break
		}

		// Increment page number
		r.QueryParams().Page = response.Meta.Pagination.CurrentPage + 1
	}

	return bookings, nil
}
