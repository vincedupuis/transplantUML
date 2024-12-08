package internal

type StateType string

const (
	Normal      = "N"
	InitialType = "I"
	HistoryType = "H"
	FinalType   = "F"
)

type StateMachine struct {
	Name        string        `json:"name"`
	States      []*State      `json:"states"`
	Transitions []*Transition `json:"transitions"`
}

type State struct {
	Name   string `json:"name"`
	Parent string `json:"parent"`
}

type Transition struct {
	Source     string    `json:"source"`
	SourceType StateType `json:"sourceType"`
	Event      string    `json:"event"`
	Cond       string    `json:"cond"`
	Action     string    `json:"action"`
	Target     string    `json:"target"`
	DestType   StateType `json:"destType"`
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
