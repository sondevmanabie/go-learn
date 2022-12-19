package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"rest-api/ent"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockFileReader interface {
	Read(arg interface{}) (interface{}, error)
	Write(arg interface{}) error
}
type MockEntityReader struct {
	readFn  func(arg interface{}) (interface{}, error)
	writeFn func(arg interface{}) error
}

func (f *MockEntityReader) Read(arg interface{}) (interface{}, error) {
	return f.readFn(arg)
}
func (f *MockEntityReader) Write(arg interface{}) error {
	return f.writeFn(arg)
}
func TestStudentController_Get(t *testing.T) {
	tests := []struct {
		name    string
		want    interface{}
		wantErr bool
	}{
		{
			name: "Should response student data correctly",
			want: []ent.Student{{
				Id:     "1",
				Name:   "Son",
				Age:    23,
				Active: true,
			}},
			wantErr: false,
		},
		{
			name:    "Should response bad request",
			want:    errors.New("read error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			mockStudentFileReader := MockEntityReader{
				readFn: func(arg interface{}) (interface{}, error) {
					if tt.wantErr {
						return nil, errors.New("read error")
					}
					return &[]ent.Student{{
						Id:     "1",
						Name:   "Son",
						Age:    23,
						Active: true,
					}}, nil
				},
			}
			c := &StudentController{
				FileReader: &mockStudentFileReader,
			}
			router.GET("/student", c.Get)
			req, _ := http.NewRequest("GET", "/student", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			var students []ent.Student
			json.Unmarshal(w.Body.Bytes(), &students)
			// assert.NoError(t, err)
			if tt.wantErr {
				assert.Equalf(t, http.StatusNotFound, w.Code, "StudentController.Get() receive= %v, expect=%v", w.Code, http.StatusNotFound)
			} else {
				assert.Equalf(t, tt.want, students, "StudentController.Get() receive= %v, expect=%v", students, tt.want)
				assert.Equalf(t, http.StatusOK, w.Code, "StudentController.Get() receive= %v, expect=%v", w.Code, http.StatusOK)
			}
		})
	}
}

func TestStudentController_Add(t *testing.T) {
	tests := []struct {
		name    string
		want    interface{}
		actual  interface{}
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			mockStudentFileReader := MockEntityReader{
				readFn: func(arg interface{}) (interface{}, error) {
					return &[]ent.Student{}, nil
				},
				writeFn: func(arg interface{}) error {
					tt.actual = arg
					return nil
				},
			}
			c := &StudentController{
				FileReader: &mockStudentFileReader,
			}
			router.POST("/student", c.Add)
			jsonStr := []byte(`{"Id":"1","Name":"Son","Age":23,"Active":true}`)
			req, _ := http.NewRequest("POST", "/student", bytes.NewBuffer(jsonStr))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if tt.wantErr {
				assert.Equalf(t, http.StatusNotFound, w.Code, "StudentController.Add() receive= %v, expect=%v", w.Code, http.StatusNotFound)
			} else {
				assert.Equalf(t, http.StatusOK, w.Code, "StudentController.Add() receive= %v, expect=%v", w.Code, http.StatusOK)
				assert.Equalf(t, tt.want, tt.actual, "StudentController.Add() receive= %v, expect=%v", tt.actual, tt.want)
			}
		})
	}

}
