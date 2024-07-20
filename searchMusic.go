package omnicron

import "context"

type MusicSearchRequest struct {
	Song  string `json:"song"`
	Limit int    `json:"limit,omitempty"`
	Proxy string `json:"proxy,omitempty"`
}
func (c *Client) MusicSearch(ctx context.Context, req *MusicSearchRequest)(*GabsContainer, error){
	if req.Song == ""{
		return nil, ErrSongNotProvided
	}
	body, err := c.newJSONPostRequest(ctx, "/musicsearch", "", req)
	if err!= nil{
        return nil, err
    }
	musicSearchResponse, err := unmarshalJSONResponse(body)
	if err != nil {
		return nil, err
	}
	return musicSearchResponse, nil
}