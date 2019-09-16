package StatePattern

import "fmt"

type State interface {
    DoAction(context *Context)
    ToString() string
}

type StartState struct {
}

func (s *StartState) DoAction(context *Context) {
    fmt.Println("Play is in start state")
    context.SetState(s)
}

func (s *StartState) ToString() string {
    return "Start State"
}

type StopState struct {
}

func (s *StopState) DoAction(context *Context) {
    fmt.Println("Play is in stop state")
    context.SetState(s)
}

func (s *StopState) ToString() string {
    return "Stop State"
}

//-------------------------------------------------------

type Context struct {
    state State
}

func (c *Context) Context() {
    c.state = nil
}

func (c *Context) SetState(state State) {
    c.state = state
}

func (c *Context) GetState() State {
    return c.state
}

//-------------------------------------------------------
//调用
func Run()  {

    context:=new(Context)
    startState:=new(StartState)
    startState.DoAction(context)
    fmt.Println(context.GetState().ToString())

    stopState:=new(StopState)
    stopState.DoAction(context)

    fmt.Println(context.GetState().ToString())

}