package chihandler

import (
	"net/http"
	"testing"

	port "github.com/devpablocristo/nanlabs/application/port"
)

func TestChiHandler_Task(t *testing.T) {
	type fields struct {
		taskService port.Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ChiHandler{
				taskService: tt.fields.taskService,
			}
			h.Task(tt.args.w, tt.args.r)
		})
	}
}
