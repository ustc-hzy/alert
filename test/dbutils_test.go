package test

import (
	"alert/core/dao"
	"testing"
)

func TestDb(t *testing.T) {
	dao.InitDB()
}
