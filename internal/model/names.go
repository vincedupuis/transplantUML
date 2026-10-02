package model

import (
	"regexp"
	"strings"
)

// Guards, actions, do activities, invariants and named delays hold names
// only, which the code a template generates calls as functions. A name is an identifier:
// letters, digits and underscores, not starting with a digit.
var name = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

// IsName reports whether s is a name that an action or a do activity may
// hold.
func IsName(s string) bool { return name.MatchString(s) && !isOperator(s) }

// IsDelay reports whether s is a delay that a time trigger may hold: a
// duration in seconds or milliseconds, such as 5s, 1.5s or 250ms, or a name.
func IsDelay(s string) bool { return duration.MatchString(s) || IsName(s) }

var duration = regexp.MustCompile(`^[0-9]+(\.[0-9]+)?(ms|s)$`)

// IsCondition reports whether s is a condition that a guard or an invariant
// may hold: names joined by and, or, not and parentheses, as in
// "card and not (blocked or expired)". A name may be negated once.
func IsCondition(s string) bool {
	c := &condition{tokens: conditionTokens.FindAllString(s, -1)}
	if strings.Join(c.tokens, "") != strings.Join(strings.Fields(s), "") {
		return false // s holds something other than names, operators and parentheses
	}
	return c.or() && c.pos == len(c.tokens)
}

var conditionTokens = regexp.MustCompile(`[A-Za-z_][A-Za-z0-9_]*|[()]`)

func isOperator(s string) bool { return s == "and" || s == "or" || s == "not" }

// condition parses the tokens of a condition by recursive descent: or binds
// looser than and, and not applies to one name or parenthesized condition.
type condition struct {
	tokens []string
	pos    int
}

func (c *condition) peek() string {
	if c.pos < len(c.tokens) {
		return c.tokens[c.pos]
	}
	return ""
}

func (c *condition) or() bool {
	for ok := c.and(); ok; ok = c.and() {
		if c.peek() != "or" {
			return true
		}
		c.pos++
	}
	return false
}

func (c *condition) and() bool {
	for ok := c.not(); ok; ok = c.not() {
		if c.peek() != "and" {
			return true
		}
		c.pos++
	}
	return false
}

func (c *condition) not() bool {
	if c.peek() == "not" {
		c.pos++
	}
	return c.single()
}

func (c *condition) single() bool {
	switch tok := c.peek(); {
	case tok == "(":
		c.pos++
		if !c.or() || c.peek() != ")" {
			return false
		}
		c.pos++
		return true
	case IsName(tok):
		c.pos++
		return true
	}
	return false
}
