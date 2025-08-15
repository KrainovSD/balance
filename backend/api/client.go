package api

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type RequestMethod string
type MethodEnum struct {
	GET     RequestMethod
	POST    RequestMethod
	DELETE  RequestMethod
	PUT     RequestMethod
	PATCH   RequestMethod
	OPTIONS RequestMethod
}

var Methods = MethodEnum{
	GET:     "GET",
	POST:    "POST",
	DELETE:  "DELETE",
	PUT:     "PUT",
	PATCH:   "PATCH",
	OPTIONS: "OPTIONS",
}

type RequestContentType string
type ContentTypeEnum struct {
	JSON   RequestContentType
	Text   RequestContentType
	Form   RequestContentType
	Stream RequestContentType
}

var ContentTypes = ContentTypeEnum{
	JSON:   "application/json",
	Text:   "plain/text",
	Form:   "application/x-www-form-urlencoded",
	Stream: "application/octet-stream",
}

type ApiClient struct {
}
type Request struct {
	Url         string
	Queries     map[string][]string
	Headers     map[string]string
	Method      RequestMethod
	ContentType RequestContentType
	Body        io.Reader
}
type Response struct {
	Data   []byte
	Status int
	Header http.Header
}

func CreateApiClient() (*ApiClient, error) {
	return &ApiClient{}, nil
}

func (c *ApiClient) Send(request Request) (Response, error) {
	var result Response
	var content []byte
	var err error
	var req *http.Request
	var res *http.Response
	var requestUrl *url.URL
	var client *http.Client = &http.Client{}

	if requestUrl, err = url.Parse(request.Url); err != nil {
		return result, err
	}
	if len(request.Queries) > 0 {
		var queries = requestUrl.Query()
		for k, vArr := range request.Queries {
			if len(vArr) == 0 {
				continue
			}

			for _, v := range vArr {
				queries.Add(k, v)
			}
		}
	}
	if req, err = http.NewRequest(string(request.Method), requestUrl.String(), request.Body); err != nil {
		return result, err
	}
	if len(request.Headers) > 0 {
		for k, v := range request.Headers {
			req.Header.Add(k, v)
		}
	}
	req.Header.Add("Content-Type", string(request.ContentType))

	if res, err = client.Do(req); err != nil {
		return result, err
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		err = errors.New("bad status code " + strconv.Itoa(res.StatusCode))
		return result, err
	}

	if content, err = io.ReadAll(res.Body); err != nil {
		return result, err
	}

	result = Response{
		Data:   content,
		Status: res.StatusCode,
		Header: res.Header,
	}

	return result, nil

}
