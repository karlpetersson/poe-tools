// +build windows

package notifications

import (
    "bitbucket.org/StephenPatrick/go-winaudio/winaudio"
)

var wavfile winaudio.Audio

func Play() {
    wavfile.Play()    
}

func Init() {
    winaudio.InitWinAudio()
    wavfile, _ = winaudio.LoadWav("./audio/Turn.wav")
    wavfile.SetLooping(false)
}