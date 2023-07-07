package dtos

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

type SectionResponse struct {
	Data entities.Section `json:"data"`
}
