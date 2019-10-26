package ulule

import (
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"log"
	"net/http"
	"time"
)

func exampleHttp() {
	rate := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  1,
	}
	store := memory.NewStore()
	instance := limiter.New(store, rate)
	middleware := stdlib.NewMiddleware(instance)
	http.Handle("/", middleware.Handler(http.HandlerFunc(index)))
	log.Fatal(http.ListenAndServe(":7777", nil))
}
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(`{"message": "ok"}`))
}
func example() {
	rate := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  1,
	}
	store := memory.NewStore()
	instance := limiter.New(store, rate)
}
