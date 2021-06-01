package nanocoinrpc

import (
	"errors"
)

type NanoService struct {
	client *Client
}

type AccountBalanceRequest struct {
	Action               string `json:"action"`
	Account              string `json:"account"`
	IncludeOnlyConfirmed string `json:"include_only_confirmed,omitempty"`
}

type AccountBalanceResponse struct {
	Balance string `json:"balance"`
	Pending string `json:"pending"`

	Error string `json:"error"`
}

func (n *NanoService) AccountBalance(r AccountBalanceRequest) (*AccountBalanceResponse, error) {
	r.Action = "account_balance"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &AccountBalanceResponse{}
	_, err = n.client.Do(req, jsonResp)
	if err != nil {
		return nil, err
	}

	// Everything went well but the Nano server sent us back an error. Make this feel more like golang and
	// move the json error to be returned as an error message for handling.
	if jsonResp.Error != "" {
		jsonErr := errors.New("Error from nano server: " + jsonResp.Error)
		return jsonResp, jsonErr
	}

	return jsonResp, nil
}

type AccountBlockCountRequest struct {
	Action  string `json:"action"`
	Account string `json:"account"`
}

type AccountBlockCountResponse struct {
	BlockCount string `json:"block_count"`

	Error string `json:"error"`
}

func (n *NanoService) AccountBlockCount(r AccountBlockCountRequest) (*AccountBlockCountResponse, error) {
	r.Action = "account_block_count"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &AccountBlockCountResponse{}
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

type AccountGetRequest struct {
	Action string `json:"action"`
	Key    string `json:"key"`
}

type AccountGetResponse struct {
	Account string `json:"account"`

	Error string `json:"error"`
}

func (n *NanoService) AccountGet(r AccountGetRequest) (*AccountGetResponse, error) {
	r.Action = "account_get"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &AccountGetResponse{}
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

type AccountHistoryRequest struct {
	Action        string   `json:"action"`
	Account       string   `json:"account"`
	Count         string   `json:"count"`
	Raw           bool     `json:"raw,omitempty"`
	Head          string   `json:"head,omitempty"`
	Offset        int      `json:"offset,omitempty"`
	Reverse       bool     `json:"reverse,omitempty"`
	AccountFilter []string `json:"account_filter,omitempty"`
}

type AccountHistoryResponse struct {
	Account string `json:"account"`
	History []struct {
		Type           string `json:"type"`
		Account        string `json:"account"`
		Amount         string `json:"amount"`
		LocalTimestamp string `json:"local_timestamp"`
		Height         string `json:"height"`
		Hash           string `json:"hash"`
	} `json:"history"`
	Previous string `json:"previous"`

	Error string `json:"error"`
}

func (n *NanoService) AccountHistory(r AccountHistoryRequest) (*AccountHistoryResponse, error) {
	r.Action = "account_history"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &AccountHistoryResponse{}
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

type AccountInfoRequest struct {
	Action           string `json:"action"`
	Account          string `json:"account"`
	IncludeConfirmed string `json:"include_confirmed,omitempty"`
	Representative   string `json:"representative,omitempty"`
	Weight           string `json:"weight,omitempty"`
	Pending          string `json:"pending,omitempty"`
}

type AccountInfoResponse struct {
	Frontier                   string `json:"frontier"`
	OpenBlock                  string `json:"open_block"`
	RepresentativeBlock        string `json:"representative_block"`
	Balance                    string `json:"balance"`
	ModifiedTimestamp          string `json:"modified_timestamp"`
	BlockCount                 string `json:"block_count"`
	AccountVersion             string `json:"account_version"`
	ConfirmationHeight         string `json:"confirmation_height"`
	ConfirmationHeightFrontier string `json:"confirmation_height_frontier"`
	ConfirmedBalance           string `json:"confirmed_balance"`
	ConfirmedHeight            string `json:"confirmed_height"`
	ConfirmedFrontier          string `json:"confirmed_frontier"`
	ConfirmedRepresentative    string `json:"confirmed_representative"`
	Pending                    string `json:"pending"`
	ConfirmedPending           string `json:"confirmed_pending"`
	Weight                     string `json:"weight"`

	Error string `json:"error"`
}

func (n *NanoService) AccountInfo(r AccountInfoRequest) (*AccountInfoResponse, error) {
	r.Action = "account_info"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &AccountInfoResponse{}
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

type AccountKeyRequest struct {
	Action  string `json:"action"`
	Account string `json:"account"`
}

type AccountKeyResponse struct {
	Key string `json:"key"`

	Error string `json:"error"`
}

func (n *NanoService) AccountKey(r AccountKeyRequest) (*AccountKeyResponse, error) {
	r.Action = "account_key"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &AccountKeyResponse{}
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

type AccountRepresentativeRequest struct {
	Action  string `json:"action"`
	Account string `json:"account"`
}

type AccountRepresentativeResponse struct {
	Representative string `json:"representative"`

	Error string `json:"error"`
}

func (n *NanoService) AccountRepresentative(r AccountRepresentativeRequest) (*AccountRepresentativeResponse, error) {
	r.Action = "account_representative"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &AccountRepresentativeResponse{}
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

type AccountWeightRequest struct {
	Action  string `json:"action"`
	Account string `json:"account"`
}

type AccountWeightResponse struct {
	Representative string `json:"representative"`

	Error string `json:"error"`
}

func (n *NanoService) AccountWeight(r AccountWeightRequest) (*AccountWeightResponse, error) {
	r.Action = "account_weight"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &AccountWeightResponse{}
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

type AccountsBalancesRequest struct {
	Action   string   `json:"action"`
	Accounts []string `json:"accounts"`
}

type AccountsBalancesResponse struct {
	Balances map[string]struct {
		Balance string `json:"balance"`
		Pending string `json:"pending"`
	} `json:"balances"`

	Error string `json:"error"`
}

func (n *NanoService) AccountsBalances(r AccountsBalancesRequest) (*AccountsBalancesResponse, error) {
	r.Action = "accounts_balances"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &AccountsBalancesResponse{}
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

type AccountsFrontiersRequest struct {
	Action   string   `json:"action"`
	Accounts []string `json:"accounts"`
}

type AccountsFrontiersResponse struct {
	Frontiers map[string]string `json:"frontiers"`

	Error string `json:"error"`
}

func (n *NanoService) AccountsFrontiers(r AccountsFrontiersRequest) (*AccountsFrontiersResponse, error) {
	r.Action = "accounts_frontiers"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &AccountsFrontiersResponse{}
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

type AccountsPendingRequest struct {
	Action   string   `json:"action"`
	Accounts []string `json:"accounts"`
}

type AccountsPendingResponse struct {
	Frontiers map[string]string `json:"frontiers"`

	Error string `json:"error"`
}

func (n *NanoService) AccountPending(r AccountsPendingRequest) (*AccountsPendingResponse, error) {
	r.Action = "accounts_pending"
	req, err := n.client.NewRequest(r)
	if err != nil {
		return nil, err
	}

	jsonResp := &AccountsPendingResponse{}
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
