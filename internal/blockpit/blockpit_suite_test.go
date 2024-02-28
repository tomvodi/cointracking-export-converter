package blockpit_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBlockpit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Blockpit Suite")
}
