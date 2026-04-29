package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/luizhen1/go-clean-gin-crud/docs"
	"github.com/luizhen1/go-clean-gin-crud/internal/domain"
	"github.com/luizhen1/go-clean-gin-crud/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitProductRoutes(r *gin.Engine, service *services.ProductService) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(1),
	))

	r.POST("/products", createProductHandler(service))
	r.GET("/products", listProductsHandler(service))
	r.GET("/products/:id", getProductByIDHandler(service))
	r.PUT("/products/:id", updateProductHandler(service))
	r.DELETE("/products/:id", deleteProductHandler(service))
}

// @Summary      Criar produto
// @Description  Cria um novo produto no banco de dados
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      domain.Product  true  "Dados do Produto"
// @Success      201      {object}  domain.Product
// @Failure      400      {object}  map[string]string "Erro nos dados enviados"
// @Failure      500      {object}  map[string]string "Erro interno no servidor"
// @Router       /products [post]
func createProductHandler(service *services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product domain.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := service.CreateProduct(&product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, product)
	}
}

// @Summary      Listar produtos
// @Description  Retorna uma lista de todos os produtos cadastrados
// @Tags         products
// @Produce      json
// @Success      200  {array}   domain.Product
// @Failure      500  {object}  map[string]string "Erro ao buscar produtos"
// @Router       /products [get]
func listProductsHandler(service *services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := service.FetchProducts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

// @Summary      Buscar por ID
// @Description  Retorna um único produto pelo ID
// @Tags         products
// @Produce      json
// @Param        id   path      int  true  "ID do Produto"
// @Success      200  {object}  domain.Product
// @Failure      400  {object}  map[string]string "ID inválido"
// @Failure      404  {object}  map[string]string "Produto não encontrado"
// @Router       /products/{id} [get]
func getProductByIDHandler(service *services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}
		product, err := service.GetProductByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if product == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// @Summary      Atualizar produto
// @Description  Atualiza os dados de um produto existente
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      int             true  "ID do Produto"
// @Param        product  body      domain.Product  true  "Novos dados"
// @Success      200      {object}  domain.Product
// @Failure      400      {object}  map[string]string "Dados inválidos"
// @Failure      500      {object}  map[string]string "Erro ao atualizar"
// @Router       /products/{id} [put]
func updateProductHandler(service *services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, _ := strconv.ParseInt(idParam, 10, 64)
		var product domain.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product.ID = id
		err := service.UpdateProduct(&product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// @Summary      Deletar produto
// @Description  Remove um produto do banco de dados
// @Tags         products
// @Param        id   path      int  true  "ID do Produto"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string "ID inválido"
// @Failure      500  {object}  map[string]string "Erro ao deletar"
// @Router       /products/{id} [delete]
func deleteProductHandler(service *services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, _ := strconv.ParseInt(idParam, 10, 64)
		err := service.DeleteProduct(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Produto deletado com sucesso"})
	}
}
