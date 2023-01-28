package flight_test

import (
	"citypair/internal/flight"
	"citypair/pkg/log"
	"context"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type FlightServiceTestSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	service flight.Service
}

func (suite *FlightServiceTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())

	logger := log.New()
	validate := validator.New()

	suite.service = flight.NewService(validate, logger)
}

func (suite *FlightServiceTestSuite) TestOneFlights() {
	src, dest, err := suite.service.GetCityPair(context.Background(), &flight.GetCityPairRequest{
		Flights: [][]string{{"SFO", "EWR"}},
	})

	suite.NoError(err)
	suite.Equal("SFO", src)
	suite.Equal("EWR", dest)
}

func (suite *FlightServiceTestSuite) TestTwoFlights() {
	src, dest, err := suite.service.GetCityPair(context.Background(), &flight.GetCityPairRequest{
		Flights: [][]string{{"SFO", "ATL"}, {"ATL", "EWR"}},
	})

	suite.NoError(err)
	suite.Equal("SFO", src)
	suite.Equal("EWR", dest)
}

func (suite *FlightServiceTestSuite) TestFourFlights() {
	src, dest, err := suite.service.GetCityPair(context.Background(), &flight.GetCityPairRequest{
		Flights: [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
	})

	suite.NoError(err)
	suite.Equal("SFO", src)
	suite.Equal("EWR", dest)
}

func (suite *FlightServiceTestSuite) TestNonConnectedFlights() {
	_, _, err := suite.service.GetCityPair(context.Background(), &flight.GetCityPairRequest{
		Flights: [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"ATL", "GSO"}},
	})

	suite.Error(err, "flight is not connected")
}

func (suite *FlightServiceTestSuite) TestCirculateFlights() {
	_, _, err := suite.service.GetCityPair(context.Background(), &flight.GetCityPairRequest{
		Flights: [][]string{{"IND", "EWR"}, {"EWR", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
	})

	suite.Error(err, "circulate flight")
}

func TestFlightServiceTestSuite(t *testing.T) {
	suite.Run(t, new(FlightServiceTestSuite))
}
