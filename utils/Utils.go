package utils

type JsonStatusResponse struct {
	Status string `json:"status"`
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
