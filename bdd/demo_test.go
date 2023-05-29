package bdd_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ginkgo Demo", Ordered, func() {

	BeforeAll(func() {
		fmt.Println("----BeforeAll()----")
	})

	BeforeEach(func() {
		fmt.Println("----BeforeEach()----")
	})

	AfterEach(func() {
		fmt.Println("----AfterEach()----")
	})

	AfterAll(func() {
		fmt.Println("----AfterAll()----")
	})

	Describe("Scenario 1", func() {
		It("test 1", func() {
			fmt.Println("----test1()----")
			Expect("siva").To(Equal("siva"))
		})

		It("test 2", func() {
			fmt.Println("----test2()----")
			Expect("siva2").To(Equal("siva2"))
		})
	})

	Describe("Scenario 2", func() {
		It("test 3", func() {
			fmt.Println("----test3()----")
			Expect("siva").To(Equal("siva"))
		})

		It("test 4", func() {
			fmt.Println("----test4()----")
			Expect("siva2").To(Equal("siva2"))
		})
	})
})
