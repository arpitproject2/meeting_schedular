package services

import (
	"database/sql"
	"meeting_scheduler/internal/db"
	"meeting_scheduler/internal/models"
)

func CreateMeeting(meeting *models.Meeting) error {
	query := `
        INSERT INTO meetings (title, room_id, start_time, end_time, created_by)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	err := db.DB.QueryRow(query, meeting.Title, meeting.RoomID, meeting.StartTime, meeting.EndTime, meeting.CreatedBy).Scan(&meeting.ID)
	return err
}

func GetMeeting(id int) (*models.Meeting, error) {
	query := `
        SELECT id, title, room_id, start_time, end_time, created_by
        FROM meetings
        WHERE id = $1
    `
	meeting := &models.Meeting{}
	err := db.DB.QueryRow(query, id).Scan(&meeting.ID, &meeting.Title, &meeting.RoomID, &meeting.StartTime, &meeting.EndTime, &meeting.CreatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return meeting, nil
}

func UpdateMeeting(meeting *models.Meeting) error {
	query := `
        UPDATE meetings
        SET title = $1, room_id = $2, start_time = $3, end_time = $4, created_by = $5
        WHERE id = $6
    `
	_, err := db.DB.Exec(query, meeting.Title, meeting.RoomID, meeting.StartTime, meeting.EndTime, meeting.CreatedBy, meeting.ID)
	return err
}

func DeleteMeeting(id int) error {
	query := `
        DELETE FROM meetings
        WHERE id = $1
    `
	_, err := db.DB.Exec(query, id)
	return err
}
