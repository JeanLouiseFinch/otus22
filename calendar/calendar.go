package calendar

import (
	"fmt"
	"otus22/internal/sql"
	"time"
)

type Calendar struct {
	ID int
}

type Event struct {
	Title       string
	Description string
	Start       time.Time
	End         time.Time
}

func NewCalendar() *Calendar {
	c := &Calendar{}
	return c
}

func (c *Calendar) NewEvent(title, description string, start time.Time, end time.Time) (string, error) {
	result, err := sql.NewEvent(c.ID, title, description, start, end)
	return result, err
}
func (c *Calendar) String() string {
	eventsSQL, err := sql.GetAllEvents(c.ID)
	if err != nil {
		return ""
	}
	var result string
	result = fmt.Sprintf("Calendar:\n---\n")
	for key, val := range eventsSQL {
		result += fmt.Sprintf("\tEvent->\n\t\tID: %v\n\t\tTitle: %s\n\t\tDescription: %s\n\t\tStart: %v\n\t\tEnd: %v\n", key, val.Title, val.Description, val.Start, val.End)
	}
	result += fmt.Sprintf("---\n")
	return result
}

func (c *Calendar) GetEvent(id string) (Event, error) {
	result, err := sql.GetEvent(id)
	if err != nil {
		return Event{}, err
	}
	ev := Event{
		Title:       result.Title,
		Description: result.Description,
		Start:       result.Start,
		End:         result.End,
	}
	return ev, nil
}

func (c *Calendar) ModifyEvent(id string, e Event) error {
	_, err := sql.ModifyEvent(id, e.Title, e.Description, e.Start, e.End)
	return err
}

func (c *Calendar) RemoveEvent(id string) error {
	return sql.RemoveEvent(id)
}
