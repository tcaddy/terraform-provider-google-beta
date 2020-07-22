// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------
package google

import (
	"encoding/json"
	"fmt"
	"time"
)

type FilestoreOperationWaiter struct {
	Config  *Config
	Project string
	CommonOperationWaiter
}

func (w *FilestoreOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	// Returns the proper get.
	url := fmt.Sprintf("https://file.googleapis.com/v1beta1/%s", w.CommonOperationWaiter.Op.Name)
	return sendRequest(w.Config, "GET", w.Project, url, nil)
}

func createFilestoreWaiter(config *Config, op map[string]interface{}, project, activity string) (*FilestoreOperationWaiter, error) {
	if val, ok := op["name"]; !ok || val == "" {
		// This was a synchronous call - there is no operation to wait for.
		return nil, nil
	}
	w := &FilestoreOperationWaiter{
		Config:  config,
		Project: project,
	}
	if err := w.CommonOperationWaiter.SetOp(op); err != nil {
		return nil, err
	}
	return w, nil
}

// nolint: deadcode,unused
func filestoreOperationWaitTimeWithResponse(config *Config, op map[string]interface{}, response *map[string]interface{}, project, activity string, timeout time.Duration) error {
	w, err := createFilestoreWaiter(config, op, project, activity)
	if err != nil || w == nil {
		// If w is nil, the op was synchronous.
		return err
	}
	if err := OperationWait(w, activity, timeout, config.PollInterval); err != nil {
		return err
	}
	return json.Unmarshal([]byte(w.CommonOperationWaiter.Op.Response), response)
}

func filestoreOperationWaitTime(config *Config, op map[string]interface{}, project, activity string, timeout time.Duration) error {
	w, err := createFilestoreWaiter(config, op, project, activity)
	if err != nil || w == nil {
		// If w is nil, the op was synchronous.
		return err
	}
	return OperationWait(w, activity, timeout, config.PollInterval)
}
