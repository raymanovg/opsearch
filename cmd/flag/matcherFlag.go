package flag

import (
	"fmt"

	"github.com/raymanovg/opsearch/matcher"
)

type MatcherFlag struct {
	Engine matcher.Matcher
}

func (mf *MatcherFlag) String() string {
	if mf.Engine == nil {
		return ""
	}

	return mf.Engine.StringName()
}

func (mf *MatcherFlag) Set(value string) error {
	if match, exist := matcher.Matchers[value]; exist {
		mf.Engine = match
		return nil
	}

	return fmt.Errorf("There is no matcher %s", value)

}
