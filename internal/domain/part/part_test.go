package part

import (
	"reflect"
	"testing"
)

func TestExpectedConsumption(t *testing.T) {
	tests := []struct {
		name string
		part Part
		want int
	}{
		{
			name: "consumo padrão",
			part: Part{AverageDailySales: 4, LeadTimeDays: 5},
			want: 20,
		},
		{
			name: "venda zero resulta em consumo zero",
			part: Part{AverageDailySales: 0, LeadTimeDays: 10},
			want: 0,
		},
		{
			name: "lead time alto",
			part: Part{AverageDailySales: 3, LeadTimeDays: 365},
			want: 1095,
		},
		{
			name: "lead time zero resulta em consumo zero",
			part: Part{AverageDailySales: 8, LeadTimeDays: 0},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.part.ExpectedConsumption(); got != tt.want {
				t.Errorf("ExpectedConsumption() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestProjectedStock(t *testing.T) {
	tests := []struct {
		name string
		part Part
		want int
	}{
		{
			name: "estoque projetado positivo",
			part: Part{CurrentStock: 100, AverageDailySales: 4, LeadTimeDays: 5},
			want: 80,
		},
		{
			name: "estoque projetado negativo (consumo maior que estoque)",
			part: Part{CurrentStock: 15, AverageDailySales: 4, LeadTimeDays: 5},
			want: -5,
		},
		{
			name: "estoque zerado com lead time alto",
			part: Part{CurrentStock: 10, AverageDailySales: 5, LeadTimeDays: 100},
			want: -490,
		},
		{
			name: "venda zero mantém estoque atual",
			part: Part{CurrentStock: 30, AverageDailySales: 0, LeadTimeDays: 50},
			want: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.part.ProjectedStock(); got != tt.want {
				t.Errorf("ProjectedStock() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestNeedRestock(t *testing.T) {
	tests := []struct {
		name string
		part Part
		want bool
	}{
		{
			name: "precisa de reposição quando projetado abaixo do mínimo",
			part: Part{CurrentStock: 15, MinimumStock: 20, AverageDailySales: 4, LeadTimeDays: 5},
			want: true,
		},
		{
			name: "não precisa quando projetado acima do mínimo",
			part: Part{CurrentStock: 100, MinimumStock: 20, AverageDailySales: 4, LeadTimeDays: 5},
			want: false,
		},
		{
			name: "não precisa quando projetado igual ao mínimo",
			part: Part{CurrentStock: 40, MinimumStock: 20, AverageDailySales: 4, LeadTimeDays: 5},
			want: false,
		},
		{
			name: "precisa quando estoque projetado é negativo",
			part: Part{CurrentStock: 8, MinimumStock: 10, AverageDailySales: 2, LeadTimeDays: 5},
			want: true,
		},
		{
			name: "venda zero e estoque acima do mínimo não precisa",
			part: Part{CurrentStock: 25, MinimumStock: 20, AverageDailySales: 0, LeadTimeDays: 100},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.part.NeedRestock(); got != tt.want {
				t.Errorf("NeedRestock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrgencyScore(t *testing.T) {
	tests := []struct {
		name string
		part Part
		want int
	}{
		{
			name: "score padrão",
			part: Part{CurrentStock: 15, MinimumStock: 20, AverageDailySales: 4, LeadTimeDays: 5, CriticalityLevel: 3},
			want: 75, // (20 - (-5)) * 3
		},
		{
			name: "estoque projetado negativo aumenta o score",
			part: Part{CurrentStock: 8, MinimumStock: 10, AverageDailySales: 4, LeadTimeDays: 5, CriticalityLevel: 3},
			want: 66, // (10 - (-12)) * 3
		},
		{
			name: "criticidade máxima com lead time alto",
			part: Part{CurrentStock: 0, MinimumStock: 5, AverageDailySales: 2, LeadTimeDays: 100, CriticalityLevel: 5},
			want: 1025, // (5 - (-200)) * 5
		},
		{
			name: "venda zero mantém score baixo",
			part: Part{CurrentStock: 5, MinimumStock: 20, AverageDailySales: 0, LeadTimeDays: 100, CriticalityLevel: 2},
			want: 30, // (20 - 5) * 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.part.UrgencyScore(); got != tt.want {
				t.Errorf("UrgencyScore() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestPrioritize(t *testing.T) {
	tests := []struct {
		name  string
		parts []Part
		want  []string // ordem esperada dos nomes
	}{
		{
			name: "filtra peças que não precisam de reposição",
			parts: []Part{
				{Name: "Não precisa", CurrentStock: 100, MinimumStock: 20, AverageDailySales: 1, LeadTimeDays: 1, CriticalityLevel: 3},
				{Name: "Precisa", CurrentStock: 5, MinimumStock: 20, AverageDailySales: 2, LeadTimeDays: 3, CriticalityLevel: 2},
			},
			want: []string{"Precisa"},
		},
		{
			name: "ordena por urgencyScore decrescente",
			parts: []Part{
				{Name: "Baixo", CurrentStock: 18, MinimumStock: 20, AverageDailySales: 1, LeadTimeDays: 2, CriticalityLevel: 1},
				{Name: "Alto", CurrentStock: 0, MinimumStock: 20, AverageDailySales: 5, LeadTimeDays: 4, CriticalityLevel: 5},
				{Name: "Médio", CurrentStock: 10, MinimumStock: 20, AverageDailySales: 2, LeadTimeDays: 3, CriticalityLevel: 2},
			},
			want: []string{"Alto", "Médio", "Baixo"},
		},
		{
			name: "desempate por criticalityLevel",
			parts: []Part{
				// mesmo urgencyScore = 20, mas criticidades diferentes
				{Name: "Crit baixa", CurrentStock: 0, MinimumStock: 10, AverageDailySales: 0, LeadTimeDays: 0, CriticalityLevel: 2}, // (10-0)*2=20
				{Name: "Crit alta", CurrentStock: 15, MinimumStock: 20, AverageDailySales: 0, LeadTimeDays: 0, CriticalityLevel: 4}, // (20-15)*4=20
			},
			want: []string{"Crit alta", "Crit baixa"},
		},
		{
			name: "desempate por averageDailySales quando score e criticidade empatam",
			parts: []Part{
				// score = 20, criticidade = 2, vendas diferentes
				{Name: "Venda baixa", CurrentStock: 0, MinimumStock: 10, AverageDailySales: 0, LeadTimeDays: 0, CriticalityLevel: 2},
				{Name: "Venda alta", CurrentStock: 10, MinimumStock: 20, AverageDailySales: 10, LeadTimeDays: 0, CriticalityLevel: 2},
			},
			want: []string{"Venda alta", "Venda baixa"},
		},
		{
			name: "desempate final por ordem alfabética",
			parts: []Part{
				// score, criticidade e vendas idênticos
				{Name: "Zebra", CurrentStock: 0, MinimumStock: 10, AverageDailySales: 0, LeadTimeDays: 0, CriticalityLevel: 2},
				{Name: "Alpha", CurrentStock: 0, MinimumStock: 10, AverageDailySales: 0, LeadTimeDays: 0, CriticalityLevel: 2},
				{Name: "Mango", CurrentStock: 0, MinimumStock: 10, AverageDailySales: 0, LeadTimeDays: 0, CriticalityLevel: 2},
			},
			want: []string{"Alpha", "Mango", "Zebra"},
		},
		{
			name:  "lista vazia retorna vazio",
			parts: []Part{},
			want:  []string{},
		},
		{
			name: "nenhuma peça precisa de reposição",
			parts: []Part{
				{Name: "A", CurrentStock: 100, MinimumStock: 10, AverageDailySales: 1, LeadTimeDays: 1, CriticalityLevel: 3},
				{Name: "B", CurrentStock: 200, MinimumStock: 20, AverageDailySales: 2, LeadTimeDays: 2, CriticalityLevel: 4},
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Prioritize(tt.parts)

			gotNames := make([]string, 0, len(got))
			for _, p := range got {
				gotNames = append(gotNames, p.Name)
			}

			if !reflect.DeepEqual(gotNames, tt.want) {
				t.Errorf("Prioritize() ordem = %v, want %v", gotNames, tt.want)
			}
		})
	}
}
