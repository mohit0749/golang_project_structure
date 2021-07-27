package module1

import "context"

var defaultService Service

type SeatStatus int

const (
	Confirm   SeatStatus = 1
	Failed    SeatStatus = 2
	InProcess SeatStatus = 3
)

type Details struct {
	From, To, Description string
	Cost                  int
}

type SeatInfo struct {
	Status SeatStatus
}

type Service interface {
	BookSeat(ctx context.Context, id int64, seatID int64) (*SeatInfo, error)
	//GetAllFlights(ctx context.Context, from, to string) ([]Details,error)
	GetFlightDetails(ctx context.Context, id int64) (*Details, error)
}

func NewService(service Service) {
	defaultService = service
}

func GetService() Service {
	return defaultService
}
