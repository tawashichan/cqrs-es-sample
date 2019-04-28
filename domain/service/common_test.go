package service

import (
	"cqrs-es/infra/onMemory"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitializeRepositories()

	code := m.Run()

	os.Exit(code)
}

func InitializeRepositories() {
	onMemory.InitializeUserEventRepository()
}
