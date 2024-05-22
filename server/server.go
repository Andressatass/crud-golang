package server

import "net/http"

func StartServer(port string) error {
	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}

	return nil
}
