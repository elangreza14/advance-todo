package api

func (s *server) newRouter() {
	// init middleware
	middleware := newMiddleware(s.conf, s, s.core.AuthService)

	s.fbr.Use(middleware.ValidateSecret)

	// auth handlers
	handlerAuth := newAuthHandler(s.conf, s, s.core.AuthService)
	routerAuth := s.fbr.Group("/auth")
	routerAuth.Post("/register", handlerAuth.HandleRegister)
	routerAuth.Post("/login", handlerAuth.HandleLogin)
	routerAuth.Get("/profile", middleware.ValidateAccessToken, handlerAuth.HandleGetProfile)
}
