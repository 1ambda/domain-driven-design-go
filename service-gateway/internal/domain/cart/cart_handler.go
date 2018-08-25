package cart

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/product"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/user"
	e "github.com/1ambda/domain-driven-design-go/service-gateway/internal/exception"
	dto "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagmodel"
	"github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagserver/swagapi"
	cartapi "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagserver/swagapi/cart"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
)

type CartHandler interface {
	Configure(handlerRegistry *swagapi.GatewayAPI)
	GetUserCart(uid string) (*dto.GetCartItemsOKBody, e.Exception)
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
		sessionStore:      sessionStore,
	}
}

func (h *cartHandlerImpl) Configure(registry *swagapi.GatewayAPI) {
	registry.CartGetCartItemsHandler = cartapi.GetCartItemsHandlerFunc(
		func(params cartapi.GetCartItemsParams) middleware.Responder {
			// TODO: refactor: AuthUtil
			uid, ex := user.HasAuthenticatedSession(h.sessionStore, params.HTTPRequest)
			if ex != nil {
				return cartapi.NewGetCartItemsDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			response, ex := h.GetUserCart(uid)
			if ex != nil {
				return cartapi.NewGetCartItemsDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			return cartapi.NewGetCartItemsOK().WithPayload(response)
		})

	registry.CartAddCartItemHandler = cartapi.AddCartItemHandlerFunc(
		func(params cartapi.AddCartItemParams) middleware.Responder {
			uid, ex := user.HasAuthenticatedSession(h.sessionStore, params.HTTPRequest)
			if ex != nil {
				return cartapi.NewGetCartItemsDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			response, ex := h.AddCartItem()
			return cartapi.NewAddCartItemOK().WithPayload()

		})
}

func (h *cartHandlerImpl) AddCartItem(uid string) (*dto.AddCartItemOKBodyItems, e.Exception) {

	var response *dto.AddCartItemOKBodyItems

	domain.Transact(h.db, func(tx *gorm.DB) e.Exception {
		aid, ex := h.userRepository.FindAuthIdentityByUID(uid)
		if ex != nil {
			ex.Wrap("User does not exist")
			return ex
		}

		u := aid.User
		modelCart, ex := h.cartRepository.CreateCartIfNotExist(tx, u)
		if ex != nil {
			ex.Wrap("Failed to get Cart")
			return ex
		}

		return nil
	})

	return response, nil
}

func (h *cartHandlerImpl) GetUserCart(uid string) (*dto.GetCartItemsOKBody, e.Exception) {

	var response *dto.GetCartItemsOKBody

	domain.Transact(h.db, func(tx *gorm.DB) e.Exception {
		aid, ex := h.userRepository.FindAuthIdentityByUID(uid)
		if ex != nil {
			ex.Wrap("User does not exist")
			return ex
		}

		u := aid.User
		modelCart, ex := h.cartRepository.CreateCartIfNotExist(tx, u)
		if ex != nil {
			ex.Wrap("Failed to get Cart")
			return ex
		}

		modelCartItems, ex := h.cartRepository.FindAllCartItems(tx, modelCart)
		if ex != nil {
			ex.Wrap("Failed to get CartItem list")
			return ex
		}

		if tx.Error != nil {
			ex := e.NewInternalServerException(tx.Error, "Unknown error occurred")
			return ex
		}

		dtoCart := modelCart.convertToDTO(len(modelCartItems))
		dtoCartItems := make([]*dto.CartItem, 0)

		for i := range modelCartItems {
			modelCartItem := modelCartItems[i]
			// TODO: price
			dtoCartItems = append(dtoCartItems, modelCartItem.convertToDTO())
		}

		response = &dto.GetCartItemsOKBody{
			Cart:         dtoCart,
			CartItemList: dtoCartItems,
		}

		return nil
	})

	return response, nil
}
