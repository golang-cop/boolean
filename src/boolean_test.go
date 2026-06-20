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
	ginkgo.It("can inspect its type", func() {
		gomega.Expect(Boolean.True().Inspect().Type().ToGoString()).To(gomega.ContainSubstring("Boolean"))
	})
	ginkgo.It("can inspect its address", func() {
		gomega.Expect(Boolean.True().Inspect().Addr().ToGoString()).To(gomega.ContainSubstring("0x"))
	})
	ginkgo.It("can inspect its data", func() {
		gomega.Expect(Boolean.True().Inspect().Data().Payload()).To(gomega.ContainSubstring("value"))
	})
})
