package repository

import (
	"context"
	"database/sql"

	domainpart "github.com/ThalesMonteir0/backend-test/internal/domain/part"
)

type Parts struct {
	db *sql.DB
}

func New(db *sql.DB) *Parts {
	return &Parts{db: db}
}

func (r *Parts) Create(ctx context.Context, p domainpart.Part) (domainpart.Part, error) {
	const query = `
		INSERT INTO parts
			(name, category, current_stock, minimum_stock, average_daily_sales, lead_time_days, criticality_level, unit_cost)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	err := r.db.QueryRowContext(ctx, query,
		p.Name,
		p.Category,
		p.CurrentStock,
		p.MinimumStock,
		p.AverageDailySales,
		p.LeadTimeDays,
		p.CriticalityLevel,
		p.UnitCoast,
	).Scan(&p.ID)
	if err != nil {
		return domainpart.Part{}, err
	}

	return p, nil
}

func (r *Parts) GetAll(ctx context.Context) ([]domainpart.Part, error) {
	const query = `
		SELECT id, name, category, current_stock, minimum_stock, average_daily_sales, lead_time_days, criticality_level, unit_cost
		FROM parts`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	parts := make([]domainpart.Part, 0)
	for rows.Next() {
		var p domainpart.Part
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Category,
			&p.CurrentStock,
			&p.MinimumStock,
			&p.AverageDailySales,
			&p.LeadTimeDays,
			&p.CriticalityLevel,
			&p.UnitCoast,
		); err != nil {
			return nil, err
		}
		parts = append(parts, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return parts, nil
}

func (r *Parts) Update(ctx context.Context, p domainpart.Part) (domainpart.Part, error) {
	const query = `
		UPDATE parts
		SET name = $1,
			category = $2,
			current_stock = $3,
			minimum_stock = $4,
			average_daily_sales = $5,
			lead_time_days = $6,
			criticality_level = $7,
			unit_cost = $8
		WHERE id = $9`

	result, err := r.db.ExecContext(ctx, query,
		p.Name,
		p.Category,
		p.CurrentStock,
		p.MinimumStock,
		p.AverageDailySales,
		p.LeadTimeDays,
		p.CriticalityLevel,
		p.UnitCoast,
		p.ID,
	)
	if err != nil {
		return domainpart.Part{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domainpart.Part{}, err
	}
	if rowsAffected == 0 {
		return domainpart.Part{}, sql.ErrNoRows
	}

	return p, nil
}

func (r *Parts) Delete(ctx context.Context, id string) error {
	const query = `DELETE FROM parts WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
