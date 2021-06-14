package nanocoinrpc

import (
	"errors"
)

type BootsrapRequest struct {
	Action                     string `json:"action"`
	Address                    string `json:"address"`
	Port                       string `json:"port"`
	BypassFrontierConfirmation string `json:"bypass_frontier_confirmation,omitempty"`
	ID                         string `json:"id,omitempty"`
}

type BootsrapResponse struct {
	Success string `json:"success"`

	Error string `json:"error"`
}

func (n *NanoService) Bootsrap(r BootsrapRequest) (*BootsrapResponse, error) {
	r.Action = "bootsrap"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &BootsrapResponse{}
	_, err = n.client.Do(req, jsonResp)
	if err != nil {
		return nil, err
	}

	if jsonResp.Error != "" {
		jsonErr := errors.New("Error from nano server: " + jsonResp.Error)
		return jsonResp, jsonErr
	}

	return jsonResp, nil
}

type BootsrapAnyRequest struct {
	Action  string `json:"action"`
	Force   string `json:"force,omitempty"`
	ID      string `json:"id,omitempty"`
	Account string `json:"account,omitempty"`
}

type BootsrapAnyResponse struct {
	Success string `json:"success"`

	Error string `json:"error"`
}

func (n *NanoService) BootsrapAny(r BootsrapAnyRequest) (*BootsrapAnyResponse, error) {
	r.Action = "bootsrap_any"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &BootsrapAnyResponse{}
	_, err = n.client.Do(req, jsonResp)
	if err != nil {
		return nil, err
	}

	if jsonResp.Error != "" {
		jsonErr := errors.New("Error from nano server: " + jsonResp.Error)
		return jsonResp, jsonErr
	}

	return jsonResp, nil
}

type BootsrapLazyRequest struct {
	Action string `json:"action"`
	Hash   string `json:"hash"`
	Force  string `json:"force,omitempty"`
	ID     string `json:"id,omitempty"`
}

type BootsrapLazyResponse struct {
	Started     string `json:"started"`
	KeyInserted string `json:"key_inserted"`

	Error string `json:"error"`
}

func (n *NanoService) BootsrapLazy(r BootsrapLazyRequest) (*BootsrapLazyResponse, error) {
	r.Action = "bootsrap_lazy"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &BootsrapLazyResponse{}
	_, err = n.client.Do(req, jsonResp)
	if err != nil {
		return nil, err
	}

	if jsonResp.Error != "" {
		jsonErr := errors.New("Error from nano server: " + jsonResp.Error)
		return jsonResp, jsonErr
	}

	return jsonResp, nil
}

type BootsrapStatusRequest struct {
	Action string `json:"action"`
}

type BootsrapStatusResponse struct {
	BootstrapThreads     string `json:"bootstrap_threads"`
	RunningAttemptsCount string `json:"running_attempts_count"`
	TotalAttemptsCount   string `json:"total_attempts_count"`
	Connections          struct {
		Clients           string `json:"clients"`
		Connections       string `json:"connections"`
		Idle              string `json:"idle"`
		TargetConnections string `json:"target_connections"`
		Pulls             string `json:"pulls"`
	} `json:"connections"`
	Attempts []struct {
		ID            string `json:"id"`
		Mode          string `json:"mode"`
		Started       string `json:"started"`
		Pulling       string `json:"pulling"`
		TotalBlocks   string `json:"total_blocks"`
		RequeuedPulls string `json:"requeued_pulls"`
		Duration      string `json:"duration"`

		FrontierPulls               string `json:"frontier_pulls"`
		FrontierReceived            string `json:"frontiers_received"`
		FrontierConfirmed           string `json:"frontiers_confirmed"`
		FrontierConfirmationPending string `json:"frontiers_confirmation_pending"`
		FrontierAge                 string `json:"frontiers_age"`
		LastAccount                 string `json:"last_account"`

		LazyBlocks         string `json:"lazy_blocks"`
		LazyStateBacklog   string `json:"lazy_state_backlog"`
		LazyBalances       string `json:"lazy_balances"`
		LazyDestinations   string `json:"lazy_destinations"`
		LazyUndefinedLinks string `json:"lazy_undefined_links"`
		LazyPulls          string `json:"lazy_pulls"`
		LazyKeys           string `json:"lazy_keys"`
		LazyKey1           string `json:"lazy_key_1"`
	} `json:"attempts"`

	Error string `json:"error"`
}

func (n *NanoService) BootsrapStatus(r BootsrapStatusRequest) (*BootsrapStatusResponse, error) {
	r.Action = "bootsrap_status"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &BootsrapStatusResponse{}
	_, err = n.client.Do(req, jsonResp)
	if err != nil {
		return nil, err
	}

	if jsonResp.Error != "" {
		jsonErr := errors.New("Error from nano server: " + jsonResp.Error)
		return jsonResp, jsonErr
	}

	return jsonResp, nil
}
