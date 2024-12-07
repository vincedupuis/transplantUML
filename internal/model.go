package internal

type StateType string

const (
	Normal      = "N"
	InitialType = "I"
	HistoryType = "H"
	FinalType   = "F"
)

type StateMachine struct {
	Name        string
	States      []*State
	Transitions []*Transition
}

type State struct {
	Name   string
	Parent string
}

type Transition struct {
	Source     string
	SourceType StateType
	Event      string
	Cond       string
	Action     string
	Target     string
	DestType   StateType
}

func (sm *StateMachine) GetInnerStates(name string) []*State {
	states := make([]*State, 0)
	for _, state := range sm.States {
		if state.Parent == name {
			states = append(states, state)
		}
	}
	return states
}

func (sm *StateMachine) GetOutgoingTransitions(name string, sourceType, destType StateType) []*Transition {
	transitions := make([]*Transition, 0)
	for _, transition := range sm.Transitions {
		if transition.Source == name && transition.SourceType == sourceType && transition.DestType == destType {
			transitions = append(transitions, transition)
		}
	}
	return transitions
}

func (sm *StateMachine) GetIncomingTransitions(name string, sourceType, destType StateType) []*Transition {
	transitions := make([]*Transition, 0)
	for _, transition := range sm.Transitions {
		if transition.Target == name && transition.SourceType == sourceType && transition.DestType == destType {
			transitions = append(transitions, transition)
		}
	}
	return transitions
}
