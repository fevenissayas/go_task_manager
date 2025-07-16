package data

import "restfulapi/models"


func Add(task *models.Task) bool {
	for _,existingTask := range models.Tasks {
		if task.ID == existingTask.ID {
			return false
		}
	}
	models.Tasks = append(models.Tasks, task)
	return true
}


func GetAll() []*models.Task{
	return models.Tasks
}

func Get(id string) (*models.Task, bool) {
	for _, task := range models.Tasks {
		if task.ID == id {
			return task, true
		}
	}
	return nil, false
}


func Delete(id string) bool{
	for i, task := range models.Tasks {
		if task.ID == id {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			return true
		}
	}
	return false

}

func Update(id string, updatedTask *models.Task) bool {
	for i, task := range models.Tasks {
		if task.ID == id {
			models.Tasks[i].Title = updatedTask.Title
			models.Tasks[i].Description = updatedTask.Description
			models.Tasks[i].DueDate = updatedTask.DueDate
			models.Tasks[i].Status = updatedTask.Status
			return true
		}
	}
	return false
}