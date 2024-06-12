package venuesuite

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewBookingByIDGet() BookingByIDGet {
	r := BookingByIDGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BookingByIDGet struct {
	client      *Client
	queryParams *BookingByIDGetQueryParams
	pathParams  *BookingByIDGetPathParams
	method      string
	headers     http.Header
	requestBody BookingByIDGetBody
}

func (r BookingByIDGet) NewQueryParams() *BookingByIDGetQueryParams {
	return &BookingByIDGetQueryParams{}
}

type BookingByIDGetQueryParams struct {
}

func (p BookingByIDGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BookingByIDGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r BookingByIDGet) NewPathParams() *BookingByIDGetPathParams {
	return &BookingByIDGetPathParams{}
}

type BookingByIDGetPathParams struct {
	BookingID int `schema:"booking_id"`
}

func (p *BookingByIDGetPathParams) Params() map[string]string {
	return map[string]string{
		"booking_id": strconv.Itoa(p.BookingID),
	}
}

func (r *BookingByIDGet) PathParams() *BookingByIDGetPathParams {
	return r.pathParams
}

func (r *BookingByIDGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BookingByIDGet) SetMethod(method string) {
	r.method = method
}

func (r *BookingByIDGet) Method() string {
	return r.method
}

func (r BookingByIDGet) NewRequestBody() BookingByIDGetBody {
	return BookingByIDGetBody{}
}

type BookingByIDGetBody struct {
}

func (r *BookingByIDGet) RequestBody() *BookingByIDGetBody {
	return nil
}

func (r *BookingByIDGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *BookingByIDGet) SetRequestBody(body BookingByIDGetBody) {
	r.requestBody = body
}

func (r *BookingByIDGet) NewResponseBody() *BookingByIDGetResponseBody {
	return &BookingByIDGetResponseBody{}
}

type BookingByIDGetResponseBody Booking

func (r *BookingByIDGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/bookings/{{.booking_id}}", r.PathParams())
	return &u
}

func (r *BookingByIDGet) Do() (BookingByIDGetResponseBody, error) {
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
