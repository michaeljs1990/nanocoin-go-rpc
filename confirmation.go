package nanocoinrpc

import (
	"errors"
)

type ConfirmationActiveRequest struct {
	Action        string `json:"action"`
	Announcements string `json:"announcements,omitempty"`
}

type ConfirmationActiveResponse struct {
	Confirmations []string `json:"confirmations"`
	Unconfirmed   string   `json:"unconfirmed"`
	Confirmated   string   `json:"confirmed"`

	Error string `json:"error"`
}

func (n *NanoService) ConfirmationActive(r ConfirmationActiveRequest) (*ConfirmationActiveResponse, error) {
	r.Action = "confirmation_active"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &ConfirmationActiveResponse{}
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

type ConfirmationHeightCurrentlyProcessingRequest struct {
	Action string `json:"action"`
}

type ConfirmationHeightCurrentlyProcessingResponse struct {
	Hash string `json:"hash"`

	Error string `json:"error"`
}

func (n *NanoService) ConfirmationHeightCurrentlyProcessing(r ConfirmationHeightCurrentlyProcessingRequest) (
	*ConfirmationHeightCurrentlyProcessingResponse, error) {

	r.Action = "confirmation_height_currently_processing"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &ConfirmationHeightCurrentlyProcessingResponse{}
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
