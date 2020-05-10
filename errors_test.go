package godfatherr

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestError(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "The Error tests")
}

var _ = Describe("Empty Error tests", func() {
	When("empty Error is created", func() {
		err := Empty()
		It("declares as not present", func() {
			Expect(err.IsPresent()).To(BeFalse())
		})
		It("contains wrapped error which is nil", func() {
			Expect(err.Error()).To(BeNil())
		})
		It("contains an empty message", func() {
			Expect(err.Message()).To(BeEmpty())
		})
		It("is not marked fatal", func() {
			Expect(err.IsFatal()).To(BeFalse())
		})
		It("contains no context", func() {
			Expect(err.ctx).To(BeEmpty())
		})
		It("displays itself as '[empty err]'", func() {
			Expect(err.String()).To(Equal("[empty err]"))
		})
	})
	When("empty Error exists", func() {
		err := Empty()
		It("ignores the context assignment", func() {
			err.WithCtx("test")
			Expect(err.ctx).To(BeEmpty())
		})
		It("ignores the panic call", func() {
			err.Panic()
		})
		It("ignores the watch command", func() {
			err.Watch()
			Expect(err.isWatched()).To(BeFalse())
		})
		It("ignores the unwatch command", func() {
			unwatched := err.unwatch()
			Expect(err).To(Equal(unwatched))
		})
	})
})
