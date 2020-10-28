package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/Jopoleon/AtlantTest/logger"

	"github.com/Jopoleon/AtlantTest/models"
)

func TestCSVFetchClient_GetCSV(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		equals(t, req.URL.String(), "/")
		rw.Write([]byte(testData))
	}))
	defer server.Close()

	client := NewCSVFetchClient(logger.NewLogger("0"))
	pr, err := client.GetCSV(server.URL)
	if err != nil {
		t.Error(err)
	}
	equals(t, pr[0], models.Product{
		Name:      "product1",
		LastPrice: 45.66,
	})
}

var testData = `
"NAME","PRICE"
# this is a comment
"product1",45.66
product2,"16.88"
"product3","99.00"
`

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
