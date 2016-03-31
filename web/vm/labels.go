package vm

import (
	"database/sql"
	"fmt"

	"github.com/yext/revere"
)

type Label struct {
	*revere.Label
	Monitors *LabelMonitors
	Triggers []*Trigger
}

func NewLabel(db *sql.DB, id int) (*Label, error) {
	l, err := revere.LoadLabel(db, uint(id))
	if err != nil {
		return nil, err
	}
	if l == nil {
		return nil, fmt.Errorf("Label not found: %d", id)
	}

	return newLabelFromModel(db, l)
}

func newLabelFromModel(db *sql.DB, l *revere.Label) (*Label, error) {
	var err error
	viewmodel := new(Label)
	viewmodel.Label = l
	viewmodel.Triggers, err = NewTriggersFromLabelTriggers(l.Triggers)
	if err != nil {
		return nil, err
	}
	viewmodel.Monitors, err = NewLabelMonitors(db, l.Monitors)
	if err != nil {
		return nil, err
	}
	return viewmodel, nil
}

func newLabelsFromModels(db *sql.DB, rls []*revere.Label) ([]*Label, error) {
	var err error
	labels := make([]*Label, len(rls))
	for i, rl := range rls {
		// TODO(psingh): Batch fetch monitor model data
		labels[i], err = newLabelFromModel(db, rl)
		if err != nil {
			return nil, err
		}
	}
	return labels, nil
}

func BlankLabel(db *sql.DB) (*Label, error) {
	var err error
	viewmodel := new(Label)
	viewmodel.Label = new(revere.Label)
	viewmodel.Triggers = []*Trigger{
		BlankTrigger(),
	}
	viewmodel.Monitors, err = BlankLabelMonitors(db)
	if err != nil {
		return nil, err
	}

	return viewmodel, nil
}

func AllLabels(db *sql.DB) ([]*Label, error) {
	rls, err := revere.LoadLabels(db)
	if err != nil {
		return nil, err
	}

	return newLabelsFromModels(db, rls)
}