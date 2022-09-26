package main

import (
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type Product struct{
	ID int `json:"id"`
	Name string `json:"name"`
}

var(
	product = map[int]*Product{}
	nomor=1
)

func main()  {	
	e:= echo.New()
	
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.GET("/get", GetProd)
	e.POST("/post", CreateProd)
	e.PUT("/Product/:id", updateProd)
	e.DELETE("/Product/:id", deleteProd)

	e.Logger.Fatal(e.Start(":1323"))
}

func GetProd(c echo.Context) error {
	return c.JSON(http.StatusOK, product)
}

func CreateProd(c echo.Context) error {
 p := &Product{
	ID: nomor,
 }
 if err := c.Bind(p); err != nil{
	return err
 }
 product[p.ID]=p
 nomor++
 return c.JSON(http.StatusCreated, p)
}

func updateProd(c echo.Context) error {
	u := new(Product)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	product[id].Name = u.Name
	return c.JSON(http.StatusOK, product[id])
}

func deleteProd(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(product, id)
	return c.NoContent(http.StatusNoContent)
}