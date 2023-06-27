package responses

import "tonothan/sangaku-pipeline-server/models"

type ManifestResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"body"`
}

type ManifestContentResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    models.Manifest `json:"body"`
}
