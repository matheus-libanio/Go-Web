package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/matheus-libanio/Go-Web/product_crud/internal/application"
	"github.com/matheus-libanio/Go-Web/product_crud/internal/handler"
	"github.com/matheus-libanio/Go-Web/product_crud/internal/model"
	"github.com/matheus-libanio/Go-Web/product_crud/internal/repository"
	"github.com/matheus-libanio/Go-Web/product_crud/internal/service"
	"github.com/stretchr/testify/require"
)

func TestProductHandler_Create(t *testing.T) {
	sdb := repository.NewProductDB()
	db := map[string]*model.Product{} // Usamos um mapa vazio para facilitar os testes
	sdb.DB = db
	st := service.NewServiceProducts(sdb)
	app := application.NewApplicationProducts(st)
	p := handler.NewProductHandler(app)

	tests := []struct {
		name         string
		requestBody  map[string]interface{}
		expectedCode int
		expectedBody string
	}{
		{
			name: "success to create product",
			requestBody: map[string]interface{}{
				"name":         "Oil - Margarine",
				"quantity":     439,
				"code_value":   "S82254D",
				"is_published": true,
				"expiration":   "15/12/2021",
				"price":        71.42,
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "bad request due to invalid body",
			requestBody: map[string]interface{}{
				// Corpo da requisição malformado
				"quantity": 100,
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"message":"Bad Request : invalid request body","data":null,"error":true}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/products", bytes.NewBuffer(body))
			res := httptest.NewRecorder()

			p.Create(res, req)

			require.Equal(t, tt.expectedCode, res.Code)

			if tt.expectedCode == http.StatusCreated {
				var response handler.ResponseBodyProduct
				err := json.Unmarshal(res.Body.Bytes(), &response)
				require.NoError(t, err)

				// Verifique se o ID retornado é válido
				_, err = uuid.Parse(response.Data.Id)
				require.NoError(t, err) // Deve ser um UUID válido

				// Além disso, verifique outras partes do retorno, se necessário
				require.Equal(t, "Product created", response.Message)
				require.Equal(t, "Oil - Margarine", response.Data.Name)
				require.Equal(t, 439, response.Data.Quantity)
				require.Equal(t, "S82254D", response.Data.Code_value)
				require.Equal(t, true, response.Data.Is_published)
				require.Equal(t, "15/12/2021", response.Data.Expiration)
				require.Equal(t, 71.42, response.Data.Price)
			} else {
				require.JSONEq(t, tt.expectedBody, res.Body.String())
			}
		})
	}
}
func TestProductHandler_GetAll(t *testing.T) {
	t.Run("success to get products", func(t *testing.T) {
		// arrange

		sdb := repository.NewProductDB()
		db := map[string]*model.Product{
			"1": {Name: "Oil - Margarine", Quantity: 439, Code_value: "S82254D", Is_published: true, Expiration: "15/12/2021", Price: 71.42},
			"2": {Name: "Pineapple - Canned, Rings", Quantity: 345, Code_value: "M4637", Is_published: true, Expiration: "09/08/2021", Price: 352.79},
		}
		sdb.DB = db
		st := service.NewServiceProducts(sdb)
		app := application.NewApplicationProducts(st)
		hd := handler.NewProductHandler(app)
		// act

		req := httptest.NewRequest("GET", "/products", nil)
		res := httptest.NewRecorder()

		hd.GetAll(res, req)
		//assert
		expectedCode := http.StatusOK
		expectedBody := `{
			"message": "success to get products",
			"data": [
				{"id": "", "name": "Oil - Margarine", "quantity": 439, "code_value": "S82254D", "is_published": true, "expiration": "15/12/2021", "price": 71.42},
				{"id": "", "name": "Pineapple - Canned, Rings", "quantity": 345, "code_value": "M4637", "is_published": true, "expiration": "09/08/2021", "price": 352.79}
			],
			"error": false
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}
