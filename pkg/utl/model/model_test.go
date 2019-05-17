package model_test

import (
	"testing"

	"gitlab.com/nguyencatpham/go-effective-study/pkg/utl/mock"
	"gitlab.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

func TestBeforeInsert(t *testing.T) {
	base := &model.Base{
		ID: 1,
	}
	base.BeforeInsert(nil)
	if base.CreatedAt.IsZero() {
		t.Error("CreatedAt was not changed")
	}
	if base.UpdatedAt.IsZero() {
		t.Error("UpdatedAt was not changed")
	}
}

func TestBeforeUpdate(t *testing.T) {
	base := &model.Base{
		ID:        1,
		CreatedAt: mock.TestTime(2000),
	}
	base.BeforeUpdate(nil)
	if base.UpdatedAt == mock.TestTime(2001) {
		t.Error("UpdatedAt was not changed")
	}

}

func TestPaginationTransform(t *testing.T) {
	p := &model.PaginationReq{
		Limit: 5000, Page: 5,
	}

	pag := p.Transform()

	if pag.Limit != 1000 {
		t.Error("Default limit not set")
	}

	if pag.Offset != 5000 {
		t.Error("Offset not set correctly")
	}

	p.Limit = 0
	newPag := p.Transform()

	if newPag.Limit != 100 {
		t.Error("Min limit not set")
	}

}
