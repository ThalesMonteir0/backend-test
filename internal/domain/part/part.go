package part

import "sort"

type Part struct {
	ID                string
	Name              string
	Category          string
	CurrentStock      int
	MinimumStock      int
	AverageDailySales int
	LeadTimeDays      int
	CriticalityLevel  int
	UnitCoast         float64
}

// ExpectedConsumption Calcula o consumo esperado durante o lead time
func (p *Part) ExpectedConsumption() int {
	return p.AverageDailySales * p.LeadTimeDays
}

// ProjectedStock Calcula o estoque projetado
func (p *Part) ProjectedStock() int {
	return p.CurrentStock - p.ExpectedConsumption()
}

// NeedRestock Verifica se precisa de reposição
func (p *Part) NeedRestock() bool {
	return p.ProjectedStock() < p.MinimumStock
}

// UrgencyScore Calcula Score de Prioridade
func (p *Part) UrgencyScore() int {
	return (p.MinimumStock - p.ProjectedStock()) * p.CriticalityLevel
}

// Prioritize filtra as peças que precisam de reposição e as ordena por
// prioridade: maior urgencyScore primeiro. Em caso de empate, aplica os
// critérios de desempate: maior criticalityLevel, maior averageDailySales
// e, por fim, ordem alfabética pelo nome.
func Prioritize(parts []Part) []Part {
	needing := make([]Part, 0, len(parts))
	for _, p := range parts {
		if p.NeedRestock() {
			needing = append(needing, p)
		}
	}

	sort.Slice(needing, func(i, j int) bool {
		a, b := needing[i], needing[j]

		if a.UrgencyScore() != b.UrgencyScore() {
			return a.UrgencyScore() > b.UrgencyScore()
		}
		if a.CriticalityLevel != b.CriticalityLevel {
			return a.CriticalityLevel > b.CriticalityLevel
		}
		if a.AverageDailySales != b.AverageDailySales {
			return a.AverageDailySales > b.AverageDailySales
		}
		return a.Name < b.Name
	})

	return needing
}
