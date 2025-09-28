package service

import (
	"context"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/litencatt/uniar/repository"
)

func TestImportMusicFromCSV(t *testing.T) {
	testCases := []struct {
		name         string
		csvData      string
		validateOnly bool
		expectError  bool
		expectFailed int
		expectSuccess int
	}{
		{
			name:         "ValidMusicCSV",
			csvData:      "name,normal,pro,master,length,color_type_id,live_id\nTest Music,500,800,1200,180,1,1",
			validateOnly: true,
			expectError:  false,
			expectSuccess: 1,
			expectFailed: 0,
		},
		{
			name:         "InvalidMusicCSV_MissingName",
			csvData:      "name,normal,pro,master,length,color_type_id,live_id\n,500,800,1200,180,1,1",
			validateOnly: true,
			expectError:  false,
			expectSuccess: 0,
			expectFailed: 1,
		},
		{
			name:         "InvalidMusicCSV_InvalidNumber",
			csvData:      "name,normal,pro,master,length,color_type_id,live_id\nTest Music,invalid,800,1200,180,1,1",
			validateOnly: true,
			expectError:  false,
			expectSuccess: 0,
			expectFailed: 1,
		},
		{
			name:         "InvalidHeader",
			csvData:      "wrong,header,format\nTest Music,500,800",
			validateOnly: true,
			expectError:  true,
			expectSuccess: 0,
			expectFailed: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			if !tc.validateOnly {
				mock.ExpectBegin()
				if !tc.expectError && tc.expectFailed == 0 {
					q.EXPECT().RegistMusic(ctx, gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					mock.ExpectCommit()
				} else {
					mock.ExpectRollback()
				}
			}

			svc := &ImportService{
				DB:      db,
				Querier: q,
			}

			reader := strings.NewReader(tc.csvData)
			result, err := svc.ImportMusicFromCSV(ctx, reader, tc.validateOnly)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result.Success != tc.expectSuccess {
				t.Errorf("Expected %d successful records, got %d", tc.expectSuccess, result.Success)
			}

			if result.Failed != tc.expectFailed {
				t.Errorf("Expected %d failed records, got %d", tc.expectFailed, result.Failed)
			}

			if tc.validateOnly && tc.expectSuccess > 0 {
				if len(result.PreviewData) != tc.expectSuccess {
					t.Errorf("Expected %d preview records, got %d", tc.expectSuccess, len(result.PreviewData))
				}
			}
		})
	}
}

func TestImportPhotographFromCSV(t *testing.T) {
	testCases := []struct {
		name         string
		csvData      string
		validateOnly bool
		expectError  bool
		expectFailed int
		expectSuccess int
	}{
		{
			name:         "ValidPhotographCSV",
			csvData:      "name,group_id,photo_type\nTest Photo,1,Live",
			validateOnly: true,
			expectError:  false,
			expectSuccess: 1,
			expectFailed: 0,
		},
		{
			name:         "InvalidPhotographCSV_MissingName",
			csvData:      "name,group_id,photo_type\n,1,Live",
			validateOnly: true,
			expectError:  false,
			expectSuccess: 0,
			expectFailed: 1,
		},
		{
			name:         "InvalidPhotographCSV_InvalidGroupID",
			csvData:      "name,group_id,photo_type\nTest Photo,invalid,Live",
			validateOnly: true,
			expectError:  false,
			expectSuccess: 0,
			expectFailed: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			if !tc.validateOnly {
				mock.ExpectBegin()
				if !tc.expectError && tc.expectFailed == 0 {
					q.EXPECT().RegistPhotograph(ctx, gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					mock.ExpectCommit()
				} else {
					mock.ExpectRollback()
				}
			}

			svc := &ImportService{
				DB:      db,
				Querier: q,
			}

			reader := strings.NewReader(tc.csvData)
			result, err := svc.ImportPhotographFromCSV(ctx, reader, tc.validateOnly)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result.Success != tc.expectSuccess {
				t.Errorf("Expected %d successful records, got %d", tc.expectSuccess, result.Success)
			}

			if result.Failed != tc.expectFailed {
				t.Errorf("Expected %d failed records, got %d", tc.expectFailed, result.Failed)
			}
		})
	}
}

func TestImportSceneFromCSV(t *testing.T) {
	testCases := []struct {
		name         string
		csvData      string
		validateOnly bool
		expectError  bool
		expectFailed int
		expectSuccess int
	}{
		{
			name:         "ValidSceneCSV",
			csvData:      "photograph_id,member_id,color_type_id,vocal_max,dance_max,performance_max,center_skill,expected_value,ssr_plus\n1,1,1,5000,4800,4900,100,500,1",
			validateOnly: true,
			expectError:  false,
			expectSuccess: 1,
			expectFailed: 0,
		},
		{
			name:         "InvalidSceneCSV_InvalidPhotographID",
			csvData:      "photograph_id,member_id,color_type_id,vocal_max,dance_max,performance_max,center_skill,expected_value,ssr_plus\ninvalid,1,1,5000,4800,4900,100,500,1",
			validateOnly: true,
			expectError:  false,
			expectSuccess: 0,
			expectFailed: 1,
		},
		{
			name:         "InvalidSceneCSV_InvalidVocalMax",
			csvData:      "photograph_id,member_id,color_type_id,vocal_max,dance_max,performance_max,center_skill,expected_value,ssr_plus\n1,1,1,invalid,4800,4900,100,500,1",
			validateOnly: true,
			expectError:  false,
			expectSuccess: 0,
			expectFailed: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			q := repository.NewMockQuerier(mockCtrl)

			if !tc.validateOnly {
				mock.ExpectBegin()
				if !tc.expectError && tc.expectFailed == 0 {
					q.EXPECT().RegistScene(ctx, gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					mock.ExpectCommit()
				} else {
					mock.ExpectRollback()
				}
			}

			svc := &ImportService{
				DB:      db,
				Querier: q,
			}

			reader := strings.NewReader(tc.csvData)
			result, err := svc.ImportSceneFromCSV(ctx, reader, tc.validateOnly)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result.Success != tc.expectSuccess {
				t.Errorf("Expected %d successful records, got %d", tc.expectSuccess, result.Success)
			}

			if result.Failed != tc.expectFailed {
				t.Errorf("Expected %d failed records, got %d", tc.expectFailed, result.Failed)
			}
		})
	}
}

func TestValidateHeader(t *testing.T) {
	testCases := []struct {
		name        string
		header      []string
		expected    []string
		expectError bool
	}{
		{
			name:        "ValidHeader",
			header:      []string{"name", "normal", "pro", "master"},
			expected:    []string{"name", "normal", "pro", "master"},
			expectError: false,
		},
		{
			name:        "InvalidHeader_WrongOrder",
			header:      []string{"normal", "name", "pro", "master"},
			expected:    []string{"name", "normal", "pro", "master"},
			expectError: true,
		},
		{
			name:        "InvalidHeader_MissingFields",
			header:      []string{"name", "normal"},
			expected:    []string{"name", "normal", "pro", "master"},
			expectError: true,
		},
		{
			name:        "InvalidHeader_WrongFieldName",
			header:      []string{"title", "normal", "pro", "master"},
			expected:    []string{"name", "normal", "pro", "master"},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateHeader(tc.header, tc.expected)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}