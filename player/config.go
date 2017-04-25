package player

type Config struct {
    CurrentWealth int
}

var PlayerConfig := Config { 0 }

func UpdateWealth(value int) {
    PlayerConfig.CurrentWealth = value
}
