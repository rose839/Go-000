package main

func main() {
	s := service.service{}
	err := s.TagUser("chen", 1)
	if err != nil {
		log.Errorf("tag user failed: %+v.", err)
	}
}