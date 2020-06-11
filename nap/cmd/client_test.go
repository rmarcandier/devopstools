package nap

import (

)

func TestProcessRequest(t *testing.T) {
	client := NewClient()
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response) error {

	})
	rest := NewResource("/get", "GET",
	client.ProcessRequest("https://httpbin.org"); err != nil {
		t.Fail()
	}
)
}