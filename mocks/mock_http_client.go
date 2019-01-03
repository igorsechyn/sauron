package mocks

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type MockHttpClient struct {
	mock.Mock
}

func (mockHttpClient *MockHttpClient) Get(url string) (io.ReadCloser, error) {
	args := mockHttpClient.Called(url)
	return args.Get(0).(io.ReadCloser), nil
}

func (mockHttpClient *MockHttpClient) GivenGetReturns(body io.ReadCloser) {
	mockHttpClient.ExpectedCalls = getCallsWithoutMethod(mockHttpClient.ExpectedCalls, "Get")
	mockHttpClient.On("Get", mock.Anything).Return(body)
}

func NewMockHttpClient() *MockHttpClient {
	mockHttpClient := new(MockHttpClient)
	mockHttpClient.GivenGetReturns(NewMockResponseBody())
	return mockHttpClient
}
