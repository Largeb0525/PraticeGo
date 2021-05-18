package server

func Init() {
	mongointi()
	go gininit()
	lineinit()
}
