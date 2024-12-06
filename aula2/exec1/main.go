package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Produto representa a estrutura de um produto.
type Produto struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// ProdutoStore armazena os produtos e fornece controle de acesso.
type ProdutoStore struct {
	sync.RWMutex
	Produtos []Produto
}

// CarregarProdutos carrega os produtos de um arquivo JSON.
func (ps *ProdutoStore) CarregarProdutos(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("não foi possível abrir o arquivo: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&ps.Produtos); err != nil {
		return fmt.Errorf("erro ao decodificar o JSON: %v", err)
	}

	return nil
}

// ListarProdutos retorna todos os produtos.
func (ps *ProdutoStore) ListarProdutos(w http.ResponseWriter, r *http.Request) {
	ps.RLock()
	defer ps.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ps.Produtos); err != nil {
		http.Error(w, "Erro ao listar produtos", http.StatusInternalServerError)
	}
}

// BuscarProduto retorna um produto pelo ID.
func (ps *ProdutoStore) BuscarProduto(w http.ResponseWriter, r *http.Request) {
	ps.RLock()
	defer ps.RUnlock()

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}

	var produtoEncontrado *Produto
	for _, p := range ps.Produtos {
		if fmt.Sprintf("%d", p.ID) == id {
			produtoEncontrado = &p
			break
		}
	}

	if produtoEncontrado == nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(produtoEncontrado); err != nil {
		http.Error(w, "Erro ao retornar produto", http.StatusInternalServerError)
	}
}

// BuscarProdutosPorPreco realiza a busca de produtos com preço maior que o valor priceGt.
func (ps *ProdutoStore) BuscarProdutosPorPreco(w http.ResponseWriter, r *http.Request) {
	priceGtStr := r.URL.Query().Get("priceGt") // Recupera o parâmetro priceGt
	if priceGtStr == "" {
		http.Error(w, "Parâmetro priceGt não fornecido", http.StatusBadRequest)
		return
	}

	// Converte priceGt para float64
	priceGt, err := strconv.ParseFloat(priceGtStr, 64)
	if err != nil {
		http.Error(w, "Valor de priceGt inválido", http.StatusBadRequest)
		return
	}

	// Filtra os produtos cujo preço é maior que priceGt
	ps.RLock()
	defer ps.RUnlock()

	var produtosFiltrados []Produto
	for _, p := range ps.Produtos {
		if p.Price > priceGt {
			produtosFiltrados = append(produtosFiltrados, p)
		}
	}

	// Se nenhum produto for encontrado, retorna um erro 404
	if len(produtosFiltrados) == 0 {
		http.Error(w, "Nenhum produto encontrado com preço superior a "+fmt.Sprintf("%.2f", priceGt), http.StatusNotFound)
		return
	}

	// Retorna a lista de produtos filtrados
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(produtosFiltrados); err != nil {
		http.Error(w, "Erro ao retornar produtos filtrados", http.StatusInternalServerError)
	}
}

func main() {
	// Inicializando o armazenamento de produtos
	produtoStore := &ProdutoStore{}

	// Carregar produtos a partir do arquivo JSON
	err := produtoStore.CarregarProdutos("products.json")
	if err != nil {
		log.Fatalf("Erro ao carregar produtos: %v", err)
	}

	// Criando o roteador Chi
	r := chi.NewRouter()

	// Middleware de log de requisições
	r.Use(middleware.Logger)

	// Definindo rotas da API
	r.Get("/produtos", produtoStore.ListarProdutos)                // Listar todos os produtos
	r.Get("/produto/{id}", produtoStore.BuscarProduto)             // Buscar produto por ID
	r.Get("/products/search", produtoStore.BuscarProdutosPorPreco) // Buscar produtos com preço maior que priceGt

	// Iniciar o servidor
	log.Println("Iniciando servidor na porta 8080...")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
