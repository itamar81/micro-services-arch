
func NewProducts(l *log.Logger) Products{
	return &Products{l}
}

func (p *Products) ServerHTTP(rw http.ResponseWriter , r *http.Request){
	if r.Method == http.MethodGet {package hendlers

import (
	"log"
	"net/http"
	"github.com/itamar81/api-gateway/data"
)


type Products struct{
	l *log.Logger
}
		p.getProducts(wr,r)
		return
	}

}

func (p *Products) getProducts(rw http.ResponseWriter , r *http.Request){
	p.l.Println("Handle get Products")
	listProducts := data.gatewayProducts

}