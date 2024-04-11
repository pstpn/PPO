package postgres

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
)

type SortDirection int8

const (
	ASC SortDirection = iota
	DESC
)

func (s SortDirection) String() string {
	switch s {
	case DESC:
		return "DESC"
	default:
		return "ASC"
	}
}

type SortOptions struct {
	Direction SortDirection
	Columns   []string
}

func (s SortOptions) Format() string {
	return fmt.Sprintf("%s %s", strings.Join(s.Columns, ","), s.Direction.String())
}

type FilterOptions struct {
	Pattern string
	Column  string
}

type Pagination struct {
	PageNumber uint64
	PageSize   uint64
	Filter     FilterOptions
	Sort       SortOptions
}

func (p *Pagination) ToSQL(s squirrel.SelectBuilder) squirrel.SelectBuilder {
	if p.Filter.Column != "" {
		s.Where(squirrel.ILike{p.Filter.Column + "::text": fmt.Sprintf("%%%s%%", p.Filter.Pattern)})
	}
	return s.
		OrderBy(p.Sort.Format()).
		Limit(p.PageSize).
		Offset(p.PageNumber * p.PageSize)
}
