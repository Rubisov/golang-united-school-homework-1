package solution

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func helloWorld() string {
	worldMessage := emoji.Sprint("Hello :world_map:!")
	fmt.Println(worldMessage)
	return worldMessage
}
