package flight_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"citypair/internal/flight"
	"citypair/pkg/log"
	mock "citypair/testutil/mocks/flight"
)

type FlightHttpTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	ts   *httptest.Server
}

func (suite *FlightHttpTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	flightHTTP := NewUserHTTPMock(suite.ctrl, validator.New())
	handler := flight.RegisterHTTPHandlers(flightHTTP)
	suite.ts = httptest.NewServer(handler)
}

func (suite *FlightHttpTestSuite) TearDownTest() {
	suite.ctrl.Finish()
	suite.ts.Close()
}

func (suite *FlightHttpTestSuite) TestGetCityPair() {
	resp, body := testRequest(suite.T(), suite.ts, http.MethodPost, "/calculate", strings.NewReader("{}"))
	suite.Assert().Equalf(200, resp.StatusCode, "Expected status code: 200, got: %d", resp.StatusCode)
	suite.Assert().Containsf(resp.Header.Get("Content-Type"), "application/json", "Invalid content type, expected application/json, got %s", resp.Header.Get("Content-Type"))
	suite.Assert().NotEmpty(body)
	suite.Assert().Equal(body, "\"[FOO, BAR]\"\n")
}

func TestFlightHttpTestSuite(t *testing.T) {
	suite.Run(t, new(FlightHttpTestSuite))
}

func NewUserHTTPMock(ctrl *gomock.Controller, validator *validator.Validate) flight.HTTP {
	svc := mock.NewMockService(ctrl)

	req := &flight.GetCityPairRequest{}

	svc.EXPECT().GetCityPair(gomock.Any(), req).Return("FOO", "BAR", nil)
	logger := log.New()
	return flight.GetFlightHTTP(svc, logger)
}

func testRequest(
	t *testing.T,
	ts *httptest.Server,
	method,
	path string,
	body io.Reader,
) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	return resp, string(respBody)
}
