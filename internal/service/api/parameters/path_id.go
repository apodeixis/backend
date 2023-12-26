package parameters

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func FetchPathID(r *http.Request) (int64, error) {
	const pathIDKey = "id"
	pathID := chi.URLParam(r, pathIDKey)
	postID, err := strconv.Atoi(pathID)
	return int64(postID), err
}
