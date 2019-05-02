package swagger

import (
	"github.com/rs/xid"
	"net/http"
	"strconv"
	"time"
)

func AddSecret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(formatResponse("Invalid input"))
		return
	}
	formData := r.Form
	expireAfterViews, err := strconv.ParseInt(formData.Get("expireAfterViews"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(formatResponse("Invalid input"))
		return
	}
	expireAfter, err := strconv.ParseInt(formData.Get("expireAfter"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(formatResponse("Invalid input"))
		return
	}
	params := SecretParams{
		formData.Get("secret"),
		expireAfterViews,
		expireAfter,
	}
	createAt := time.Now()
	secret := Secret{
		xid.New().String(),
		params.Secret,
		createAt,
		createAt.Add(time.Duration(params.ExpireAfter) * time.Minute),
		params.ExpireAfterViews,
	}
	secret.save()
	w.WriteHeader(http.StatusOK)
	w.Write(formatResponse(secret.Hash))
}

func GetSecret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	hash, err := getSecretHash(r, "/secret/")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(formatResponse("Secret not found"))
		return
	}
	secret := getSecret(hash)
	if secret == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(formatResponse("Secret not found"))
	} else {
		updateRemainingViews(secret)
		if secret.expired() {
			secret.delete()
			w.WriteHeader(http.StatusNotFound)
			w.Write(formatResponse("Secret is expired!!"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(formatResponse(secret))
	}
}
