package productbatchesdto

type CreateProductBatchesDTO struct {
	BatchNumber        int    `json:"batch_number" binding:"required"`
	CurrentQuantity    int    `json:"current_quantity" binding:"required"`
	CurrentTemperature float64    `json:"current_temperature" binding:"required"`
	DueDate            string `json:"due_date" binding:"required"` // Example: "2023-07-06"
	InitialQuantity    int    `json:"initial_quantity" binding:"required"`
	ManufacturingDate  string `json:"manufacturing_date" binding:"required"`
	ManufacturingHour  string `json:"manufacturing_hour" binding:"required"`
	MinimumTemperature float64    `json:"minimum_temperature" binding:"required"`
	ProductID          int    `json:"product_id" binding:"required"`
	SectionID          int    `json:"section_id" binding:"required"`
}
