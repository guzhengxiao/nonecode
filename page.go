package nonecode

import (
	"gorm.io/gorm"
)

type PageParam struct {
	Page   int `json:"page"`
	Size   int `json:"size"`
	Offset int `json:"-"`
}

func (p *PageParam) Def() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Size <= 0 || p.Size > 100 {
		p.Size = 20
	}
	p.Offset = (p.Page - 1) * p.Size
}

func (p *PageParam) Do(stmt *gorm.DB, resModel interface{}) *PageResponse {
	p.Def()
	rs := &PageResponse{PageParam: p}
	stmt.Count(&rs.Count)
	stmt.Limit(p.Size).Offset(p.Offset).Find(resModel)
	rs.Data = resModel
	return rs
}

type PageResponse struct {
	*PageParam
	Count int64       `json:"count"`
	Data  interface{} `json:"data"`
}
