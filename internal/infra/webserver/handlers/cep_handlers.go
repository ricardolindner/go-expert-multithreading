package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetCepHandler(w http.ResponseWriter, r *http.Request, cep string) {
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	resultCh := make(chan string, 2)

	go fetchFromAPI(ctx, "ViaCEP", "https://viacep.com.br/ws/"+cep+"/json/", resultCh)
	go fetchFromAPI(ctx, "BrasilAPI", "https://brasilapi.com.br/api/cep/v1/"+cep, resultCh)

	select {
	case msg := <-resultCh:
		fmt.Fprintln(w, msg)
	case <-ctx.Done():
		http.Error(w, "❌ Timeout fetching CEP information", http.StatusGatewayTimeout)
	}
}

func fetchFromAPI(ctx context.Context, apiName, url string, ch chan<- string) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		ch <- fmt.Sprintf("❌ Error in request %s: %v", apiName, err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// Pode ser erro de timeout, cancelamento, DNS, etc.
		ch <- fmt.Sprintf("❌ Error in request %s: %v", apiName, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("❌ Error fetching data from %s: %v", apiName, err)
		return
	}

	select {
	case ch <- fmt.Sprintf("✅ Winning API: %s\n%s", apiName, string(body)):
	case <-ctx.Done():
	}
}
