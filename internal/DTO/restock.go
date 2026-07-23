package DTO

type Priorities struct {
	Id             string `json:"partId"`
	Name           string `json:"name"`
	CurrentStock   int    `json:"currentStock"`
	ProjectedStock int    `json:"projectedStock"`
	MinimumStock   int    `json:"minimumStock"`
	UrgencyStock   int    `json:"urgencyStock"`
}

type RestockResponse struct {
	Priorities []Priorities `json:"priorities"`
}
