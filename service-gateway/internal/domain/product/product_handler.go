package product

import (
	e "github.com/1ambda/domain-driven-design-go/service-gateway/internal/exception"
	dto "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagmodel"
	"github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagserver/swagapi"
	productapi "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagserver/swagapi/product"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"strconv"
)

type ProductHandler interface {
	Configure(handlerRegistry *swagapi.GatewayAPI)
	FindAllProducts(itemCountPerPage int, currentPageOffset int) (int, []*Product, e.Exception)
}

type productHandlerImpl struct {
	productRepository Repository
}

func NewProductHandler(repo Repository) ProductHandler {
	return &productHandlerImpl{
		productRepository: repo,
	}
}

func (h *productHandlerImpl) Configure(registry *swagapi.GatewayAPI) {
	registry.ProductFindAllHandler = productapi.FindAllHandlerFunc(
		func(params productapi.FindAllParams) middleware.Responder {
			currentPageOffset := int(*params.CurrentPageOffset)
			itemCountPerPage := int(*params.ItemCountPerPage)
			totalCount, productList, ex := h.FindAllProducts(currentPageOffset, itemCountPerPage)
			totalItemCount := int64(totalCount)

			if ex != nil {
				return productapi.NewFindAllDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			rows := make([]*dto.Product, 0)
			for i := range productList {
				model := productList[i]
				converted := model.convertToDTO()
				rows = append(rows, converted)
			}

			response := dto.FindAllOKBody{
				Pagination: &dto.Pagination{
					CurrentPageOffset: params.CurrentPageOffset,
					ItemCountPerPage:  params.ItemCountPerPage,
					TotalItemCount:    &totalItemCount,
				},

				Rows: rows,
			}

			return productapi.NewFindAllOK().WithPayload(&response)
		})

	registry.ProductFindOneWithOptionsHandler = productapi.FindOneWithOptionsHandlerFunc(
		func(params productapi.FindOneWithOptionsParams) middleware.Responder {
			if params.ProductID == nil {
				err := errors.New("Got invalid Product ID")
				ex := e.NewBadRequestException(err, "Can't find empty Product")
				return productapi.NewFindOneWithOptionsDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			var productID, err = strconv.ParseUint(*params.ProductID, 10, 64)
			if err != nil {
				ex := e.NewInternalServerException(err, "Failed to find Product")
				return productapi.NewFindOneWithOptionsDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			product, productOptions, ex := h.FindProductWithOptions(uint(productID))
			if ex != nil {
				return productapi.NewFindOneWithOptionsDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			options := make([]*dto.ProductOption, 0)
			for i := range productOptions {
				model := productOptions[i]
				converted := model.convertToDTO()
				options = append(options, converted)
			}

			response := dto.FindOneWithOptionsOKBody{
				Product: product.convertToDTO(),
				Options: options,
			}

			return productapi.NewFindOneWithOptionsOK().WithPayload(&response)
		})
}

func (h *productHandlerImpl) FindAllProducts(currentPageOffset int, itemCountPerPage int) (int, []*Product, e.Exception) {
	if itemCountPerPage <= 0 || itemCountPerPage > 20 {
		itemCountPerPage = 20
	}
	if currentPageOffset < 0 {
		currentPageOffset = 0
	}

	totalCount, productList, ex := h.productRepository.FindAllProducts(itemCountPerPage, currentPageOffset)
	if ex != nil {
		return 0, nil, ex
	}

	return totalCount, productList, nil
}

func (h *productHandlerImpl) FindProductWithOptions(productID uint) (*Product, []*ProductOption, e.Exception) {
	return h.productRepository.FindProductWithOptions(productID)
}
