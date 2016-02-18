package vm

import (
	"html/template"

	"github.com/yext/revere"
	"github.com/yext/revere/probes"
)

type Monitor struct {
	*revere.Monitor
	Probe    *Probe
	Triggers []*revere.MonitorTrigger
}

func NewMonitor(m *revere.Monitor) (*Monitor, error) {
	viewmodel := new(Monitor)

	viewmodel.Monitor = m
	viewmodel.Triggers = m.Triggers

	probeType, err := probes.ProbeTypeById(m.ProbeType)
	if err != nil {
		return nil, err
	}

	probe, err := probeType.Load(m.ProbeJson)
	if err != nil {
		return nil, err
	}
	viewmodel.Probe = NewProbe(probe)

	return viewmodel, nil
}

func BlankMonitor() (*Monitor, error) {
	viewmodel := new(Monitor)

	viewmodel.Monitor = new(revere.Monitor)

	viewmodel.Triggers = []*revere.MonitorTrigger{
		&revere.MonitorTrigger{
			Trigger: revere.Trigger{
				TargetTemplate: template.HTML("PLACEHOLDER DELETE"), //TODO(fchen): code cleanup[targets.DefaultTargetTemplate()]
			},
		},
	}

	viewmodel.Probe = DefaultProbe()

	return viewmodel, nil
}

func (m *Monitor) GetProbeType() probes.ProbeType {
	return m.Probe.ProbeType()
}
