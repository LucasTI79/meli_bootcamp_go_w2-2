package sections

import "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"

type SectionResponse struct {
	Data domain.Section `json:"data"`
}
