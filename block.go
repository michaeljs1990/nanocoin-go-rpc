package nanocoinrpc

import (
	"errors"
)

type BlockAccountRequest struct {
	Action string `json:"action"`
	Hash   string `json:"hash"`
}

type BlockAccountResponse struct {
	Account string `json:"account"`

	Error string `json:"error"`
}

func (n *NanoService) BlockAccount(r BlockAccountRequest) (*BlockAccountResponse, error) {
	r.Action = "block_account"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &BlockAccountResponse{}
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

type BlockConfirmRequest struct {
	Action string `json:"action"`
	Hash   string `json:"hash"`
}

type BlockConfirmResponse struct {
	Started string `json:"started"`

	Error string `json:"error"`
}

func (n *NanoService) BlockConfirm(r BlockConfirmRequest) (*BlockConfirmResponse, error) {
	r.Action = "block_confirm"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &BlockConfirmResponse{}
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

type BlockCountRequest struct {
	Action string `json:"action"`
}

type BlockCountResponse struct {
	Count     string `json:"count"`
	Unchecked string `json:"unchecked"`
	Cemented  string `json:"cemented"`

	Error string `json:"error"`
}

func (n *NanoService) BlockCount(r BlockCountRequest) (*BlockCountResponse, error) {
	r.Action = "block_count"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &BlockCountResponse{}
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

type BlockCreateRequest struct {
	Action         string `json:"action"`
	JSONBlock      string `json:"json_block,omitempty"`
	Type           string `json:"type"`
	Balance        string `json:"balance"`
	Key            string `json:"key"`
	Representative string `json:"representative"`
	Link           string `json:"link"`
	Previous       string `json:"previous"`
	Wallet         string `json:"wallet,omitempty"`
	Account        string `json:"account,omitempty"`
	Source         string `json:"source,omitempty"`
	Destination    string `json:"destination,omitempty"`
	Work           string `json:"work,omitempty"`
	Version        string `json:"version,omitempty"`
	Difficulty     string `json:"difficulty,omitempty"`
}

type BlockCreateResponse struct {
	Hash       string `json:"hash"`
	Difficulty string `json:"difficulty"`
	Block      struct {
		Type           string `json:"type"`
		Account        string `json:"account"`
		Previous       string `json:"previous"`
		Representative string `json:"representative"`
		Balance        string `json:"balance"`
		Link           string `json:"link"`
		LinkAsAccount  string `json:"link_as_account"`
		Signature      string `json:"signature"`
		Work           string `json:"work"`
	} `json:"block"`

	Error string `json:"error"`
}

func (n *NanoService) BlockCreate(r BlockCreateRequest) (*BlockCreateResponse, error) {
	r.Action = "block_create"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &BlockCreateResponse{}
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

type BlockHashRequest struct {
	Action         string `json:"action"`
	JSONBlock      string `json:"json_block,omitempty"`
	Type           string `json:"type"`
	Balance        string `json:"balance"`
	Key            string `json:"key"`
	Representative string `json:"representative"`
	Link           string `json:"link"`
	Previous       string `json:"previous"`
	Wallet         string `json:"wallet,omitempty"`
	Account        string `json:"account,omitempty"`
	Source         string `json:"source,omitempty"`
	Destination    string `json:"destination,omitempty"`
	Work           string `json:"work,omitempty"`
	Version        string `json:"version,omitempty"`
	Difficulty     string `json:"difficulty,omitempty"`
}

type BlockHashResponse struct {
	Hash       string `json:"hash"`
	Difficulty string `json:"difficulty"`
	Block      struct {
		Type           string `json:"type"`
		Account        string `json:"account"`
		Previous       string `json:"previous"`
		Representative string `json:"representative"`
		Balance        string `json:"balance"`
		Link           string `json:"link"`
		LinkAsAccount  string `json:"link_as_account"`
		Signature      string `json:"signature"`
		Work           string `json:"work"`
	} `json:"block"`

	Error string `json:"error"`
}

func (n *NanoService) BlockHash(r BlockHashRequest) (*BlockHashResponse, error) {
	r.Action = "block_hash"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &BlockHashResponse{}
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
