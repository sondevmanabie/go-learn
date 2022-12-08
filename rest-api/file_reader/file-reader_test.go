package file_reader

import (
	"reflect"
	"rest-api/ent"
	"testing"
)

func TestEntityReader_Read(t *testing.T) {
	type fields struct {
		Entity string
		Path   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
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
					Title:   "The Hobbit",
					Author:  "Son",
					Publish: true,
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
			want:    nil,
			wantErr: true,
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
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should write data correctly",
			fields: fields{
				Entity: "student",
				Path:   "../data-test",
			},
			args: args{
				data: ent.Student{
					Id:     "1",
					Name:   "Son",
					Age:    23,
					Active: true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EntityReader{
				Entity: tt.fields.Entity,
				Path:   tt.fields.Path,
			}
			if err := r.Write(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("EntityReader.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
