package main

import (
	"fmt"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/*
Author: Alistair Chambers
Game: Flappy bird clone

i already recreated this game in the past but i did it in android studio. I could add the gravity but that isnt in the specifcations
*/

func main() {
	rl.InitWindow(800, 450, "Flappy bird clone")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// Load textures
	birdSprite := rl.LoadTexture("textures/flick.png")
	//use a rec boxes for the pipes

	// Set up bird
	birdPos := rl.NewVector2(100, 200)
	birdSize := rl.NewVector2(50, 50)

	//added gradual gravity where the bird will fall faster the longer it falls
	birdVelocity := 0.0
	birdGravity := 0.1
	birdJump := -2.0
	maxFallSpeed := 2.0 // Limit max falling speed

	// Set up pipes
	pipeWidth := 50 // width is stable for all pipes
	pipeGap := 100
	pipeSpeed := 4
	pipeColor := rl.NewColor(0, 255, 0, 255)
	pipe1 := rl.NewRectangle(800, 0, float32(pipeWidth), float32(rand.IntN(200)+100))
	pipe2 := rl.NewRectangle(800, pipe1.Height+float32(pipeGap), float32(pipeWidth), 450-pipe1.Height-float32(pipeGap))

	//score variable
	score := 0
	gameOver := false

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Blue)
		//game loop
		if !gameOver {
			/*
				// Player movement (up/down)
				if rl.IsKeyDown(rl.KeyW) && birdPos.Y > 0 {
					birdPos.Y -= 3
				}
				if rl.IsKeyDown(rl.KeyS) && birdPos.Y < 450-birdSize.Y {
					birdPos.Y += 3
				}*/

			// Player fly
			if rl.IsKeyPressed(rl.KeySpace) {
				birdVelocity = birdJump // Reset velocity when jumping
			}

			//outlines of the bird, the "real" size of the bird caused some issues with the pipes
			rl.DrawRectangle(int32(birdPos.X), int32(birdPos.Y), int32(birdSize.X), int32(birdSize.Y), rl.Black)

			// Move pipes
			pipe1.X -= float32(pipeSpeed)
			pipe2.X -= float32(pipeSpeed)

			// Reset pipes if off-screen
			if pipe1.X < -float32(pipeWidth) {
				pipe1.X = 800
				pipe1.Height = float32(rand.IntN(200) + 100)
				pipe2.X = 800
				pipe2.Y = pipe1.Height + float32(pipeGap)
				pipe2.Height = 450 - pipe2.Y

				// Increment score when pipes reset
				score++
			}

			// Check for collision
			birdHitbox := rl.NewRectangle(birdPos.X, birdPos.Y, birdSize.X, birdSize.Y)
			pipe1Hitbox := rl.NewRectangle(pipe1.X, pipe1.Y, pipe1.Width, pipe1.Height)
			pipe2Hitbox := rl.NewRectangle(pipe2.X, pipe2.Y, pipe2.Width, pipe2.Height)

			if rl.CheckCollisionRecs(birdHitbox, pipe1Hitbox) || rl.CheckCollisionRecs(birdHitbox, pipe2Hitbox) {
				gameOver = true
			}
		}
		// Apply gravity
		birdVelocity += birdGravity
		if birdVelocity > maxFallSpeed {
			birdVelocity = maxFallSpeed // Limit max falling speed
		}
		birdPos.Y += float32(birdVelocity)

		//game over if the bird hits the ground
		if birdPos.Y > 450 {
			gameOver = true
		}

		// Draw bird
		rl.DrawTextureEx(birdSprite, birdPos, 0, birdSize.X/float32(birdSprite.Width), rl.White)

		// Draw pipes
		rl.DrawRectangleRec(pipe1, pipeColor)
		rl.DrawRectangleRec(pipe2, pipeColor)

		// Draw score
		rl.DrawText(fmt.Sprintf("Score: %d", score), 700, 10, 20, rl.White)

		// Handle game over state
		if gameOver {
			rl.DrawText("Game Over", 300, 200, 50, rl.White)
			rl.DrawText("Press R to Restart", 290, 270, 25, rl.White)

			// Restart on key press
			if rl.IsKeyPressed(rl.KeyR) {
				birdPos = rl.NewVector2(100, 200)
				pipe1 = rl.NewRectangle(800, 0, float32(pipeWidth), float32(rand.IntN(200)+100))
				pipe2 = rl.NewRectangle(800, pipe1.Height+float32(pipeGap), float32(pipeWidth), 450-pipe1.Height-float32(pipeGap))
				score = 0
				//switches the game back to running
				gameOver = false
			}
		}

		rl.EndDrawing()
	}
}
