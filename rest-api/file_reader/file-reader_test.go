package file_reader

import (
	"errors"
	"reflect"
	"rest-api/ent"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntityReader_Read(t *testing.T) {
	type fields struct {
		Entity string
		Path   string
	}
	tests := []struct {
		name       string
		fields     fields
		want       interface{}
		wantErr    bool
		wantErrMsg error
	}{
		{
			name: "Should read student file correctly",
			fields: fields{
				Entity: "student",
				Path:   "../data-test",
			},
			want: []ent.Student{
				{
					Id:     "1",
					Name:   "Son",
					Age:    23,
					Active: true},
			},
			wantErr: false,
		},
		{
			name: "Should read book file correctly",
			fields: fields{
				Entity: "book",
				Path:   "../data-test",
			},
			want: []ent.Book{
				{
					Id:      "1",
					Title:   "Dune",
					Author:  "Son",
					Publish: false,
				},
			},
			wantErr: false,
		},
		{
			name: "Should return nil when file not found",
			fields: fields{
				Entity: "course",
				Path:   "data-test",
			},
			want:       nil,
			wantErr:    true,
			wantErrMsg: errors.New("not match any entity"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EntityReader{
				Entity: tt.fields.Entity,
				Path:   tt.fields.Path,
			}
			var arg interface{}
			if tt.fields.Entity == "student" {
				arg = &[]ent.Student{}
			} else {
				arg = &[]ent.Book{}
			}
			got, err := r.Read(arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("EntityReader.Read() receive = %v, expect %v", err, tt.wantErr)
				return
			}
			var receive interface{}
			if tt.fields.Entity == "student" {
				receive = *got.(*[]ent.Student)
			} else if tt.fields.Entity == "book" {
				receive = *got.(*[]ent.Book)
			} else {
				receive = got
			}
			if !reflect.DeepEqual(receive, tt.want) {
				t.Errorf("EntityReader.Read() receive = %v, expect %v", receive, tt.want)
			}
			if tt.wantErrMsg != nil {
				assert.Containsf(t, err.Error(), tt.wantErrMsg.Error(), "EntityReader.Read() receive error = %q, expect %q", tt.wantErrMsg, err)
			}
		})
	}
}

func TestEntityReader_Write(t *testing.T) {
	type fields struct {
		Entity string
		Path   string
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		wantErrMsg error
	}{
		{
			name: "Should write student data correctly",
			fields: fields{
				Entity: "student",
				Path:   "../data-test",
			},
			args: args{
				data: []ent.Student{{
					Id:     "1",
					Name:   "Son",
					Age:    23,
					Active: true,
				}},
			},
			wantErr: false,
		},
		{
			name: "Should write book data correctly",
			fields: fields{
				Entity: "book",
				Path:   "../data-test",
			},
			args: args{
				data: []ent.Book{{
					Id:      "1",
					Title:   "Dune",
					Author:  "Son",
					Publish: false,
				}},
			},
			wantErr: false,
		},
		{
			name: "Should throw error when file not found",
			fields: fields{
				Entity: "course",
				Path:   "../data-test",
			},
			args: args{
				data: nil,
			},
			wantErr:    true,
			wantErrMsg: errors.New("not match any entity"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EntityReader{
				Entity: tt.fields.Entity,
				Path:   tt.fields.Path,
			}
			err := r.Write(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("EntityReader.Write() error = %v, wantErr %v", err, tt.wantErrMsg)
			}
			if tt.wantErrMsg != nil {
				assert.Containsf(t, err.Error(), tt.wantErrMsg.Error(), "EntityReader.Write() receive error = %q, expect %q", tt.wantErrMsg, err)
			}
		})
	}
}
