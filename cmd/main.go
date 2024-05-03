// package main

// import (
// 	"metaballs/internal"

// 	"github.com/pkg/profile"
// )

// // import "github.com/pkg/profile"

// func main() {
// 	defer profile.Start(profile.CPUProfile,profile.ProfilePath("./profile")).Stop()
// 	// defer profile.Start(profile.TraceProfile, profile.ProfilePath("./profile"), profile.NoShutdownHook,).Stop()
// 	internal.NewMetaBallApp(
// 		6,
// 		internal.SlowSpeedBall,
// 		internal.LargeSizeBall,
//         internal.Pink,
// 		internal.FPS(45),
// 		internal.ScreenMedium,
// 		internal.Resolution1024).
// 		Run()

// }

package main

import (
	"metaballs/app"

	// "github.com/pkg/profile"
)

func main(){
	// defer profile.Start(profile.CPUProfile,profile.ProfilePath("./profile")).Stop()
    app.Run()


}