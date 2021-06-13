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
