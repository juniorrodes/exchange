package mock

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type DoerMock struct {
    mock.Mock
}

func (c *DoerMock) Do(request *http.Request) (*http.Response, error) {
    args := c.Called(request)

    return args.Get(0).(*http.Response), args.Error(1)
}
