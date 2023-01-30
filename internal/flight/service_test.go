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

	suite.Assert().ErrorContains(err, "flight is not connected")
}

func (suite *FlightServiceTestSuite) TestCirculateFlights() {
	_, _, err := suite.service.GetCityPair(context.Background(), &flight.GetCityPairRequest{
		Flights: [][]string{{"IND", "EWR"}, {"EWR", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
	})

	suite.Assert().ErrorContains(err, "circulate flight")
}

func (suite *FlightServiceTestSuite) TestInputValidation() {
	{

		cases := []struct {
			name    string
			input   [][]string
			errMsg  string
			wantErr bool
		}{
			{
				name:    "valid input",
				input:   [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
				wantErr: false,
			},
			{
				name:    "invalid len",
				input:   [][]string{{"IND"}},
				errMsg:  "Field validation for 'Flights[0]' failed on the 'len' tag",
				wantErr: true,
			},
			{
				name:    "minimum 1 flight",
				input:   [][]string{},
				errMsg:  "Field validation for 'Flights' failed on the 'min' tag",
				wantErr: true,
			},
			{
				name:    "duplicated city",
				input:   [][]string{{"IND", "IND"}},
				errMsg:  "Field validation for 'Flights[0]' failed on the 'unique' tag",
				wantErr: true,
			},
		}

		for _, c := range cases {
			_, _, err := suite.service.GetCityPair(context.Background(), &flight.GetCityPairRequest{
				Flights: c.input,
			})

			if c.wantErr {
				suite.Assert().ErrorContains(err, c.errMsg)
			} else {
				suite.NoError(err)
			}
		}

	}
}

func TestFlightServiceTestSuite(t *testing.T) {
	suite.Run(t, new(FlightServiceTestSuite))
}
