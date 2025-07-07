package startup

func (s server) registerRoutes() {
	v1 := s.router.Group("/v1")
	{
		v1.POST("/account", nil)
	}
}
