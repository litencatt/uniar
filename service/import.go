package service

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/litencatt/uniar/repository"
)

type ImportService struct {
	DB      *sql.DB
	Querier repository.Querier
}

type ImportResult struct {
	Success      int                    `json:"success"`
	Failed       int                    `json:"failed"`
	Errors       []ImportError          `json:"errors"`
	PreviewData  []map[string]interface{} `json:"preview_data,omitempty"`
}

type ImportError struct {
	Row     int    `json:"row"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

type MusicImportRow struct {
	Name        string `csv:"name"`
	Normal      string `csv:"normal"`
	Pro         string `csv:"pro"`
	Master      string `csv:"master"`
	Length      string `csv:"length"`
	ColorTypeID string `csv:"color_type_id"`
	LiveID      string `csv:"live_id"`
}

type PhotographImportRow struct {
	Name      string `csv:"name"`
	GroupID   string `csv:"group_id"`
	PhotoType string `csv:"photo_type"`
}

type SceneImportRow struct {
	PhotographID    string `csv:"photograph_id"`
	MemberID        string `csv:"member_id"`
	ColorTypeID     string `csv:"color_type_id"`
	VocalMax        string `csv:"vocal_max"`
	DanceMax        string `csv:"dance_max"`
	PerformanceMax  string `csv:"performance_max"`
	CenterSkill     string `csv:"center_skill"`
	ExpectedValue   string `csv:"expected_value"`
	SSRPlus         string `csv:"ssr_plus"`
}

func (s *ImportService) ImportMusicFromCSV(ctx context.Context, reader io.Reader, validateOnly bool) (*ImportResult, error) {
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	if len(records) == 0 {
		return &ImportResult{Success: 0, Failed: 0, Errors: []ImportError{}}, nil
	}

	header := records[0]
	expectedFields := []string{"name", "normal", "pro", "master", "length", "color_type_id", "live_id"}
	if err := validateHeader(header, expectedFields); err != nil {
		return nil, err
	}

	result := &ImportResult{
		Success: 0,
		Failed:  0,
		Errors:  []ImportError{},
	}

	if validateOnly {
		result.PreviewData = make([]map[string]interface{}, 0)
	}

	var tx *sql.Tx
	if !validateOnly {
		var err error
		tx, err = s.DB.BeginTx(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to begin transaction: %w", err)
		}
		defer func() {
			_ = tx.Rollback() // ロールバックエラーは無視
		}()
	}

	for i, record := range records[1:] {
		rowNum := i + 2

		if len(record) < len(expectedFields) {
			result.Failed++
			result.Errors = append(result.Errors, ImportError{
				Row:     rowNum,
				Field:   "record",
				Message: fmt.Sprintf("not enough fields: expected %d, got %d", len(expectedFields), len(record)),
			})
			continue
		}

		musicRow := MusicImportRow{
			Name:        strings.TrimSpace(record[0]),
			Normal:      strings.TrimSpace(record[1]),
			Pro:         strings.TrimSpace(record[2]),
			Master:      strings.TrimSpace(record[3]),
			Length:      strings.TrimSpace(record[4]),
			ColorTypeID: strings.TrimSpace(record[5]),
			LiveID:      strings.TrimSpace(record[6]),
		}

		if validateOnly {
			previewData := map[string]interface{}{
				"name":          musicRow.Name,
				"normal":        musicRow.Normal,
				"pro":           musicRow.Pro,
				"master":        musicRow.Master,
				"length":        musicRow.Length,
				"color_type_id": musicRow.ColorTypeID,
				"live_id":       musicRow.LiveID,
			}
			result.PreviewData = append(result.PreviewData, previewData)
			result.Success++
			continue
		}

		if err := s.validateAndInsertMusic(ctx, tx, musicRow, rowNum, result); err != nil {
			continue
		}

		result.Success++
	}

	if !validateOnly && result.Failed == 0 {
		if err := tx.Commit(); err != nil {
			return nil, fmt.Errorf("failed to commit transaction: %w", err)
		}
	}

	return result, nil
}

func (s *ImportService) ImportPhotographFromCSV(ctx context.Context, reader io.Reader, validateOnly bool) (*ImportResult, error) {
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	if len(records) == 0 {
		return &ImportResult{Success: 0, Failed: 0, Errors: []ImportError{}}, nil
	}

	header := records[0]
	expectedFields := []string{"name", "group_id", "photo_type"}
	if err := validateHeader(header, expectedFields); err != nil {
		return nil, err
	}

	result := &ImportResult{
		Success: 0,
		Failed:  0,
		Errors:  []ImportError{},
	}

	if validateOnly {
		result.PreviewData = make([]map[string]interface{}, 0)
	}

	var tx *sql.Tx
	if !validateOnly {
		var err error
		tx, err = s.DB.BeginTx(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to begin transaction: %w", err)
		}
		defer func() {
			_ = tx.Rollback() // ロールバックエラーは無視
		}()
	}

	for i, record := range records[1:] {
		rowNum := i + 2

		if len(record) < len(expectedFields) {
			result.Failed++
			result.Errors = append(result.Errors, ImportError{
				Row:     rowNum,
				Field:   "record",
				Message: fmt.Sprintf("not enough fields: expected %d, got %d", len(expectedFields), len(record)),
			})
			continue
		}

		photographRow := PhotographImportRow{
			Name:      strings.TrimSpace(record[0]),
			GroupID:   strings.TrimSpace(record[1]),
			PhotoType: strings.TrimSpace(record[2]),
		}

		if validateOnly {
			previewData := map[string]interface{}{
				"name":       photographRow.Name,
				"group_id":   photographRow.GroupID,
				"photo_type": photographRow.PhotoType,
			}
			result.PreviewData = append(result.PreviewData, previewData)
			result.Success++
			continue
		}

		if err := s.validateAndInsertPhotograph(ctx, tx, photographRow, rowNum, result); err != nil {
			continue
		}

		result.Success++
	}

	if !validateOnly && result.Failed == 0 {
		if err := tx.Commit(); err != nil {
			return nil, fmt.Errorf("failed to commit transaction: %w", err)
		}
	}

	return result, nil
}

func (s *ImportService) ImportSceneFromCSV(ctx context.Context, reader io.Reader, validateOnly bool) (*ImportResult, error) {
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	if len(records) == 0 {
		return &ImportResult{Success: 0, Failed: 0, Errors: []ImportError{}}, nil
	}

	header := records[0]
	expectedFields := []string{"photograph_id", "member_id", "color_type_id", "vocal_max", "dance_max", "performance_max", "center_skill", "expected_value", "ssr_plus"}
	if err := validateHeader(header, expectedFields); err != nil {
		return nil, err
	}

	result := &ImportResult{
		Success: 0,
		Failed:  0,
		Errors:  []ImportError{},
	}

	if validateOnly {
		result.PreviewData = make([]map[string]interface{}, 0)
	}

	var tx *sql.Tx
	if !validateOnly {
		var err error
		tx, err = s.DB.BeginTx(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to begin transaction: %w", err)
		}
		defer func() {
			_ = tx.Rollback() // ロールバックエラーは無視
		}()
	}

	for i, record := range records[1:] {
		rowNum := i + 2

		if len(record) < len(expectedFields) {
			result.Failed++
			result.Errors = append(result.Errors, ImportError{
				Row:     rowNum,
				Field:   "record",
				Message: fmt.Sprintf("not enough fields: expected %d, got %d", len(expectedFields), len(record)),
			})
			continue
		}

		sceneRow := SceneImportRow{
			PhotographID:   strings.TrimSpace(record[0]),
			MemberID:       strings.TrimSpace(record[1]),
			ColorTypeID:    strings.TrimSpace(record[2]),
			VocalMax:       strings.TrimSpace(record[3]),
			DanceMax:       strings.TrimSpace(record[4]),
			PerformanceMax: strings.TrimSpace(record[5]),
			CenterSkill:    strings.TrimSpace(record[6]),
			ExpectedValue:  strings.TrimSpace(record[7]),
			SSRPlus:        strings.TrimSpace(record[8]),
		}

		if validateOnly {
			previewData := map[string]interface{}{
				"photograph_id":   sceneRow.PhotographID,
				"member_id":       sceneRow.MemberID,
				"color_type_id":   sceneRow.ColorTypeID,
				"vocal_max":       sceneRow.VocalMax,
				"dance_max":       sceneRow.DanceMax,
				"performance_max": sceneRow.PerformanceMax,
				"center_skill":    sceneRow.CenterSkill,
				"expected_value":  sceneRow.ExpectedValue,
				"ssr_plus":        sceneRow.SSRPlus,
			}
			result.PreviewData = append(result.PreviewData, previewData)
			result.Success++
			continue
		}

		if err := s.validateAndInsertScene(ctx, tx, sceneRow, rowNum, result); err != nil {
			continue
		}

		result.Success++
	}

	if !validateOnly && result.Failed == 0 {
		if err := tx.Commit(); err != nil {
			return nil, fmt.Errorf("failed to commit transaction: %w", err)
		}
	}

	return result, nil
}

func validateHeader(header, expected []string) error {
	if len(header) < len(expected) {
		return fmt.Errorf("insufficient header fields: expected %v, got %v", expected, header)
	}

	for i, expectedField := range expected {
		if strings.TrimSpace(header[i]) != expectedField {
			return fmt.Errorf("invalid header at position %d: expected '%s', got '%s'", i, expectedField, strings.TrimSpace(header[i]))
		}
	}

	return nil
}

func (s *ImportService) validateAndInsertMusic(ctx context.Context, tx *sql.Tx, row MusicImportRow, rowNum int, result *ImportResult) error {
	if row.Name == "" {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "name",
			Message: "name is required",
		})
		return fmt.Errorf("validation failed")
	}

	normal, err := strconv.ParseInt(row.Normal, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "normal",
			Message: "invalid normal score",
		})
		return fmt.Errorf("validation failed")
	}

	pro, err := strconv.ParseInt(row.Pro, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "pro",
			Message: "invalid pro score",
		})
		return fmt.Errorf("validation failed")
	}

	master, err := strconv.ParseInt(row.Master, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "master",
			Message: "invalid master score",
		})
		return fmt.Errorf("validation failed")
	}

	length, err := strconv.ParseInt(row.Length, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "length",
			Message: "invalid length",
		})
		return fmt.Errorf("validation failed")
	}

	colorTypeID, err := strconv.ParseInt(row.ColorTypeID, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "color_type_id",
			Message: "invalid color_type_id",
		})
		return fmt.Errorf("validation failed")
	}

	liveID, err := strconv.ParseInt(row.LiveID, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "live_id",
			Message: "invalid live_id",
		})
		return fmt.Errorf("validation failed")
	}


	if err := s.Querier.RegistMusic(ctx, tx, repository.RegistMusicParams{
		Name:        row.Name,
		Normal:      normal,
		Pro:         pro,
		Master:      master,
		Length:      length,
		ColorTypeID: colorTypeID,
		LiveID:      liveID,
	}); err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "database",
			Message: fmt.Sprintf("failed to insert: %v", err),
		})
		return fmt.Errorf("database error")
	}

	return nil
}

func (s *ImportService) validateAndInsertPhotograph(ctx context.Context, tx *sql.Tx, row PhotographImportRow, rowNum int, result *ImportResult) error {
	if row.Name == "" {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "name",
			Message: "name is required",
		})
		return fmt.Errorf("validation failed")
	}

	groupID, err := strconv.ParseInt(row.GroupID, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "group_id",
			Message: "invalid group_id",
		})
		return fmt.Errorf("validation failed")
	}

	if err := s.Querier.RegistPhotograph(ctx, tx, repository.RegistPhotographParams{
		Name:      row.Name,
		GroupID:   groupID,
		PhotoType: row.PhotoType,
	}); err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "database",
			Message: fmt.Sprintf("failed to insert: %v", err),
		})
		return fmt.Errorf("database error")
	}

	return nil
}

func (s *ImportService) validateAndInsertScene(ctx context.Context, tx *sql.Tx, row SceneImportRow, rowNum int, result *ImportResult) error {
	photographID, err := strconv.ParseInt(row.PhotographID, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "photograph_id",
			Message: "invalid photograph_id",
		})
		return fmt.Errorf("validation failed")
	}

	memberID, err := strconv.ParseInt(row.MemberID, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "member_id",
			Message: "invalid member_id",
		})
		return fmt.Errorf("validation failed")
	}

	colorTypeID, err := strconv.ParseInt(row.ColorTypeID, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "color_type_id",
			Message: "invalid color_type_id",
		})
		return fmt.Errorf("validation failed")
	}

	vocalMax, err := strconv.ParseInt(row.VocalMax, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "vocal_max",
			Message: "invalid vocal_max",
		})
		return fmt.Errorf("validation failed")
	}

	danceMax, err := strconv.ParseInt(row.DanceMax, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "dance_max",
			Message: "invalid dance_max",
		})
		return fmt.Errorf("validation failed")
	}

	performanceMax, err := strconv.ParseInt(row.PerformanceMax, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "performance_max",
			Message: "invalid performance_max",
		})
		return fmt.Errorf("validation failed")
	}

	centerSkill, err := strconv.ParseInt(row.CenterSkill, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "center_skill",
			Message: "invalid center_skill",
		})
		return fmt.Errorf("validation failed")
	}

	expectedValue, err := strconv.ParseInt(row.ExpectedValue, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "expected_value",
			Message: "invalid expected_value",
		})
		return fmt.Errorf("validation failed")
	}

	ssrPlus, err := strconv.ParseInt(row.SSRPlus, 10, 64)
	if err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "ssr_plus",
			Message: "invalid ssr_plus",
		})
		return fmt.Errorf("validation failed")
	}

	centerSkillNull := sql.NullString{
		String: strconv.FormatInt(centerSkill, 10),
		Valid:  true,
	}
	expectedValueNull := sql.NullString{
		String: strconv.FormatInt(expectedValue, 10),
		Valid:  true,
	}

	if err := s.Querier.RegistScene(ctx, tx, repository.RegistSceneParams{
		PhotographID:   photographID,
		MemberID:       memberID,
		ColorTypeID:    colorTypeID,
		VocalMax:       vocalMax,
		DanceMax:       danceMax,
		PerformanceMax: performanceMax,
		CenterSkill:    centerSkillNull,
		ExpectedValue:  expectedValueNull,
		SsrPlus:        ssrPlus,
	}); err != nil {
		result.Failed++
		result.Errors = append(result.Errors, ImportError{
			Row:     rowNum,
			Field:   "database",
			Message: fmt.Sprintf("failed to insert: %v", err),
		})
		return fmt.Errorf("database error")
	}

	return nil
}