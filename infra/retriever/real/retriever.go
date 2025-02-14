package real

import "net/http"

type Retriever struct {
	UserAgent string
	TimeOut   int
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp.Status
}
