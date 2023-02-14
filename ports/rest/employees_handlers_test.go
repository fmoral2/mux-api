package api

import (
	"net/http/httptest"
	"testing"
)

// func TestGetEmployees(t *testing.T) {
//
// 	req, err := http.NewRequest("GET", "/api/employees", nil)
//
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	resp := httptest.NewRecorder()
//
// 	appEmpMock := &mocks.EmpApp{}
// 	target := EmpHandler{app: appEmpMock}
// 	mockList := []model.Employee{
// 		{
// 			ID:   "1",
// 			Name: "chico",
// 			Role: "dev",
// 		},
// 	}
//
// 	appEmpMock.On("GetFilterEmployees", model.Employee{}).Return(mockList, nil)
// 	target.GetFilterEmployees(resp, req)
//
// 	response := resp.Result()
// 	data, _ := ioutil.ReadAll(response.Body)
// 	parsedData := strings.Replace(string(data), "\n", "", 1)
//
// 	expected := `[{"id":"1","name":"chico","role":"dev"}]`
// 	assert.Equal(t, expected, parsedData)
//
// }
//
// func TestCreateEmployee(t *testing.T) {
//
// }

func BenchmarkGetPageRequest(b *testing.B) {
	req := httptest.NewRequest("GET", "/?page=5&size=10", nil)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = getPageRequest(req)
	}
}
