package DTO

type PartRestockResponse struct {
	Id             string `json:"partId"`
	Name           string `json:"name"`
	CurrentStock   int    `json:"currentStock"`
	ProjectedStock int    `json:"projectedStock"`
	MinimumStock   int    `json:"minimumStock"`
	UrgencyStock   int    `json:"urgencyStock"`
}

type Priorities struct {
	PartRestockResponse PartRestockResponse
}

type RestockResponse struct {
	Priorities []Priorities `json:"priorities"`
}
