package model

const (
	PageDefault = 0
	SizeDefault = 10
	MaxSize     = 100
)

type EmployeesResponse struct {
	PageResponse
	Items []Employee
}

type PageResponse struct {
	Page       int64 `json:"page"`
	TotalItems int64 `json:"total_items"`
	TotalPages int64 `json:"total_pages"`
}

type PageRequest struct {
	Page int64
	Size int64
}

func (p *PageRequest) WithDefaultValues() *PageRequest {
	if p.Page == 0 {
		p.Page = PageDefault
	}

	if p.Size == 0 {
		p.Size = SizeDefault
	} else if p.Size > MaxSize {
		p.Size = MaxSize
	}

	return &PageRequest{
		Page: p.Page, Size: p.Size,
	}
}
