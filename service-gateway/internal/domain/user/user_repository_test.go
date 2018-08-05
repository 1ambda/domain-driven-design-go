package user_test

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/user"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/test"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserRepository", func() {

	var db *gorm.DB
	var repo user.Repository

	BeforeEach(func() {
		db = test.GetTestDatabase(false)
		repo = user.NewRepository(db)
	})

	AfterEach(func() {
	})

	Describe("Delete()", func() {
		When("trying to delete non-existing record", func() {
			It("should return not found exception", func() {
				record, ex := repo.DeleteUser(0)

				Expect(record).To(BeFalse())
				Expect(ex).NotTo(BeNil())
				Expect(ex.IsNotFoundException()).To(BeTrue())
			})
		})
	})

	Describe("Find()", func() {
		When("the record dose not exist", func() {
			It("should return NotFoundException", func() {
				invalidId := uint(0)

				record, ex := repo.FindUserById(invalidId)
				Expect(record).To(BeNil())
				Expect(ex).NotTo(BeNil())
				Expect(ex.IsNotFoundException()).To(BeTrue())

			})
		})
	})

	Describe("CreateAuthIdentity()", func() {
		When("got non-existing uid", func() {
			It("should create User and AuthIdentity", func() {
				// TODO

			})
		})

		When("got existing uid", func() {
			It("should throw exception", func() {
				// TODO
			})
		})
	})

	Describe("FindAuthIdentityByUID()", func() {
		When("got non-existing uid", func() {
			It("should create User and AuthIdentity", func() {
				// TODO

			})
		})

		When("got non-existing uid", func() {
			It("should throw UnauthorizedException", func() {
				// TODO
			})
		})
	})
})
