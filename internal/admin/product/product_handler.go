package product

type ProductHandler struct {
	Service ProductService
}

func NewProductHandler(service ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}
