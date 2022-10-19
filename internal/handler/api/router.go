package api

func (s *server) newRouter() {

	// auth handlers
	var handlerAuth iAuthApiHandler = NewAuthHandler(s, s.core.AuthService)
	routerAuth := s.fbr.Group("/auth")
	routerAuth.Post("/register", handlerAuth.HandleRegister)
	routerAuth.Post("/login", handlerAuth.HandleLogin)
}
