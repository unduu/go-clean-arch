package handler

import ( 
    "encoding/json" 
    "net/http" 
)

func catch(err error) {
    if err != nil {
        panic(err)
    }
}

// respondwithError return error message
func respondErrorJSON(w http.ResponseWriter, code int, msg string) {
    response, _ := json.Marshal(map[string]string{"message": msg})

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(500)
    w.Write(response)
}

// respondwithJSON write json response format
func respondSuccessJSON(w http.ResponseWriter, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(response)
}