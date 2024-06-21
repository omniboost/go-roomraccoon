package venuesuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewBookingActivitiesGet() BookingActivitiesGet {
	r := BookingActivitiesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BookingActivitiesGet struct {
	client      *Client
	queryParams *BookingActivitiesGetQueryParams
	pathParams  *BookingActivitiesGetPathParams
	method      string
	headers     http.Header
	requestBody BookingActivitiesGetBody
}

func (r BookingActivitiesGet) NewQueryParams() *BookingActivitiesGetQueryParams {
	return &BookingActivitiesGetQueryParams{}
}

type BookingActivitiesGetQueryParams struct {
	Source       CommaSeparatedQueryParam `schema:"source,omitempty"`
	Status       CommaSeparatedQueryParam `schema:"status,omitempty"`
	PeriodStart  Date                     `schema:"period[start],omitempty"`
	PeriodEnd    Date                     `schema:"period[end],omitempty"`
	CreatedFrom  Date                     `schema:"created_from,omitempty"`
	CreatedUntil Date                     `schema:"created_until,omitempty"`
	Limit        int                      `schema:"limit,omitempty"`
	Page         int                      `schema:"page,omitempty"`
}

func (p BookingActivitiesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BookingActivitiesGet) QueryParams() *BookingActivitiesGetQueryParams {
	return r.queryParams
}

func (r BookingActivitiesGet) NewPathParams() *BookingActivitiesGetPathParams {
	return &BookingActivitiesGetPathParams{}
}

type BookingActivitiesGetPathParams struct {
	BookingID int `schema:"booking_id"`
}

func (p *BookingActivitiesGetPathParams) Params() map[string]string {
	return map[string]string{
		"booking_id": strconv.Itoa(p.BookingID),
	}
}

func (r *BookingActivitiesGet) PathParams() *BookingActivitiesGetPathParams {
	return r.pathParams
}

func (r *BookingActivitiesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BookingActivitiesGet) SetMethod(method string) {
	r.method = method
}

func (r *BookingActivitiesGet) Method() string {
	return r.method
}

func (r BookingActivitiesGet) NewRequestBody() BookingActivitiesGetBody {
	return BookingActivitiesGetBody{}
}

type BookingActivitiesGetBody struct {
}

func (r *BookingActivitiesGet) RequestBody() *BookingActivitiesGetBody {
	return nil
}

func (r *BookingActivitiesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *BookingActivitiesGet) SetRequestBody(body BookingActivitiesGetBody) {
	r.requestBody = body
}

func (r *BookingActivitiesGet) NewResponseBody() *BookingActivitiesGetResponseBody {
	return &BookingActivitiesGetResponseBody{}
}

type BookingActivitiesGetResponseBody struct {
	Data BookingActivities `json:"data"`
	Meta struct {
		Pagination Pagination `json:"pagination"`
	} `json:"meta"`
}

func (r *BookingActivitiesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/bookings/{{.booking_id}}/activity", r.PathParams())
	return &u
}

func (r *BookingActivitiesGet) Do() (BookingActivitiesGetResponseBody, error) {
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

func (r *BookingActivitiesGet) All() (BookingActivities, error) {
	activities := BookingActivities{}
	response := *r.NewResponseBody()
	for {
		resp, err := r.Do()
		if err != nil {
			return activities, err
		}

		// Break out of loop when no activities are found
		if len(resp.Data) == 0 {
			break
		}

		// Add activities to list
		activities = append(activities, resp.Data...)

		if response.Meta.Pagination.CurrentPage == response.Meta.Pagination.TotalPages {
			break
		}

		// Increment page number
		r.QueryParams().Page = response.Meta.Pagination.CurrentPage + 1
	}

	return activities, nil
}
