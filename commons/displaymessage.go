package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/estebanMe/edcomments/models"
)

//DisplayMessage Return a message
func DisplayMessage (w http.ResponseWriter, m models.Message ) {
   j, err := json.Marshal(m)

   if err != nil {
	   log.Fatal("Error al convertir el mensaje: &s", err)
   }

   w.WriteHeader(m.Code)
   w.Write(j)
}