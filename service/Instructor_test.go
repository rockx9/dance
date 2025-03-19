package service

import (
	"dance/models"
	"testing"
)

func TestInstructor_CreateInstance(t *testing.T) {
	type args struct {
		obj *models.Instructor
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				obj: &models.Instructor{
					ID:   10,
					Name: "test1",
				},
			},
		},
		{
			name: "",
			args: args{
				obj: &models.Instructor{
					ID:   5,
					Name: "test2",
				},
			},
		},
		{
			name: "",
			args: args{
				obj: &models.Instructor{
					ID:   15,
					Name: "test3",
				},
			},
		},
		{
			name: "",
			args: args{
				obj: &models.Instructor{
					ID:   2,
					Name: "test4",
				},
			},
		},
	}
	i := Instructor{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := i.CreateInstance(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("CreateInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInstructor_GetInstance(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Instructor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				id: 1,
			},
		},
		{
			args: args{
				id: 10,
			},
		},
		{
			args: args{
				id: 11,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Instructor{}
			got, err := i.GetInstance(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			i.PrintInstance(got)
		})
	}
}

func TestInstructor_DeleteInstance(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				id: 1,
			},
		},
		{
			args: args{
				id: 20,
			},
		},
	}
	i := Instructor{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := i.DeleteInstance(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	i.PrintData()
}

func TestInstructor_UpdateInstance(t *testing.T) {
	type args struct {
		id  int
		obj *models.Instructor
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				id: 1,
				obj: &models.Instructor{
					Name: "updated_name",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Instructor{}
			if err := i.UpdateInstance(tt.args.id, tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("UpdateInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
			i.PrintData()
		})
	}

}
