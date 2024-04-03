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
	case ASC:
		return "ASC"
	case DESC:
		return "DESC"
	default:
		return ""
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
	return s.
		Where(squirrel.ILike{p.Filter.Column: fmt.Sprintf("%%%s%%", p.Filter.Pattern)}).
		OrderBy(p.Sort.Format()).
		Limit(p.PageSize).
		Offset(p.PageNumber * p.PageSize)
}
