package lokalise

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestProcessService_Retrieve(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	processId := "test"
	mux.HandleFunc(
		fmt.Sprintf("/projects/%s/processes/%s", testProjectID, processId),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			testMethod(t, r, "GET")
			testHeader(t, r, apiTokenHeader, testApiToken)

			_, _ = fmt.Fprint(w, `{
				"process_id": "`+processId+`",
				"type": "type",
				"status": "status"
			}`)
		})

	r, err := client.Processes().Retrieve(testProjectID, processId)
	if err != nil {
		t.Errorf("Processes.Retrieve returned error: %v", err)
	}

	want := Process{
		ProcessID: processId,
		Type:      "type",
		Status:    "status",
	}

	if !reflect.DeepEqual(r, want) {
		t.Errorf("Processes.Retrieve returned %+v, want %+v", r, want)
	}
}
