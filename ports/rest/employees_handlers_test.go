package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/morlfm/rest-api/application/model"
	"github.com/morlfm/rest-api/ports/rest/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetEmployees(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/employees", nil)

	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	appEmpMock := &mocks.EmpApp{}
	target := EmpHandler{app: appEmpMock}
	mockList := []model.Employee{
		{
			ID:   "1",
			Name: "chico",
			Role: "dev",
		},
	}

	appEmpMock.On("GetFilterEmployees", model.Employee{}).Return(mockList, nil)
	target.GetFilterEmployees(resp, req)

	response := resp.Result()
	data, _ := ioutil.ReadAll(response.Body)
	parsedData := strings.Replace(string(data), "\n", "", 1)

	expected := `[{"id":"1","name":"chico","role":"dev"}]`
	assert.Equal(t, expected, parsedData)

}

func TestCreateEmployee(t *testing.T) {

}
