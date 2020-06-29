package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title      string    `gorm:"size:255;not null;unique" json:"title"`
	Content    string    `gorm:"size:255;not null;" json:"content"`
	EndingAt   time.Time `gorm:"default:NULL" json:"ending_at"`
	Finished   bool      `gorm: "default:FALSE" json:"finished"`
	SubtaskIDs []uint64
	Author     User      `json:"author"`
	AuthorID   uint32    `gorm:"not null" json:"author_id"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (t *Task) Prepare() {
	t.ID = 0
	t.Title = html.EscapeString(strings.TrimSpace(t.Title))
	t.Content = html.EscapeString(strings.TrimSpace(t.Content))
	t.Author = User{}
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}

func (t *Task) Validate() error {

	if t.Title == "" {
		return errors.New("Required Title")
	}
	if t.Content == "" {
		return errors.New("Required Content")
	}
	if t.AuthorID < 1 {
		return errors.New("Required Author")
	}
	return nil
}

func (t *Task) SaveTask(db *gorm.DB) (*Task, error) {
	var err error
	err = db.Debug().Model(&Task{}).Create(&t).Error
	if err != nil {
		return &Task{}, err
	}
	if t.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", t.AuthorID).Take(&t.Author).Error
		if err != nil {
			return &Task{}, err
		}
	}
	return t, nil
}

func (t *Task) FindAllTasks(db *gorm.DB) (*[]Task, error) {
	var err error
	tasks := []Task{}
	err = db.Debug().Model(&Task{}).Limit(100).Find(&tasks).Error
	if err != nil {
		return &[]Task{}, err
	}
	if len(tasks) > 0 {
		for i := range tasks {
			err := db.Debug().Model(&User{}).Where("id = ?", tasks[i].AuthorID).Take(&tasks[i].Author).Error
			if err != nil {
				return &[]Task{}, err
			}
		}
	}
	return &tasks, nil
}

func (t *Task) FindTaskByID(db *gorm.DB, pid uint64) (*Task, error) {
	var err error
	err = db.Debug().Model(&Task{}).Where("id = ?", pid).Take(&t).Error
	if err != nil {
		return &Task{}, err
	}
	if t.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", t.AuthorID).Take(&t.Author).Error
		if err != nil {
			return &Task{}, err
		}
	}
	return t, nil
}

func (t *Task) UpdateATask(db *gorm.DB, pid uint64) (*Task, error) {

	var err error
	db = db.Debug().Model(&Task{}).Where("id = ?", pid).Take(&Task{}).UpdateColumns(
		map[string]interface{}{
			"title":      t.Title,
			"content":    t.Content,
			"updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&t).Error
	if err != nil {
		return &Task{}, err
	}
	if t.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", t.AuthorID).Take(&t.Author).Error
		if err != nil {
			return &Task{}, err
		}
	}
	return t, nil
}

func (t *Task) DeleteATask(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Task{}).Where("id = ? and author_id = ?", pid, uid).Take(&Task{}).Delete(&Task{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
