package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alramdein/template-go/model"
	"github.com/alramdein/template-go/usecase"
)

type productHandler struct {
	ProductUsecase usecase.IProductUsecase
}

func NewProductHandler(productUsecase usecase.IProductUsecase) productHandler {
	return productHandler{
		ProductUsecase: productUsecase,
	}
}

func (p productHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload model.Product
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "errornya pake json yang rapih", http.StatusBadRequest)
		return
	}

	err = p.ProductUsecase.CreateUser(r.Context(), payload)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	w.Write([]byte("Success")) // pastikan reponsenya JSON
}
