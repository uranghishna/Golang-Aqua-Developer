package handler

import (
	"ecommerce/entity"
	"ecommerce/usecase"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductUsecase *usecase.ProductUsecase
}

func NewProductHandler(ProductUsecase *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{ProductUsecase}
}

// Create products godoc
// @Summary Create Product.
// @Description create product.
// @Tags products
// @Param Body body entity.ProductRequest true "silahkan daftarkan data product"
// @Consumes multipart/form-data
// @Param Body formData file true "file anda"
// @Produce json
// @Success 201 {object} entity.ProductResponse
// @Router /products/create [post]
func (handler ProductHandler) CreateProduct(c echo.Context) error {
	name := c.FormValue("name")
	price, _ := strconv.Atoi(c.FormValue("price"))
	category := c.FormValue("category")
	description := c.FormValue("description")
	file, err := c.FormFile("photo")

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// destinattion
	dst, err := os.Create(fmt.Sprintf("%s%s", "storage/", file.Filename))
	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), dst.Name())

	req := entity.ProductRequest{
		Name:        name,
		Photo:       filePath,
		Price:       price,
		Category:    category,
		Description: description,
	}
	Product, err := handler.ProductUsecase.CreateProduct(req)

	if err != nil {
		return err
	}

	return c.JSON(201, Product)
}

func (handler ProductHandler) GetAllProduct(c echo.Context) error {

	Product, err := handler.ProductUsecase.GetAllProduct()

	if err != nil {
		return err
	}

	return c.JSON(200, Product)
}

func (handler ProductHandler) GetProductByID(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))

	product, err := handler.ProductUsecase.GetProductById(productId)
	if err != nil {
		return err
	}

	return c.JSON(200, product)
}

func (handler ProductHandler) GetProductByCategory(c echo.Context) error {
	productCategory := c.Param("category")

	product, err := handler.ProductUsecase.GetProductByCategory(productCategory)
	if err != nil {
		return err
	}

	return c.JSON(200, product)
}

func (handler ProductHandler) GetProductByPrice(c echo.Context) error {
	priceMin, _ := strconv.Atoi(c.Param("priceMin"))
	priceMax, _ := strconv.Atoi(c.Param("priceMax"))

	product, err := handler.ProductUsecase.GetProductByPrice(priceMin, priceMax)
	if err != nil {
		return err
	}
	return c.JSON(200, product)
}

func (handler ProductHandler) UpdateProduct(c echo.Context) error {
	ProductId, err := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")
	price, _ := strconv.Atoi(c.FormValue("price"))
	category := c.FormValue("category")
	description := c.FormValue("description")
	file, err := c.FormFile("photo")
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// destinattion
	dst, err := os.Create(fmt.Sprintf("%s%s", "storage/", file.Filename))
	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), dst.Name())

	req := entity.UpdateProductRequest{
		Name:        name,
		Photo:       filePath,
		Price:       price,
		Category:    category,
		Description: description,
	}
	ProductRequest := entity.UpdateProductRequest{}
	if err := c.Bind(&ProductRequest); err != nil {
		return err
	}
	Product, err := handler.ProductUsecase.UpdateProduct(req, ProductId)
	if err != nil {
		return err
	}

	return c.JSON(200, Product)
}

func (handler ProductHandler) DeleteProduct(c echo.Context) error {
	ProductId, _ := strconv.Atoi(c.Param("id"))
	err := handler.ProductUsecase.DeleteProduct(ProductId)
	if err != nil {
		return err
	}
	return c.JSON(200, ProductId)
}
