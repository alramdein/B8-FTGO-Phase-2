package main

// Interface Segregation
// -- Bad
type Eater interface {
	Eat(string)
}

type Worker interface {
	Work()
}

type Sleeper interface {
	Sleep()
}

type Human struct {
}

func (h *Human) Eat(food string) {
	println("Human eat", food)
}

func (h *Human) Sleep() {
	println("Human sleep")
}

func (h *Human) Work() {
	println("Human work")
}

func NewHumanWorker() Worker {
	return &Human{}
}

func NewHumanEater() Eater {
	return &Human{}
}

func NewHumanSleeper() Worker {
	return &Human{}
}

// ------- ROBOT --------------
type Robot struct{}

func (r *Robot) Work() {
	println("Robot work")
}

func NewRobot() Worker {
	return &Robot{}
}

func main42() {
	worker1 := NewHuman()
	worker1.Eat("banana")
	worker1.Work()
	worker1.Sleep()

	worker2 := NewRobot()
	worker2.Work()
}

// -- Pastikan, semua object/fungsi yang menggunakan interface tsb
// -- harus memngimplament SEMUA fungsi yang ada
