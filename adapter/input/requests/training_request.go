package requests

import "github.com/LordRadamanthys/centralized-health/application/domain"

type TrainingRequest struct {
	Training []domain.TrainingDetails `json:"day"`
}
