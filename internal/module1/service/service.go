package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"test.com/project_structure/internal/module1"

	"github.com/eapache/go-resiliency/breaker"
)

const (
	//Http Client Default Config
	defaultHttpTimeout     = 30 * time.Second
	defaultHttpMaxIdleConn = 10

	//Circuit Breaker Default Config
	defaultBreakerErrorThreshold   = 10
	defaultBreakerSuccessThreshold = 5
	defaultBreakerTimeout          = 30 * time.Second
)

type testService struct {
	addr       string
	httpClient *http.Client
	breaker    *breaker.Breaker
}

type Option func(cfg *testService)

func WithHttpOption(timeout time.Duration, maxIdleConn int) Option {
	return func(service *testService) {
		if timeout == 0 {
			timeout = defaultHttpTimeout
		}

		if maxIdleConn == 0 {
			maxIdleConn = defaultHttpMaxIdleConn
		}

		service.httpClient = &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:    maxIdleConn,
				IdleConnTimeout: timeout,
			},
		}
	}
}

func WithBreakerOption(errorThreshold, successThreshold int, timeout time.Duration) Option {
	return func(service *testService) {
		if errorThreshold == 0 {
			errorThreshold = defaultBreakerErrorThreshold
		}

		if successThreshold == 0 {
			successThreshold = defaultBreakerSuccessThreshold
		}

		if timeout == 0 {
			timeout = defaultBreakerTimeout
		}

		service.breaker = breaker.New(errorThreshold, successThreshold, timeout)
	}
}

func NewTestService(addr string, options ...Option) *testService {

	service := &testService{
		addr: addr,
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:    defaultHttpMaxIdleConn,
				IdleConnTimeout: defaultHttpTimeout,
			},
		},
		breaker: breaker.New(defaultBreakerErrorThreshold, defaultBreakerSuccessThreshold, defaultBreakerTimeout),
	}

	for _, opt := range options {
		opt(service)
	}

	return service
}

func (fs *testService) BookSeat(ctx context.Context, flightId int64, seatID int64) (*module1.SeatInfo, error) {
	//TODO: add tracing
	reqBody := bookSeatRequest{FlightID: flightId, SeatID: seatID}
	resp, err := fs.newBookSeatRequest(ctx, reqBody)
	if err != nil {
		//TODO: add logging
		return nil, err
	}
	return parseBookSeatResponse(resp)
}

type bookSeatRequest struct {
	FlightID int64 `json:"flight_id"`
	SeatID   int64 `json:"seat_id"`
}

type bookSeatResponse struct {
	Status int `json:"status"`
}

func (fs *testService) newBookSeatRequest(ctx context.Context, reqBody bookSeatRequest) (*http.Response, error) {
	var err error
	var payload []byte
	var resp *http.Response
	payload, err = json.Marshal(reqBody)
	if err != nil {
		return resp, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fs.addr, bytes.NewBuffer(payload))
	err = fs.breaker.Run(func() error {
		resp, err = fs.httpClient.Do(req)
		return err
	})

	//TODO: handler non-200 status code
	if resp.StatusCode > 200 {
		return nil, fmt.Errorf("service responded with %s", resp.Status)
	}

	if err != nil {
		return resp, err
	}
	return resp, nil
}

func parseBookSeatResponse(resp *http.Response) (*module1.SeatInfo, error) {
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var data bookSeatResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	var seatInfo module1.SeatInfo
	seatInfo.Status = module1.SeatStatus(data.Status)
	return &seatInfo, nil
}

func (fs *testService) GetFlightDetails(ctx context.Context, id int64) (*module1.Details, error) {
	//TODO: add tracing
	reqBody := flightDetailsRequest{Id: id}
	resp, err := fs.newFlightDetailsRequest(ctx, reqBody)
	if err != nil {
		//TODO:add logging
		return nil, err
	}
	return parseFlightDetailsResponse(resp)
}

type flightDetailsRequest struct {
	Id int64 `json:"id"`
}

type flightDetailsResponse struct {
	Src         string `json:"src"`
	Dest        string `json:"dest"`
	Description string `json:"desc,omitempty"`
	Cost        int    `json:"cost,omitempty"`
}

func (fs *testService) newFlightDetailsRequest(ctx context.Context, reqBody flightDetailsRequest) (*http.Response, error) {
	var err error
	var payload []byte
	var resp *http.Response
	payload, err = json.Marshal(reqBody)
	if err != nil {
		return resp, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fs.addr, bytes.NewBuffer(payload))
	err = fs.breaker.Run(func() error {
		resp, err = fs.httpClient.Do(req)
		return err
	})

	//TODO: handler non-200 status code
	if resp.StatusCode > 200 {
		return nil, fmt.Errorf("service responded with %s", resp.Status)
	}
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func parseFlightDetailsResponse(resp *http.Response) (*module1.Details, error) {
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var data flightDetailsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	var details module1.Details
	details.Cost = data.Cost
	details.Description = data.Description
	details.From = data.Src
	details.To = data.Dest
	return &details, nil
}
