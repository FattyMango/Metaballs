package main

//import "github.com/pkg/profile"


func main() {
	// defer profile.Start(profile.CPUProfile,profile.ProfilePath("./profile")).Stop()
	NewMetaBallApp(
		4,
		SlowSpeedBall,
		LargeSizeBall,
		FPS(45),
		ScreenSmall,
		Resolution512).
		Run()

}
