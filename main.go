package main

//import "github.com/pkg/profile"


func main() {
	// defer profile.Start(profile.CPUProfile,profile.ProfilePath("./profile")).Stop()
	NewMetaBallApp(
		4,
		SlowSpeedBall,
		SmallSizeBall,
		FPS60,
		ScreenLarge,
		Resolution1024).
		Run()

}
