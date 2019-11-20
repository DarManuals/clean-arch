package main

func main() {
	err := Initialize().Invoke(func(s Server) {
		s.Run()
	})
	if err != nil {
		panic(err)
	}
}
