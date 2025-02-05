package main

// Interface Segregation
// -- Bad
// type Worker interface {
// 	Eat(food string)
// 	Sleep()
// 	Work()
// }

// type Human struct {
// }

// func (h *Human) Eat(food string) {
// 	println("Human eat", food)
// }

// func (h *Human) Sleep() {
// 	println("Human sleep")
// }

// func (h *Human) Work() {
// 	println("Human work")
// }

// func NewHuman() Worker {
// 	return &Human{}
// }

// // ------- ROBOT --------------
// type Robot struct{}

// func (r *Robot) Eat() {
// 	//
// }

// func (r *Robot) Work() {
// 	//
// }

// func (r *Robot) Work() {
// 	println("Robot work")
// }

// func NewRobot() Worker {
// 	return &Robot{}
// }

// func main4() {
// 	worker1 := NewHuman()
// 	worker1.Eat("banana")
// 	worker1.Work()
// 	worker1.Sleep()

// 	worker2 := NewRobot()
// 	worker2.Work()
// }
