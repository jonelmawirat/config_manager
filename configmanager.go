
import (
	"errors"
	"fmt"
    "sync"
)

type Config interface {
    set(string,string)
    get(string)     (string, error)
}

type config struct {
    data        map[string]string
}

var instance *config
var once sync.Once

func NewConfig() Config {
    once.Do(func(){
            if instance == nil {
                instance = &config{
                    data: make(map[string]string),
                }
            }
        })
    return instance
}

func (c *config) String() string {
    var output string
    output += fmt.Sprintf("{")
    for k,v := range c.data {
        output += fmt.Sprintf("%v: %v,",k,v)
    }
    output += fmt.Sprintf("}")

    return output
}


func (c *config) set(k,v string) {
    c.data[k] = v
}

func (c *config) get(k string) (string,error) {
    v,err := c.data[k]
    if err == false {
        return "",errors.New("Invalid Key")
    }
    return v,nil
}


