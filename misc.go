package nanocoinrpc

import (
	"errors"
)

type AvailableSupplyRequest struct {
	Action string `json:"action"`
}

type AvailableSupplyResponse struct {
	Available string `json:"available"`

	Error string `json:"error"`
}

func (n *NanoService) AvailableSupply(r AvailableSupplyRequest) (*AvailableSupplyResponse, error) {
	r.Action = "available_supply"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &AvailableSupplyResponse{}
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

type ChainRequest struct {
	Action  string `json:"action"`
	Block   string `json:"block"`
	Count   string `json:"count"`
	Offset  string `json:"offset,omitempty"`
	Reverse string `json:"reverse,omitempty"`
}

type ChainResponse struct {
	Blocks []string `json:"blocks"`

	Error string `json:"error"`
}

func (n *NanoService) Chain(r ChainRequest) (*ChainResponse, error) {
	r.Action = "chain"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &ChainResponse{}
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
