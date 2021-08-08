package wire01

// @Description  go语言学习练习
// @Author playclouds
// @Update    2021/8/8 19:54
//import "github.com/google/wire"

type apple struct {
	name  string
	score int
}

func (a *apple) say() string {
	return a.name
}

type banana struct {
	job   string
	price int
}

func (b *banana) speak() string {
	return b.job
}

func main() {

}
