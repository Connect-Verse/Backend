package rtc
import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)


func HttpSDPServer(port int) chan string {
	sdpChan := make(chan string)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		fmt.Fprintf(w, "done") //nolint: errcheck
		sdpChan <- string(body)
	})

	go func() {
		// nolint: gosec
		panic(http.ListenAndServe(":"+strconv.Itoa(port), nil))
	}()

	return sdpChan
}