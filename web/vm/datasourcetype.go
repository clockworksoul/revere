package vm

import (
	"database/sql"
	"github.com/yext/revere"
	"github.com/yext/revere/datasources"
)

type DataSourceTypeViewModel struct {
	Type        datasources.DataSourceType
	DataSources []*DataSource
}

const (
	DataSourceDir = "datasources"
)

func NewDataSourceTypeViewModel(db *sql.DB, dst *datasources.DataSourceType) (*DataSourceTypeViewModel, error) {
	dstvm := new(DataSourceTypeViewModel)
	dstvm.Type = *dst

	dataSources, err := revere.LoadDataSourcesOfType(db, dstvm.Type.Id())
	if err != nil {
		return nil, err
	}
	arr := make([]*DataSource, 0)
	for _, ds := range dataSources {
		new, err := NewDataSourceViewModel(ds)
		if err != nil {
			return nil, err
		}
		arr = append(arr, new)
	}
	if len(arr) == 0 {
		new, err := BlankDataSourceViewModelWithType(dstvm.Type.Id())
		if err != nil {
			return nil, err
		}
		arr = append(arr, new)
	}

	dstvm.DataSources = arr
	return dstvm, nil
}