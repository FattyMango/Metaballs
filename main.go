package main

// import "github.com/pkg/profile"


func main() {
	// defer profile.Start(profile.CPUProfile,profile.ProfilePath("./profile")).Stop()
	// defer profile.Start(profile.TraceProfile, profile.ProfilePath("./profile"), profile.NoShutdownHook,).Stop()
	NewMetaBallApp(
		4,
		SlowSpeedBall,
		MediumSizeBall,
		FPS(45),
		ScreenMedium,
		Resolution1024).
		Run()

}
