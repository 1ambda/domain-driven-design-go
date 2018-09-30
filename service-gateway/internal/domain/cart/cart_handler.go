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
	"github.com/pkg/errors"
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

			response, ex := h.AddCartItem(uid, params.Body)
			if ex != nil {
				return cartapi.NewGetCartItemsDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			return cartapi.NewAddCartItemOK().WithPayload(response)
		})
}

func (h *cartHandlerImpl) AddCartItem(uid string, request *dto.AddCartItemDTO) (*dto.AddCartItemDTO, e.Exception) {

	TODO: create convert.go that ibcludes converting functions for swagger types.
	TODO: service
	TODO: root aggregator (root model) and rootRepository
	reqProductID := uint(*request.ProductID)
	reqOptionIDList := make([]uint, 0)
	reqQuantity := uint(*request.Quantity)
	for i := range request.ProductOptionIDList {
		reqOptionID := uint(request.ProductOptionIDList[i])
		reqOptionIDList = append(reqOptionIDList, reqOptionID)
	}

	// TODO: handle policy validation in service layer
	if reqQuantity > 100 {
		err := errors.Errorf("requested ProductID %d should be greater than 0", reqProductID)
		ex := e.NewBadRequestException(err, "Some of requested options are not available")
		return nil, ex
	}

	var response *dto.AddCartItemDTO

	ex := domain.Transact(h.db, func(tx *gorm.DB) e.Exception {
		aid, ex := h.userRepository.FindAuthIdentityByUID(uid)
		if ex != nil {
			ex.Wrap("User does not exist")
			return ex
		}

		u := aid.User
		modelCart, ex := h.cartRepository.FindOrCreateCart(tx, u)
		if ex != nil {
			ex.Wrap("Failed to find Cart")
			return ex
		}

		modelCartItems, ex := h.cartRepository.FindAllCartItems(tx, modelCart)
		if ex != nil {
			ex.Wrap("Failed to find CartItem list")
			return ex
		}

		modelProduct, modelProductOptions, ex := h.productRepository.FindProductWithOptions(reqProductID)
		if ex != nil {
			ex.Wrap("Failed to find Product")
			return ex
		}

		totalPrice := modelProduct.Price

		// check the requested options are all available
		for _, reqOptionID := range reqOptionIDList {

			valid := false

			for _, availableOption := range modelProductOptions {
				if reqOptionID == availableOption.ProductID {
					valid = true
					totalPrice += availableOption.Price
					break
				}
			}

			if !valid {
				err := errors.Errorf("Requested option %d doesn't match with available product options", reqOptionID)
				ex := e.NewBadRequestException(err, err.Error())
				return ex
			}
		}

		_, ex = h.cartRepository.AddCartItem(
			tx, modelCart, len(modelCartItems), reqQuantity, reqProductID, reqOptionIDList,
		)

		dtoProductID := int64(reqProductID)
		dtoQuantity := int64(reqQuantity)
		dtoProductIDList := make([]int64, 0)
		for i := range reqOptionIDList {
			dtoProductIDList = append(dtoProductIDList, int64(reqOptionIDList[i]))
		}

		response = &dto.AddCartItemDTO{
			ProductID:           &dtoProductID,
			ProductOptionIDList: dtoProductIDList,
			Quantity:            &dtoQuantity,
		}

		return nil
	})

	return response, ex
}

func (h *cartHandlerImpl) GetUserCart(uid string) (*dto.GetCartItemsOKBody, e.Exception) {

	var response *dto.GetCartItemsOKBody

	ex := domain.Transact(h.db, func(tx *gorm.DB) e.Exception {
		aid, ex := h.userRepository.FindAuthIdentityByUID(uid)
		if ex != nil {
			ex.Wrap("User does not exist")
			return ex
		}

		u := aid.User
		modelCart, ex := h.cartRepository.FindOrCreateCart(tx, u)
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
		dtoCart.TotalPrice =
		dtoCartItemList := make([]*dto.GetCartItemsOKBodyCartItemListItems, 0)

		for i := range modelCartItems {
			modelCartItem := modelCartItems[i]
			//modelCartItemOptionList := make([]*CartItemOption, 0)

			dtoCartItem := &dto.GetCartItemsOKBodyCartItemListItems{
				CartItem:           modelCartItem.convertToDTO(),
				CartItemOptionList: nil, // TODO
			}

			// TODO: price
			dtoCartItemList = append(dtoCartItemList, dtoCartItem)
		}

		response = &dto.GetCartItemsOKBody{
			Cart:         dtoCart,
			CartItemList: dtoCartItemList,
		}

		return nil
	})

	return response, ex
}
