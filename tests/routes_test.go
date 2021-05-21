package tests

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"payments/app"
	"payments/logs"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		logs.Fatal("Error loading '.env' file")
	}
}

// cmd: go test -run TestGetPagos -v
func TestGetPagos(t *testing.T) {
	appTest := fiber.New()
	app.Routes(appTest)
	resp, err := appTest.Test(httptest.NewRequest(http.MethodGet, "/api/pagos", nil))

	utils.AssertEqual(t, 200, resp.StatusCode, "OK")
	t.Logf(body(resp.Body))
	if err != nil {
		utils.AssertEqual(t, 400, resp.StatusCode, "Bad request")
	}
}

// cmd: go test -run TestPostPagos -v
func TestPostPagos(t *testing.T) {
	appTest := fiber.New()
	app.Routes(appTest)

	tests := []struct {
		name    string
		request string
		want    int
		code    string
	}{
		{
			name:    "Outdated",
			request: `{"documentoIdentificacionArrendatario": "1036946622", "codigoInmueble": "8870", "valorPagado": "1000000", "fechaPago": "25/09/2020"}`,
			want:    400,
			code:    "Bad request",
		},
		{
			name:    "Invalid date format",
			request: `{"documentoIdentificacionArrendatario": "1036946622", "codigoInmueble": "8870", "valorPagado": "1000000", "fechaPago": "25-09-2020"}`,
			want:    400,
			code:    "Bad request",
		},
		{
			name:    "Pay complete",
			request: `{"documentoIdentificacionArrendatario": "1036946622", "codigoInmueble": "8870", "valorPagado": "1000000", "fechaPago": "25/06/2021"}`,
			want:    200,
			code:    "OK",
		},
		{
			name:    "Pay parcial",
			request: `{"documentoIdentificacionArrendatario": "1036946622", "codigoInmueble": "8870", "valorPagado": "400000", "fechaPago": "25/07/2021"}`,
			want:    200,
			code:    "OK",
		},
	}

	for _, tt := range tests {
		// run sub-test
		t.Run(tt.name, func(t *testing.T) {
			var json = []byte(tt.request)
			req := httptest.NewRequest(http.MethodPost, "/api/pagos", bytes.NewBuffer(json))
			req.Header.Set("Content-Type", "application/json")
			resp, err := appTest.Test(req)

			utils.AssertEqual(t, tt.want, resp.StatusCode, tt.code)
			t.Logf(body(resp.Body))
			if err != nil {
				utils.AssertEqual(t, tt.want, resp.StatusCode, tt.code)
			}
		})
	}
}

func body(resp io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp)
	b := buf.String()
	return b
}
