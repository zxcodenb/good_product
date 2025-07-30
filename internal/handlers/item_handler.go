// handlers包处理HTTP请求
package handlers

import (
	"encoding/json"
	"item-value/internal/models"
	"item-value/internal/storage"
	"net/http"
)

// ItemHandler 处理项目的HTTP请求
type ItemHandler struct {
	store storage.Store  // 存储接口实例
}

// NewItemHandler 创建一个新的ItemHandler实例
func NewItemHandler(s storage.Store) *ItemHandler {
	return &ItemHandler{store: s}
}

// RegisterRoutes 注册项目相关的路由
func (h *ItemHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /items", h.CreateItem)      // 创建项目
	mux.HandleFunc("GET /items", h.ListItems)        // 获取项目列表
	mux.HandleFunc("GET /items/{id}", h.GetItem)     // 根据ID获取项目
	mux.HandleFunc("PUT /items/{id}", h.UpdateItem)  // 根据ID更新项目
	mux.HandleFunc("DELETE /items/{id}", h.DeleteItem) // 根据ID删除项目
}

// CreateItem 处理创建新项目的请求
func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	// 解析请求体中的JSON数据
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 调用存储接口创建项目
	createdItem, err := h.store.CreateItem(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回创建成功的项目信息
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdItem)
}

// ListItems 处理获取项目列表的请求
func (h *ItemHandler) ListItems(w http.ResponseWriter, r *http.Request) {
	// 调用存储接口获取所有项目
	items, err := h.store.ListItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回项目列表
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GetItem 处理根据ID获取项目的请求
func (h *ItemHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	// 从路径参数中获取项目ID
	id := r.PathValue("id")
	
	// 调用存储接口获取项目
	item, err := h.store.GetItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// 返回项目信息
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// UpdateItem 处理更新项目的请求
func (h *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	// 从路径参数中获取项目ID
	id := r.PathValue("id")

	var item models.Item
	// 解析请求体中的JSON数据
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 调用存储接口更新项目
	updatedItem, err := h.store.UpdateItem(id, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// 返回更新后的项目信息
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedItem)
}

// DeleteItem 处理删除项目的请求
func (h *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	// 从路径参数中获取项目ID
	id := r.PathValue("id")

	// 调用存储接口删除项目
	if err := h.store.DeleteItem(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// 返回204状态码表示删除成功
	w.WriteHeader(http.StatusNoContent)
}