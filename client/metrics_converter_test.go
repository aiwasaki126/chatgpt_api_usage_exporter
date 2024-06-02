package client

import (
	"reflect"
	"testing"

	"github.com/aiwasaki126/chat_gpt_api_usage_exporter/client/response"
	"github.com/aiwasaki126/chat_gpt_api_usage_exporter/collector"
)

var (
	label1 = collector.Labels{
		OrganizationID:   "org1",
		OrganizationName: "org1",
		Operation:        "operation1",
		SnapshotID:       "snapshot1",
	}
	label2 = collector.Labels{
		OrganizationID:   "org2",
		OrganizationName: "org2",
		Operation:        "operation2",
		SnapshotID:       "snapshot2",
	}
)

func Test_metricsConverter_convertDataToMetrics(t *testing.T) {
	type fields struct {
		body *response.UsageResponse
	}
	type args struct {
		data []*rawMetrics
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*collector.Metrics
		wantErr bool
	}{
		// TODO: Add test cases.
		{"empty data", fields{nil}, args{[]*rawMetrics{}}, []*collector.Metrics{}, false},
		{"single data", fields{nil}, args{[]*rawMetrics{{1, 3, 5, label1}}}, []*collector.Metrics{{1, 1, 3, 5, label1}}, false},
		{"multiple data with the same labels", fields{nil}, args{[]*rawMetrics{{1, 3, 5, label1}, {1, 3, 5, label1}}}, []*collector.Metrics{{1, 2, 6, 10, label1}}, false},
		{"multiple data with different labels", fields{nil}, args{[]*rawMetrics{{1, 3, 5, label1}, {10, 30, 50, label2}}}, []*collector.Metrics{{1, 1, 3, 5, label1}, {1, 10, 30, 50, label2}}, false},
		{"multiple data with different labels with the same labels", fields{nil}, args{[]*rawMetrics{{1, 3, 5, label1}, {10, 30, 50, label2}, {1, 3, 5, label1}}}, []*collector.Metrics{{1, 2, 6, 10, label1}, {1, 10, 30, 50, label2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &metricsConverter{
				body: tt.fields.body,
			}
			got, err := m.convertDataToMetrics(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("metricsConverter.convertDataToMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("metricsConverter.convertDataToMetrics() = %v, want %v", got, tt.want)
			}
		})
	}
}
