package megasecond

import (
	"fmt"
	"math"
	"time"
)

// AddMegasecond adds 1 million seconds to time
func AddMegasecond(t time.Time) time.Time {
	seconds := math.Pow(10, 6)
	returnVal := t.Add(time.Second * time.Duration(seconds))
	fmt.Printf("%v", returnVal)
	return returnVal
}
