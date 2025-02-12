// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type FeaturesListItem string

const (
	FeaturesListItemArtifacts   FeaturesListItem = "ARTIFACTS"
	FeaturesListItemJunitReport FeaturesListItem = "JUNIT_REPORT"
)

var AllFeaturesListItem = []FeaturesListItem{
	FeaturesListItemArtifacts,
	FeaturesListItemJunitReport,
}

func (e FeaturesListItem) IsValid() bool {
	switch e {
	case FeaturesListItemArtifacts, FeaturesListItemJunitReport:
		return true
	}
	return false
}

func (e FeaturesListItem) String() string {
	return string(e)
}

func (e *FeaturesListItem) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FeaturesListItem(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FeaturesListItem", str)
	}
	return nil
}

func (e FeaturesListItem) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
