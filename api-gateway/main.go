package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func apiGatewayHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[1:]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Belirli URL'lere göre hangi mikroservise yönlendirileceğini kontrol et
	switch url {
	case "user":
		// Burada 'user-service' ile iletişim kurma işlemini gerçekleştirebilirsiniz.
		// Örnek bir çıktı:
		fmt.Fprintf(w, "API Gateway: User Service - %s", string(body))
	case "purchase":
		// Burada 'purchase-service' ile iletişim kurma işlemini gerçekleştirebilirsiniz.
		// Örnek bir çıktı:
		fmt.Fprintf(w, "API Gateway: Purchase Service - %s", string(body))
	default:
		http.Error(w, "Unknown endpoint", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/", apiGatewayHandler)

	// Burada kullanılacak portu ve diğer konfigürasyonları belirleyebilirsiniz.
	port := 8082
	fmt.Printf("API Gateway is running on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
