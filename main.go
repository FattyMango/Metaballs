package main

//import "github.com/pkg/profile"


func main() {
	// defer profile.Start(profile.CPUProfile,profile.ProfilePath("./profile")).Stop()
	NewMetaBallApp(
		4,
		NormalSpeedBall,
		SmallSizeBall,
		FPS(45),
		ScreenSmall,
		Resolution1024).
		Run()

}
