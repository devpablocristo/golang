package ginkgofw_test

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo/extensions/table"
	//. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	ggo "github.com/devpablocristo/go-concepts/3rd-party-libs/testing/ginkgo"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo Suite")
}

var _ = Describe("Person.IsChild()", func() {
	Context("when person is a chiled", func() {
		BeforeEach(func() {
			log.Print("Is a child")
		})
		It("resturns true", func() {
			// init
			person := ggo.Person{Age: 10}

			// execution
			response := person.IsChild()

			// validation (gomega)
			Expect(response).To(BeTrue())
		})
	})

	Context("when person is not a chiled", func() {
		BeforeEach(func() {
			log.Print("Not a child")
		})

		// Skip blocl
		// XIt("resturns true", func() {
		// The block bedore was skipped
		// FIt("resturns true", func() {
		It("resturns true", func() {
			// init
			person := ggo.Person{Age: 40}

			// execution
			response := person.IsChild()

			// validation (gomega)
			Expect(response).To(BeFalse())
		})

	})

	DescribeTable("isChild table test",
		func(age int, expectedResponse bool) {
			p := ggo.Person{Age: age}

			Expect(p.IsChild()).To(Equal(expectedResponse))
		},
		Entry("when is a child", 10, true),
		Entry("when is a child", 18, false))

})
