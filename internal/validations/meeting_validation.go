package validations

import (
	"errors"

	"github.com/arpitproject2/meeting_schedular/internal/models"
)

func ValidateMeeting(meeting *models.Meeting) error {
	if meeting.Title == "" {
		return errors.New("meeting title is required")
	}
	if meeting.StartTime.After(meeting.EndTime) {
		return errors.New("start time cannot be after end time")
	}
	return nil
}
