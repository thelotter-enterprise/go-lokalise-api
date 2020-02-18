package lokalise

import "fmt"

const (
	pathProcesses = "processes"
)

// ProcessService handles communication with the Process related
// methods of the Lokalise API.
//
// Lokalise API docs: https://lokalise.com/api2docs/curl/#resource-Processes
type ProcessService struct {
	BaseService
}

// ‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Service entity objects
// _____________________________________________________________________________________________________________________

type Process struct {
	WithCreationTime

	ProcessID string `json:"process_id"`
	Type      string `json:"type"`
	Status    string `json:"status"`
}

// Retrieves a Process
func (c *ProcessService) Retrieve(processID string) (r Process, err error) {
	resp, err := c.get(c.Ctx(), pathProcessByID(processID), &r)

	if err != nil {
		return
	}
	return r, apiError(resp)
}

func pathProcessByID(processID string) string {
	return fmt.Sprintf("%s/%s", pathProcesses, processID)
}
