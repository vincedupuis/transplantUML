package internal

import (
	"errors"
	"github.com/beevik/etree"
)

func ScxmlToStateMachine(xml string) (*StateMachine, error) {
	doc := etree.NewDocument()
	err := doc.ReadFromString(xml)
	if err != nil {
		return nil, errors.New("error parsing XML: %s" + err.Error())
	}

	scxmlElement := doc.SelectElement("scxml")
	if scxmlElement == nil {
		return nil, errors.New("error parsing SCXML: no SCXML element found")
	}

	states, transitions := getAllStatesAndTransitions(scxmlElement)

	sm := &StateMachine{
		Name:        getXmlAttribute(scxmlElement, "name"),
		States:      states,
		Transitions: transitions,
	}

	return sm, nil
}

type mapOfStateNameToParentName map[string]string

func getAllStatesAndTransitions(element *etree.Element) ([]*State, []*Transition) {
	states := make([]*State, 0)
	transitions := make([]*Transition, 0)

	historyStates := createSpecialStateMapping(element, "history")
	finalStates := createSpecialStateMapping(element, "final")

	initialTransition := getInitialTransition(element, historyStates)
	if initialTransition != nil {
		transitions = append(transitions, initialTransition)
	}

	stateElements := element.FindElements("//state")
	for _, stateElement := range stateElements {
		state := &State{
			Parent: getXmlAttribute(stateElement.Parent(), "id"),
			Name:   getXmlAttribute(stateElement, "id"),
		}
		states = append(states, state)
		initialTransition := getInitialTransition(stateElement, historyStates)
		if initialTransition != nil {
			transitions = append(transitions, initialTransition)
		}
	}

	transitionElements := element.FindElements("//transition")
	for _, transitionElement := range transitionElements {
		source, sourceType := determineSourceState(transitionElement, historyStates)
		target, targetType := determineTargetState(transitionElement, historyStates, finalStates)
		transition := &Transition{
			Source:     source,
			SourceType: sourceType,
			Event:      getXmlAttribute(transitionElement, "event"),
			Cond:       getXmlAttribute(transitionElement, "cond"),
			Action:     getXmlText(transitionElement, "script"),
			Target:     target,
			DestType:   targetType,
		}
		transitions = append(transitions, transition)
	}

	return states, transitions
}

func createSpecialStateMapping(element *etree.Element, stateType string) mapOfStateNameToParentName {
	mapping := make(mapOfStateNameToParentName)
	for _, child := range element.FindElements("//" + stateType) {
		name := getXmlAttribute(child, "id")
		mapping[name] = getXmlAttribute(child.Parent(), "id")
	}
	return mapping
}

func getInitialTransition(element *etree.Element, historyStates mapOfStateNameToParentName) *Transition {
	initial := element.SelectAttrValue("initial", "")
	if initial == "" {
		firstInnerState := element.SelectElement("state")
		if firstInnerState == nil {
			return nil
		}
		initial = getXmlAttribute(firstInnerState, "id")
	}

	var targetType StateType = Normal
	if _, exist := historyStates[initial]; exist {
		targetType = HistoryType
	}

	return &Transition{
		Source:     getXmlAttribute(element, "id"),
		SourceType: InitialType,
		Event:      "",
		Cond:       "",
		Action:     "",
		Target:     initial,
		DestType:   targetType,
	}
}

func determineTargetState(element *etree.Element, historyStates mapOfStateNameToParentName, finalStates mapOfStateNameToParentName) (string, StateType) {
	target := getXmlAttribute(element, "target")
	if stateName, exist := historyStates[target]; exist {
		return stateName, HistoryType
	}
	if stateName, exist := finalStates[target]; exist {
		return stateName, FinalType
	}
	return target, Normal
}

func determineSourceState(element *etree.Element, historyStates mapOfStateNameToParentName) (string, StateType) {
	source := getXmlAttribute(element.Parent(), "id")
	if name, exist := historyStates[source]; exist {
		return name, HistoryType
	}
	return source, Normal
}

func getXmlAttribute(element *etree.Element, name string) string {
	if element == nil {
		return ""
	}
	return element.SelectAttrValue(name, "")
}

func getXmlText(element *etree.Element, s string) string {
	el := element.SelectElement(s)
	if el != nil {
		return el.Text()
	}
	return ""
}
