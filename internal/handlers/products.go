package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/evermos/boilerplate-go/internal/domain/products/models"
	"github.com/evermos/boilerplate-go/internal/domain/products/services"
	"github.com/evermos/boilerplate-go/transport/http/middleware"
	"github.com/google/uuid"

	"github.com/go-chi/chi"
)

type ProductsHandler struct {
	Service        services.Service
	Authentication *middleware.Authentication
}

func ProvideProductsHandler(service services.Service, auth *middleware.Authentication) ProductsHandler {
	return ProductsHandler{
		Service:        service,
		Authentication: auth,
	}
}

// Router sets up the router for this domain.
func (h *ProductsHandler) Router(r chi.Router) {
	r.Route("/", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(generateRequestID)
			// Middleware for authentication using api key
			r.Use(h.Authentication.SetXApiKey)
			r.Use(h.Authentication.XApiAuthentication)
			r.Post("/products", h.CreateProduct)
			r.Get("/products", h.ListProducts)
			r.Put("/variants/{variantID}", h.UpdateVariant)
			r.Delete("/products/{productID}", h.SoftDeleteProduct)
			r.Delete("/products/hard/{productID}", h.HardDeleteProduct)
		})
	})
}

// Middleware to generate and include the requestId in the request context
func generateRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String() + "-" + "mock"
		ctx := context.WithValue(r.Context(), "requestId", requestID)
		r = r.WithContext(ctx)
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r)
	})
}

func (h *ProductsHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Define the required struct for the request body
	var req struct {
		UserID      string `json:"user_id"`
		BrandID     string `json:"brand_id"`
		WarehouseID string `json:"warehouse_id"`
		Name        string `json:"name"`
		CreatedBy   string `json:"created_by"`
		UpdatedBy   string `json:"updated_by"`
	}

	// Decode the request body into the req struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.UserID == "" || req.BrandID == "" || req.WarehouseID == "" || req.Name == "" || req.CreatedBy == "" || req.UpdatedBy == "" {
		http.Error(w, "user_id, brand_id, warehouse_id, name, created_by, and updated_by fields are required", http.StatusBadRequest)
		return
	}

	// Ensure that created_by is equal to updated_by
	if req.CreatedBy != req.UpdatedBy {
		http.Error(w, "created_by must be equal to updated_by", http.StatusBadRequest)
		return
	}

	UserID, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user_id format", http.StatusBadRequest)
		return
	}

	BrandID, err := strconv.ParseInt(req.BrandID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid brand_id format", http.StatusBadRequest)
		return
	}

	WarehouseID, err := strconv.ParseInt(req.WarehouseID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user_id format", http.StatusBadRequest)
		return
	}

	// Validate that the CreatedBy user has user_type "admin"
	users, err := h.Service.GetUserByUsername(req.CreatedBy)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	if users.UserType != "admin" {
		http.Error(w, "Only admin users are allowed to create products", http.StatusForbidden)
		return
	}

	// Create the product model from the request data
	product := &models.Product{
		UserID:      UserID,
		BrandID:     BrandID,
		WarehouseID: WarehouseID,
		Name:        req.Name,
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.UpdatedBy,
	}

	newProduct, err := h.Service.CreateProduct(product)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message": "Product successfully created",
		"product": newProduct,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *ProductsHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	q := r.URL.Query()
	brandName := q.Get("brandName")
	productName := q.Get("productName")
	variantName := q.Get("variantName")
	status := q.Get("status")
	sortBy := q.Get("sortBy")
	page, _ := strconv.Atoi(q.Get("page"))
	size, _ := strconv.Atoi(q.Get("size"))

	// Set default values for page and size if they are not provided or invalid
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}

	// Call the service to fetch the products with the specified filters and sorting
	products, err := h.Service.ListProducts(models.ProductFilter{
		BrandName:   brandName,
		ProductName: productName,
		VariantName: variantName,
		Status:      status,
	},
		sortBy,
		page,
		size)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h *ProductsHandler) UpdateVariant(w http.ResponseWriter, r *http.Request) {
	variantID := chi.URLParam(r, "variantID")
	id, err := strconv.ParseInt(variantID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid variant ID", http.StatusBadRequest)
		return
	}

	// Check if the variant with the given ID exists in the database
	_, err = h.Service.GetVariantByID(id)
	if err != nil {
		http.Error(w, "Variant not found", http.StatusBadRequest)
		return
	}

	// Define the required struct for the request body
	var req struct {
		Name      string `json:"name"`
		Price     string `json:"price"`
		Stock     string `json:"stock"`
		Status    string `json:"status"`
		UpdatedBy string `json:"updated_by"`
	}

	// Decode the request body into the req struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Name == "" || req.Price == "" || req.Stock == "" || req.Status == "" || req.UpdatedBy == "" {
		http.Error(w, "name, price, stock, status, and updated_by fields are required", http.StatusBadRequest)
		return
	}

	// Convert the price string to float64
	price, err := strconv.ParseFloat(req.Price, 64)
	if err != nil {
		http.Error(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	// Convert the stock string to int
	stock, err := strconv.Atoi(req.Stock)
	if err != nil {
		http.Error(w, "Invalid stock format", http.StatusBadRequest)
		return
	}

	// Validate that the UpdatedBy user has user_type "admin"
	users, err := h.Service.GetUserByUsername(req.UpdatedBy)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	if users.UserType != "admin" {
		http.Error(w, "Only admin users are allowed to update variants", http.StatusForbidden)
		return
	}

	// Create the variant model from the request data
	variant := &models.UpdateVariant{
		Name:      req.Name,
		Price:     price,
		Stock:     stock,
		Status:    req.Status,
		UpdatedBy: req.UpdatedBy,
	}

	updatedVariant, err := h.Service.UpdateVariant(id, variant)
	if err != nil {
		http.Error(w, "Failed to update variant", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// response body
	response := map[string]interface{}{
		"message": "Variant successfully updated",
		"variant": updatedVariant,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *ProductsHandler) SoftDeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productID")
	id, err := strconv.ParseInt(productID, 10, 64) //int64
	// id, err := strconv.Atoi(productID) //int
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Check if the product with the given ID exists in the database
	_, err = h.Service.GetProductByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusBadRequest)
		return
	}

	// Define the required struct for the request body
	var req struct {
		DeletedBy string `json:"deleted_by"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.DeletedBy == "" {
		http.Error(w, "deleted_by is required", http.StatusBadRequest)
		return
	}

	// Validate that the DeletedBy user has user_type "admin"
	users, err := h.Service.GetUserByUsername(req.DeletedBy)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	if users.UserType != "admin" {
		http.Error(w, "Only admin users are allowed to soft delete variants", http.StatusForbidden)
		return
	}

	err = h.Service.SoftDeleteProduct(id, req.DeletedBy)
	if err != nil {
		http.Error(w, "Failed to soft delete product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product soft deleted successfully"})
}

func (h *ProductsHandler) HardDeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productID")
	id, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Check if the product with the given ID exists in the database
	_, err = h.Service.GetProductByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusBadRequest)
		return
	}

	err = h.Service.HardDeleteProduct(id)
	if err != nil {
		http.Error(w, "Failed to hard delete product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product hard deleted successfully"})
}
