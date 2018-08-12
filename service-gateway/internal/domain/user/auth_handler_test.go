package user_test

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/user"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/test"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AuthHandler", func() {

	var db *gorm.DB
	var repo user.Repository
	var handler user.AuthHandler
	var sessionStore sessions.Store

	BeforeEach(func() {
		db = test.GetTestDatabase(false)
		encryptor := user.NewEncryptor(0)
		repo = user.NewRepository(db)
		sessionStore = user.NewSessionStore()
		handler = user.NewAuthHandler(repo, encryptor, sessionStore)
	})

	AfterEach(func() {
	})

	Describe("Register()", func() {
		When("got empty uid or empty", func() {
			It("should return BadRequestException", func() {
				u1, ex1 := handler.Register("  ", "email", "password")

				Expect(u1).Should(BeNil())
				Expect(ex1).ShouldNot(BeNil())
				Expect(ex1.IsBadRequestException()).Should(BeTrue())

				u2, ex2 := handler.Register("uid", "email", "  ")
				Expect(u2).Should(BeNil())
				Expect(ex2).ShouldNot(BeNil())
				Expect(ex2.IsBadRequestException()).Should(BeTrue())

				u3, ex3 := handler.Register("uid", "  ", "password")
				Expect(u3).Should(BeNil())
				Expect(ex3).ShouldNot(BeNil())
				Expect(ex3.IsBadRequestException()).Should(BeTrue())
			})
		})

		When("got valid uid and password", func() {
			It("should create AuthClaim and User", func() {
				uid := "uid"
				password := "ma password"
				email := "me@email.com"
				claim, ex := handler.Register(uid, email, password)

				Expect(ex).Should(BeNil())
				Expect(claim.UserID).Should(BeNumerically(">", 0))
				Expect(claim.UID).Should(Equal(uid))
			})
		})
	})

	Describe("Login()", func() {
		When("got empty uid or password", func() {
			It("should return UnauthorizedException", func() {
				c1, ex1 := handler.Login("  ", "password")

				Expect(c1).Should(BeNil())
				Expect(ex1).ShouldNot(BeNil())
				Expect(ex1.IsUnauthorizedException()).Should(BeTrue())

				c2, ex2 := handler.Login("uid", "  ")

				Expect(c2).Should(BeNil())
				Expect(ex2).ShouldNot(BeNil())
				Expect(ex2.IsUnauthorizedException()).Should(BeTrue())

			})
		})

		When("AuthIdentity does not exist", func() {
			It("should return BadRequestException", func() {
				claim, ex := handler.Login("uid", "password")

				Expect(claim).Should(BeNil())
				Expect(ex).ShouldNot(BeNil())
				Expect(ex.IsBadRequestException()).Should(BeTrue())
			})
		})

		When("encrypted password is different", func() {
			It("should return BadRequestException", func() {
				// given
				uid := "user"
				email := "me@email.com"
				password := "secret"
				aid, ex1 := handler.Register(uid, email, password)

				Expect(aid).ShouldNot(BeNil())
				Expect(ex1).Should(BeNil())

				// when
				claim, ex2 := handler.Login(uid, password+"1")

				// then
				Expect(claim).Should(BeNil())
				Expect(ex2).ShouldNot(BeNil())
				Expect(ex2.IsBadRequestException()).Should(BeTrue())
			})
		})

		When("got valid uid and password", func() {
			It("should return claim", func() {
				// given
				uid := "user"
				email := "me@email.com"
				password := "secret"
				aid, ex1 := handler.Register(uid, email, password)

				Expect(aid).ShouldNot(BeNil())
				Expect(ex1).Should(BeNil())

				// when
				claim, ex2 := handler.Login(uid, password)

				// then
				Expect(claim).ShouldNot(BeNil())
				Expect(ex2).Should(BeNil())

				Expect(claim.Provider).Should(Equal(user.ProviderPassword))
				Expect(claim.UID).Should(Equal(uid))
				Expect(claim.UserID).Should(BeNumerically(">", uint(0)))
			})
		})
	})
})
