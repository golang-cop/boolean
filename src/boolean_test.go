package Boolean_test

import (
	Boolean "github.com/go-composites/boolean/src"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Boolean", func() {

	//BeforeEach(func(){
	//
	//	})
	ginkgo.It("constructor can create a new Boolean from Go native type value true", func() {
		gomega.Expect(Boolean.New(true).ToGoBool()).To(gomega.BeTrue())
	})
	ginkgo.It("constructor can create a new Boolean from Go native type value false", func() {
		gomega.Expect(Boolean.New(false).ToGoBool()).To(gomega.BeFalse())
	})
	ginkgo.It("can create a new True Boolean", func() {
		gomega.Expect(Boolean.True()).To(gomega.Equal(Boolean.True()))
	})
	ginkgo.It("can create a new False Boolean", func() {
		gomega.Expect(Boolean.False()).To(gomega.Equal(Boolean.False()))
	})
	ginkgo.It("can tell if a true Boolean value is true", func() {
		gomega.Expect(Boolean.True().IsTrue()).To(gomega.BeTrue())
	})
	ginkgo.It("can tell if a true Boolean value is false", func() {
		gomega.Expect(Boolean.True().IsFalse()).To(gomega.BeFalse())
	})
	ginkgo.It("can tell if a false Boolean value is true", func() {
		gomega.Expect(Boolean.False().IsTrue()).To(gomega.BeFalse())
	})
	ginkgo.It("can tell if a false Boolean value is false", func() {
		gomega.Expect(Boolean.False().IsFalse()).To(gomega.BeTrue())
	})
	ginkgo.It("renders its Go string form", func() {
		gomega.Expect(Boolean.True().ToGoString()).To(gomega.Equal("\"true\""))
		gomega.Expect(Boolean.False().ToGoString()).To(gomega.Equal("\"false\""))
	})
	ginkgo.It("can tell whether it equals another Boolean", func() {
		gomega.Expect(Boolean.True().Equal(Boolean.True()).IsTrue()).To(gomega.BeTrue())
		gomega.Expect(Boolean.True().Equal(Boolean.False()).IsTrue()).To(gomega.BeFalse())
		gomega.Expect(Boolean.False().Equal(Boolean.False()).IsTrue()).To(gomega.BeTrue())
	})
	ginkgo.It("can inspect its type", func() {
		gomega.Expect(Boolean.True().Inspect().Type().ToGoString()).To(gomega.ContainSubstring("Boolean"))
	})
	ginkgo.It("can inspect its address", func() {
		gomega.Expect(Boolean.True().Inspect().Addr().ToGoString()).To(gomega.ContainSubstring("0x"))
	})
	ginkgo.It("can inspect its data", func() {
		gomega.Expect(Boolean.True().Inspect().Data().Payload()).To(gomega.ContainSubstring("value"))
	})
	ginkgo.It("computes logical AND over the truth table", func() {
		gomega.Expect(Boolean.True().And(Boolean.True()).ToGoBool()).To(gomega.BeTrue())
		gomega.Expect(Boolean.True().And(Boolean.False()).ToGoBool()).To(gomega.BeFalse())
		gomega.Expect(Boolean.False().And(Boolean.True()).ToGoBool()).To(gomega.BeFalse())
		gomega.Expect(Boolean.False().And(Boolean.False()).ToGoBool()).To(gomega.BeFalse())
	})
	ginkgo.It("computes logical OR over the truth table", func() {
		gomega.Expect(Boolean.True().Or(Boolean.True()).ToGoBool()).To(gomega.BeTrue())
		gomega.Expect(Boolean.True().Or(Boolean.False()).ToGoBool()).To(gomega.BeTrue())
		gomega.Expect(Boolean.False().Or(Boolean.True()).ToGoBool()).To(gomega.BeTrue())
		gomega.Expect(Boolean.False().Or(Boolean.False()).ToGoBool()).To(gomega.BeFalse())
	})
	ginkgo.It("computes logical NOT", func() {
		gomega.Expect(Boolean.True().Not().ToGoBool()).To(gomega.BeFalse())
		gomega.Expect(Boolean.False().Not().ToGoBool()).To(gomega.BeTrue())
	})
	ginkgo.It("computes exclusive OR over the truth table", func() {
		gomega.Expect(Boolean.True().Xor(Boolean.True()).ToGoBool()).To(gomega.BeFalse())
		gomega.Expect(Boolean.True().Xor(Boolean.False()).ToGoBool()).To(gomega.BeTrue())
		gomega.Expect(Boolean.False().Xor(Boolean.True()).ToGoBool()).To(gomega.BeTrue())
		gomega.Expect(Boolean.False().Xor(Boolean.False()).ToGoBool()).To(gomega.BeFalse())
	})
	ginkgo.It("reports that a real Boolean is not null", func() {
		gomega.Expect(Boolean.True().IsNull()).To(gomega.BeFalse())
		gomega.Expect(Boolean.False().IsNull()).To(gomega.BeFalse())
	})

	ginkgo.Context("the Null-Object Boolean", func() {
		ginkgo.It("reports that it is null", func() {
			gomega.Expect(Boolean.Null().IsNull()).To(gomega.BeTrue())
		})
		ginkgo.It("is neither true nor false", func() {
			gomega.Expect(Boolean.Null().IsTrue()).To(gomega.BeFalse())
			gomega.Expect(Boolean.Null().IsFalse()).To(gomega.BeFalse())
		})
		ginkgo.It("renders as the null literal", func() {
			gomega.Expect(Boolean.Null().ToGoString()).To(gomega.Equal("\"null\""))
		})
		ginkgo.It("has a false Go bool value", func() {
			gomega.Expect(Boolean.Null().ToGoBool()).To(gomega.BeFalse())
		})
		ginkgo.It("returns the null Boolean from Equal", func() {
			gomega.Expect(Boolean.Null().Equal(Boolean.True()).IsNull()).To(gomega.BeTrue())
		})
		ginkgo.It("returns the null Boolean from And", func() {
			gomega.Expect(Boolean.Null().And(Boolean.True()).IsNull()).To(gomega.BeTrue())
		})
		ginkgo.It("returns the null Boolean from Or", func() {
			gomega.Expect(Boolean.Null().Or(Boolean.True()).IsNull()).To(gomega.BeTrue())
		})
		ginkgo.It("returns the null Boolean from Not", func() {
			gomega.Expect(Boolean.Null().Not().IsNull()).To(gomega.BeTrue())
		})
		ginkgo.It("returns the null Boolean from Xor", func() {
			gomega.Expect(Boolean.Null().Xor(Boolean.True()).IsNull()).To(gomega.BeTrue())
		})
		ginkgo.It("can inspect its data", func() {
			gomega.Expect(Boolean.Null().Inspect().Data().Payload()).To(gomega.ContainSubstring("null"))
		})
	})
})
