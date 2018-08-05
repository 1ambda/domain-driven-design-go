package product_test

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/product"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/test"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProductRepository", func() {

	var db *gorm.DB
	var repo product.Repository

	GinkgoT()

	BeforeEach(func() {
		db = test.GetTestDatabase(false)
		repo = product.NewRepository(db)
	})

	AfterEach(func() {
	})

	Describe("AddCategory()", func() {
		Context("When got valid record", func() {
			It("should return the created record", func() {
				r := &product.Category{
					Path:             "/test",
					Name:             "test",
					DisplayName:      "TEST",
					Description:      "description",
					ParentCategoryID: nil,
				}
				created, ex := repo.AddCategory(r)

				Expect(ex).To(BeNil())
				Expect(created.ID).Should(BeNumerically(">", 0))
			})
		})
	})

	Describe("FindCategory()", func() {
		Context("When record does not exist", func() {
			It("should return NotFoundException", func() {
				invalidId := uint(0)
				record, ex := repo.FindCategoryById(invalidId)

				Expect(record).To(BeNil())
				Expect(ex).NotTo(BeNil())
				Expect(ex.IsNotFoundException()).To(BeTrue())
			})
		})
	})

	Describe("AddImage()", func() {
		Context("When got valid record", func() {
			It("should return the created record", func() {
				r := &product.Image{}
				record, ex := repo.AddImage(r)

				Expect(ex).To(BeNil())
				Expect(record.ID).Should(BeNumerically(">", 0))
			})
		})
	})

	Describe("FindImage()", func() {
		Context("When record does not exist", func() {
			It("should return NotFoundException", func() {
				invalidId := uint(0)
				record, ex := repo.FindImageById(invalidId)

				Expect(record).To(BeNil())
				Expect(ex).NotTo(BeNil())
				Expect(ex.IsNotFoundException()).To(BeTrue())
			})
		})
	})
})
