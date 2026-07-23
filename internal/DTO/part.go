package DTO

type Part struct {
	Id                string  `json:"id"`
	Name              string  `json:"name"`
	Category          string  `json:"category"`
	CurrentStock      int     `json:"currentStock"`
	MinimumStock      int     `json:"minimumStock"`
	AverageDailySales int     `json:"averageDailySales"`
	LeadTimeDays      int     `json:"leadTimeDays"`
	CriticalityLevel  int     `json:"criticalityLevel"`
	UnitCost          float64 `json:"unitCost"`
}
