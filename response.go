package omnicron

import "encoding/json"

type ResponseMsg struct {
	Response interface{} `json:"response"`
}

func unmarshalJSONResponse(body []byte) (*ResponseMsg, error) {
	var responseParams ResponseMsg
	if err := json.Unmarshal(body, &responseParams); err != nil {
		return nil, err
	}
	return &responseParams, nil
}
