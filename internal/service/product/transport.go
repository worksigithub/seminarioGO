package product

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)
// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method string
	path string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}	
	list = append(list, 
		&endpoint{
			method: "GET",
			path: "/product/:id",
			function: FindProductID(s),
		},
		&endpoint{
			method: "GET",
			path: "/product",
			function: FindProducts(s),
		},
		&endpoint{
			method: "POST",
			path: "/product",
			function: InsertProduct(s),
		},
		&endpoint{
			method: "PUT",
			path: "/product",
			function: UpdateProduct(s),
		},
		&endpoint{
			method: "DELETE",
			path: "/product/:id",
			function: DeleteProduct(s),
		},
	)
	return list
}

func FindProductID(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		result, err := s.FindProductID(ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"product": err,
			})
		}else {
			c.JSON(http.StatusOK, gin.H{
				"product": result,
			})
		}
	}	
}		

func FindProducts(s Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		result, err := s.FindProducts()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"product": err,
			})
		}else {
			c.JSON(http.StatusOK, gin.H{
				"product": result,
			})
		}
	}
}

func InsertProduct(s Service) gin.HandlerFunc {
	var product product
	return func(c *gin.Context) {
		c.BindJSON(&product)
		result, err := s.InsertProduct(product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"product": err,
			})
		}else {
			c.JSON(http.StatusOK, gin.H{
				"product": result,
			})
		}
	}
}

func UpdateProduct(s Service) gin.HandlerFunc {
	var product product
	return func(c *gin.Context) {
		c.BindJSON(&product)
		result, err := s.UpdateProduct(product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"product": err,
			})
		}else {
			c.JSON(http.StatusOK, gin.H{
				"product": result,
			})
		}
	}
}

func DeleteProduct(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		result, err := s.DeleteProduct(ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"product": err,
			})
		}else {
			c.JSON(http.StatusOK, gin.H{
				"product": result,
			})
		}	
	}
}

func responseStatus(result,err error,c *gin.Context){
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"product": err,
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"product": result,
		})
	}
}

func (s httpService) Register(r *gin.Engine){
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}