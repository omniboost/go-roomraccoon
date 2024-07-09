package roomraccoon

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-roomraccoon/utils"
)

func (c *Client) NewBookingNotesPost() BookingNotesPost {
	r := BookingNotesPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BookingNotesPost struct {
	client      *Client
	queryParams *BookingNotesPostQueryParams
	pathParams  *BookingNotesPostPathParams
	method      string
	headers     http.Header
	requestBody BookingNotesPostBody
}

func (r BookingNotesPost) NewQueryParams() *BookingNotesPostQueryParams {
	return &BookingNotesPostQueryParams{}
}

type BookingNotesPostQueryParams struct {
	Source       CommaSeparatedQueryParam `schema:"source,omitempty"`
	Status       CommaSeparatedQueryParam `schema:"status,omitempty"`
	PeriodStart  Date                     `schema:"period[start],omitempty"`
	PeriodEnd    Date                     `schema:"period[end],omitempty"`
	CreatedFrom  Date                     `schema:"created_from,omitempty"`
	CreatedUntil Date                     `schema:"created_until,omitempty"`
	Limit        int                      `schema:"limit,omitempty"`
	Page         int                      `schema:"page,omitempty"`
}

func (p BookingNotesPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(CommaSeparatedQueryParam{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *BookingNotesPost) QueryParams() *BookingNotesPostQueryParams {
	return r.queryParams
}

func (r BookingNotesPost) NewPathParams() *BookingNotesPostPathParams {
	return &BookingNotesPostPathParams{}
}

type BookingNotesPostPathParams struct {
	BookingID int `schema:"booking_id"`
}

func (p *BookingNotesPostPathParams) Params() map[string]string {
	return map[string]string{
		"booking_id": strconv.Itoa(p.BookingID),
	}
}

func (r *BookingNotesPost) PathParams() *BookingNotesPostPathParams {
	return r.pathParams
}

func (r *BookingNotesPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BookingNotesPost) SetMethod(method string) {
	r.method = method
}

func (r *BookingNotesPost) Method() string {
	return r.method
}

func (r BookingNotesPost) NewRequestBody() BookingNotesPostBody {
	return BookingNotesPostBody{}
}

type BookingNotesPostBody struct {
	Content string `json:"content"`
}

func (r *BookingNotesPost) RequestBody() *BookingNotesPostBody {
	return &r.requestBody
}

func (r *BookingNotesPost) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *BookingNotesPost) SetRequestBody(body BookingNotesPostBody) {
	r.requestBody = body
}

func (r *BookingNotesPost) NewResponseBody() *BookingNotesPostResponseBody {
	return &BookingNotesPostResponseBody{}
}

type BookingNotesPostResponseBody BookingNote

func (r *BookingNotesPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/bookings/{{.booking_id}}/notes", r.PathParams())
	return &u
}

func (r *BookingNotesPost) Do() (BookingNotesPostResponseBody, error) {
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
