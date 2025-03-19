package service

import (
	"dance/models"
	"dance/types"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
)

var Data = []*models.Instructor{
	{ID: 1, Name: "test1"},
	{ID: 2, Name: "test2"},
	{ID: 10, Name: "test10"},
}

type DataSlice []*models.Instructor

func (d DataSlice) Len() int      { return len(d) }
func (d DataSlice) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

func (p DataSlice) Less(i, j int) bool {
	if p[i].ID < p[j].ID {
		return true
	}
	return false
}

var ErrInstructorNotExist = errors.New("The instructor does not exist.")
var ErrInstructorNameExist = errors.New("The name already exists")

type IInstructor interface {
	GetInstance(id int) (*models.Instructor, error)
	ListAllInstances() ([]*models.Instructor, error)
	CreateInstance(obj models.Instructor) error
	DeleteInstance(id int) error
	UpdateInstance(id int, obj types.UpdateInstructorRequest) error
}

type Instructor struct {
}

func NewInstructor() IInstructor {
	return &Instructor{}
}

func (i Instructor) GetInstance(id int) (*models.Instructor, error) {
	index, err := i.getIndexById(id)
	if err != nil {
		return nil, err
	}
	return Data[index], nil
}

func (i Instructor) ListAllInstances() ([]*models.Instructor, error) {
	return Data, nil
}

func (i Instructor) CreateInstance(obj models.Instructor) error {
	ins, err := i.GetInstance(obj.ID)
	if ins != nil || !errors.Is(err, ErrInstructorNotExist) {
		return errors.New("Instructor ID already exists.")
	}

	if err = i.verifyName(obj.Name); err != nil {
		return err
	}

	Data = append(Data, &obj)
	sort.Sort(DataSlice(Data))
	return nil
}

func (i Instructor) DeleteInstance(id int) error {
	index, err := i.getIndexById(id)
	if err != nil {
		return err
	}

	if index == 0 {
		Data = Data[1:]
		return nil
	}

	length := len(Data)
	if index == length-1 {
		Data = Data[:length-1]
		return nil
	}

	Data = append(Data[:index], Data[index+1:]...)
	return nil
}

func (i Instructor) UpdateInstance(id int, obj types.UpdateInstructorRequest) error {
	index, err := i.getIndexById(id)
	if err != nil {
		return err
	}

	if obj.Name != nil && *obj.Name == "" {
		return errors.New("The name cannot be empty.")
	}

	name := types.GetStringVal(obj.Name)
	if err = i.verifyName(name); err != nil {
		return err
	}

	if name != "" {
		Data[index].Name = name
	}

	if obj.Bio != nil {
		Data[index].Bio = obj.Bio
	}

	if obj.Availability != nil {
		Data[index].Availability = obj.Availability
	}

	if obj.Specialty != nil {
		Data[index].Specialty = obj.Specialty
	}

	return nil
}

func (i Instructor) verifyName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty.")
	}
	for _, obj := range Data {
		if obj.Name == name {
			return ErrInstructorNameExist
		}
	}
	return nil
}

func (i Instructor) getIndexById(id int) (int, error) {
	// Binary Search
	left, right := 0, len(Data)-1
	for left <= right {
		mid := left + (right-left)/2
		if Data[mid].ID == id {
			return mid, nil
		} else if Data[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1, ErrInstructorNotExist
}

func (i Instructor) PrintData() {
	b, _ := json.MarshalIndent(Data, "", "\t")
	fmt.Println(string(b))
}

func (i Instructor) PrintInstance(ins *models.Instructor) {
	if ins == nil {
		return
	}
	b, _ := json.MarshalIndent(ins, "", "\t")
	fmt.Println(string(b))
}
