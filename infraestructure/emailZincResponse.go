package infraestructure

import "time"

type EmailZincResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index     string    `json:"_index"`
			Type      string    `json:"_type"`
			ID        string    `json:"_id"`
			Score     float64   `json:"_score"`
			Timestamp time.Time `json:"@timestamp"`
			Source    struct {
				Body                    string `json:"Body"`
				ContentTransferEncoding string `json:"ContentTransferEncoding"`
				ContentType             string `json:"ContentType"`
				Date                    string `json:"Date"`
				From                    string `json:"From"`
				MimeVersion             string `json:"MimeVersion"`
				Subject                 string `json:"Subject"`
				To                      string `json:"To"`
				XFileName               string `json:"XFileName"`
				XFolder                 string `json:"XFolder"`
				XFrom                   string `json:"XFrom"`
				XOrigin                 string `json:"XOrigin"`
				XTo                     string `json:"XTo"`
				Xbcc                    string `json:"Xbcc"`
				Xcc                     string `json:"Xcc"`
				ID                      string `json:"_id"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
