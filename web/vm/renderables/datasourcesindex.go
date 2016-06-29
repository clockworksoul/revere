package renderables

import (
	"github.com/yext/revere/datasources"
	"github.com/yext/revere/web/vm"
)

type DataSourceIndex struct {
	datasources []*datasources.VM
	subs        []Renderable
}

func NewDataSourceIndex(dss []*datasources.VM) *DataSourceIndex {
	dsi := DataSourceIndex{}
	dsi.datasources = dss
	dsi.subs = make([]Renderable, len(dss))
	for i, ds := range dss {
		dsi.subs[i] = NewDataSourceView(ds)
	}
	return &dsi
}

func (dsi *DataSourceIndex) name() string {
	return "Data Sources"
}

func (dsi *DataSourceIndex) template() string {
	return "datasources-index.html"
}

func (dsi *DataSourceIndex) data() interface{} {
	return datasources.AllTypes()
}

func (dsi *DataSourceIndex) scripts() []string {
	return datasources.AllScripts()
}

func (dsi *DataSourceIndex) breadcrumbs() []vm.Breadcrumb {
	return nil
}

func (dsi *DataSourceIndex) subRenderables() []Renderable {
	return dsi.subs
}

func (dsi *DataSourceIndex) renderPropagate() (*renderResult, error) {
	return renderPropagate(dsi)
}

func (dsi *DataSourceIndex) aggregatePipelineData(parent *renderResult, child *renderResult) {
	aggregatePipelineDataArray(parent, child)
}
