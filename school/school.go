package school

import (
	"errors"
)

type Grade struct {
	Student, Class string
	Value float64
}


type School struct {
	Students map[string]map[string]float64
}

func getAverage(student map[string]float64) float64 {
	var total float64 = 0

	for _, grade := range student {
		total += grade
	}

	return total / float64(len(student))
}

func (t *School) AddGrade(grade *Grade, reply *float64) error {
	if t.Students[grade.Student] == nil {
		t.Students[grade.Student] = make(map[string]float64)
	}

	t.Students[grade.Student][grade.Class] = grade.Value
	return nil
} 

func (t *School) GetStudentAverage(name string, reply *float64) error {
	student, ok := t.Students[name]
	if !ok {
		return errors.New("El usuario no se encontro")
	}

	*reply = getAverage(student)

	return nil
}

func (t *School) GetGeneralAverage(a bool, reply *float64) error {
	var total float64
	for _, student := range t.Students {
		total += getAverage(student)
	}
	if len(t.Students) > 0 {
		*reply = total / float64(len(t.Students))
	} else {
		*reply = 0
	}

	return nil
}

func (t *School) GetClassAverage(name string, reply *float64) error {
	var total float64
	var found bool
	var totalClasses int
	for _, student := range t.Students {
		grade, ok := student[name]
		if ok {
			total += grade
			totalClasses += 1
			found = true
		}
	}

	if found {
		*reply = total / float64(totalClasses)
	} else {
		return errors.New("No se encontro la clase")
	}
	return nil
}