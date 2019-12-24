package http

import(
	"net/http"
	"encoding/json"
)

// Message ...
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond json
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

// RespondSuccess ...
func RespondSuccess(w http.ResponseWriter, payload interface{}) {
	respond(w, payload, http.StatusOK)
}

// RespondResourceCreated ..
func RespondResourceCreated(w http.ResponseWriter, payload interface{}) {
	respond(w, payload, http.StatusCreated)
}

// RespondBadRequest ..
func RespondBadRequest(w http.ResponseWriter, payload interface{}) {
	respond(w, payload, http.StatusBadRequest)
}

// RespondNotFound ..
func RespondNotFound(w http.ResponseWriter, payload interface{}) {
	respond(w, payload, http.StatusNotFound)
}

// RespondNotModified ..
func RespondNotModified(w http.ResponseWriter, payload interface{}) {
	respond(w, payload, http.StatusNotModified)
}

// RespondUnauthorized ..
func RespondUnauthorized(w http.ResponseWriter, payload interface{}) {
	respond(w, payload, http.StatusUnauthorized)
}

// RespondInternalServer ..
func RespondInternalServer(w http.ResponseWriter, payload interface{}) {
	respond(w, payload, http.StatusInternalServerError)
}

// RespondTooManyRequests ..
func RespondTooManyRequests(w http.ResponseWriter, payload interface{}) {
	respond(w, payload, http.StatusTooManyRequests)
}

func respond(w http.ResponseWriter, payload interface{}, statusCode int) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(response))
}
