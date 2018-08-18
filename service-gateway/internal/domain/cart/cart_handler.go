package cart

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagserver/swagapi"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/user"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/product"
	cartapi "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagserver/swagapi/cart"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	e "github.com/1ambda/domain-driven-design-go/service-gateway/internal/exception"
	dto "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagmodel"
)

type CartHandler interface {
	Configure(handlerRegistry *swagapi.GatewayAPI)
	GetUserCart(uid string) ()
}

type cartHandlerImpl struct {
	cartRepository    Repository
	userRepository    user.Repository
	productRepository product.Repository
	sessionStore      sessions.Store
	db                *gorm.DB
}

func NewCartHandler(sessionStore sessions.Store, db *gorm.DB, cartRepo Repository, userRepo user.Repository, productRepo product.Repository) CartHandler {
	return &cartHandlerImpl{
		db:                db,
		cartRepository:    cartRepo,
		userRepository:    userRepo,
		productRepository: productRepo,
		sessionStore: sessionStore,
	}
}

func (h *cartHandlerImpl) Configure(registry *swagapi.GatewayAPI) {
	registry.CartGetUserCartHandler = cartapi.GetUserCartHandlerFunc(
		func(params cartapi.GetUserCartParams) middleware.Responder {
			// TODO: refactor: AuthUtil
			uid, ex := user.HasAuthenticatedSession(h.sessionStore, params.HTTPRequest)
			if ex != nil {
				return cartapi.NewGetUserCartDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			tx := h.db.Begin()
			if tx.Error != nil {
				ex := e.NewInternalServerException(tx.Error, "Unknown error occurred")
				return cartapi.NewGetUserCartDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			aid, ex := h.userRepository.FindAuthIdentityByUID(uid)
			if ex != nil {
				ex.Wrap("User does not exist")
				return cartapi.NewGetUserCartDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			u := aid.User
			modelCart, ex := h.cartRepository.CreateCartIfNotExist(tx, u)
			if ex != nil {
				ex.Wrap("Failed to get Cart")
				return cartapi.NewGetUserCartDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			modelCartItems := make([]*CartItem, 0)

			tx.Commit()
			if tx.Error != nil {
				ex := e.NewInternalServerException(tx.Error, "Unknown error occurred")
				return cartapi.NewGetUserCartDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			dtoCart := modelCart.convertToDTO(len(modelCartItems))
			dtoCartItems := make([]*dto.CartItem, 0)

			//for i, _ := range modelCartItems {
			//	//modelCartItem := modelCartItems[i]
			//	// modelCartItem.toSwaggerModel()
			//}

			response := cartapi.GetUserCartOKBody{
				Cart:         dtoCart,
				CartItemList: dtoCartItems,
			}
			return cartapi.NewGetUserCartOK().WithPayload(&response)

		})
}

func (cartHandlerImpl) GetUserCart(uid string) () {
	panic("implement me")
}
