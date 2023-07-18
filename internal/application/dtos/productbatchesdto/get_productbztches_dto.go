package productbatchesdto

import "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"

type ProductBatchResponse struct {
	Data domain.ProductBySection `json:"data"`
}
