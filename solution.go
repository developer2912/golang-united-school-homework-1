package solution

import "github.com/kyokomi/emoji"

const (
	message string = "Hello :world_map:!"
)

func GetMessage() string {
	return emoji.Sprint(message)
}
