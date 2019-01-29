package api

type EventJson struct {
	Host  string
	Tags  []string
	State string
}

type JsonResponse struct {
	Status string
}
