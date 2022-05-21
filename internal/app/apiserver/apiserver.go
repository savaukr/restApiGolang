package apiserver

//APIServer ...
type APIServer struct{}

func New()*APIServer {
	return &APIServer{}
}

func (s *APIServer) Start() error{
	return nil
}