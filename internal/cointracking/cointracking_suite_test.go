package cointracking_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCointracking(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cointracking Suite")
}
