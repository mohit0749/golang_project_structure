package service

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"test.com/project_structure/internal/module1"

	"github.com/eapache/go-resiliency/breaker"
)

func Test_flightService_BookSeat(t *testing.T) {
	type fields struct {
		addr       string
		httpClient *http.Client
		breaker    *breaker.Breaker
	}
	type args struct {
		ctx      context.Context
		flightId int64
		seatID   int64
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		mockHttpHandler func(w http.ResponseWriter, r *http.Request)
		want            *module1.SeatInfo
		wantErr         bool
	}{
		{
			"test success",
			fields{
				httpClient: http.DefaultClient,
				breaker:    breaker.New(11, 11, 11111),
			},
			args{
				ctx:      context.Background(),
				flightId: 1,
				seatID:   11,
			},
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(`{"status":1}`))
			},
			&module1.SeatInfo{Status: module1.Confirm},
			false,
		},

		{
			"test error",
			fields{
				httpClient: http.DefaultClient,
				breaker:    breaker.New(11, 11, 11111),
			},
			args{
				ctx:      context.Background(),
				flightId: 1,
				seatID:   11,
			},
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServ := httptest.NewServer(http.HandlerFunc(tt.mockHttpHandler))
			defer mockServ.Close()
			fs := &testService{
				addr:       mockServ.URL,
				httpClient: tt.fields.httpClient,
				breaker:    tt.fields.breaker,
			}
			got, err := fs.BookSeat(tt.args.ctx, tt.args.flightId, tt.args.seatID)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookSeat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookSeat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flightService_GetFlightDetails(t *testing.T) {
	type fields struct {
		addr       string
		httpClient *http.Client
		breaker    *breaker.Breaker
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		mockHttpHandler func(w http.ResponseWriter, r *http.Request)
		want            *module1.Details
		wantErr         bool
	}{
		{
			"test success",
			fields{
				httpClient: http.DefaultClient,
				breaker:    breaker.New(11, 11, 111),
			},
			args{
				ctx: context.Background(),
				id:  1,
			},
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(`{"src":"Delhi","dest":"Mumbai","cost":10000}`))
			},
			&module1.Details{
				From:        "Delhi",
				To:          "Mumbai",
				Description: "",
				Cost:        10000,
			},
			false,
		},
		{
			"test error",
			fields{
				httpClient: http.DefaultClient,
				breaker:    breaker.New(11, 11, 111),
			},
			args{
				ctx: context.Background(),
				id:  1,
			},
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServ := httptest.NewServer(http.HandlerFunc(tt.mockHttpHandler))
			defer mockServ.Close()
			fs := &testService{
				addr:       mockServ.URL,
				httpClient: tt.fields.httpClient,
				breaker:    tt.fields.breaker,
			}
			got, err := fs.GetFlightDetails(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFlightDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFlightDetails() got = %v, want %v", got, tt.want)
			}
		})
	}
}
