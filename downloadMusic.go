package omnicron

import (
	"context"
	"encoding/json"
)

type MusicRequest struct {
	Song string `json:"song"`
}

// SongResponse represents the structure of the response
type MusicResponse struct {
	Response []string `json:"response"`
}

// the downloadMusic function takes a song as input, downloads the song and return the direct cloudinary url. something to note: use the search music function to get the song before using it as input. Do not use any song name directly to avoid inaccuracy.
func (c *Client) DownloadMusic(ctx context.Context, req *MusicRequest) (*MusicResponse, error) {
	if req.Song == "" {
		return nil, ErrSongNotProvided
	}
	body, err := c.newJSONPostRequest(ctx, "/downloadmusic", "", req)
	if err != nil {
		return nil, err
	}
	var response MusicResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
