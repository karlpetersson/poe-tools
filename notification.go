// +build windows,!linux,!darwin

package main

import (
    "bitbucket.org/StephenPatrick/go-winaudio/winaudio"
)

var wavfile winaudio.Audio

func playNotification() {
    wavfile.Play()    
}

func initSound() {
    winaudio.InitWinAudio()
    wavfile, _ = winaudio.LoadWav("./audio/Turn.wav")
    wavfile.SetLooping(false)
}