//What do you understand by each of the functions demo_func() as shown in the below code?
//DemoStruct definition
type DemoStruct struct {
    Val int
}
//A.
func demo_func() DemoStruct {
    return DemoStruct{Val: 1}
}
//B.
func demo_func() *DemoStruct {
    return &DemoStruct{}
}
//C.
func demo_func(s *DemoStruct) {
    s.Val = 1
}
A. Since the function has a return type of the struct, the function returns a copy of the struct by setting the value as 1.
B. Since the function returns *DemoStruct, which is a pointer to the struct, it returns a pointer to the struct value created within the function.
C. Since the function expects the existing struct object as a parameter and in the function, we are setting the value of its attribute, at the end of execution the value of Val variable of the struct object is set to 1.

