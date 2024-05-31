package types

type Pagination struct {
	limit  int64 `default:"10"`
	offset int64 `default:"0"`
}

func NewPagination(limit, offset *int64) *Pagination {
	out := &Pagination{
		limit:  20,
		offset: 0,
	}
	if limit != nil {
		out.limit = *limit
	}
	if offset != nil {
		out.offset = *offset
	}

	return out
}

func (p *Pagination) GetLimit() int64 {
	return p.limit
}

func (p *Pagination) SetLimit(limit int64) {
	p.limit = limit
}

func (p *Pagination) GetOffset() int64 {
	return p.offset
}

func (p *Pagination) SetOffset(offset int64) {
	p.offset = offset
}
