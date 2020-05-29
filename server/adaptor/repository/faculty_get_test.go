package repository

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	mock_repository "github.com/k-ueki/AGPlus/server/domain/mock"

	"github.com/k-ueki/AGPlus/server/domain/model"

	"github.com/k-ueki/AGPlus/server/domain/entity"
)

func TestFacultyGetRepository_FindFacultiesByCampusID(t *testing.T) {
	type args struct {
		campusID int
	}
	tests := []struct {
		name    string
		args    args
		mock    func(*mock_repository.MockFacultyGetRepository) *mock_repository.MockFacultyGetRepository
		want    []*entity.Faculty
		wantErr bool
	}{
		{
			name: "[成功]AoyamaキャンパスのFaculty取得",
			args: args{
				campusID: 1,
			},
			mock: func(rep *mock_repository.MockFacultyGetRepository) *mock_repository.MockFacultyGetRepository {
				rep.EXPECT().FindFacultiesByCampusID(1).Return([]*entity.Faculty{
					{
						ID:     1,
						Name:   "文学部",
						Type:   model.FacultyTypeFaculty,
						Campus: model.CampusAoyama,
					},
					{
						ID:     2,
						Name:   "法学部",
						Type:   model.FacultyTypeFaculty,
						Campus: model.CampusAoyama,
					},
				}, nil)
				return rep
			},
			want: []*entity.Faculty{
				{
					ID:     1,
					Name:   "文学部",
					Type:   model.FacultyTypeFaculty,
					Campus: model.CampusAoyama,
				},
				{
					ID:     2,
					Name:   "法学部",
					Type:   model.FacultyTypeFaculty,
					Campus: model.CampusAoyama,
				},
			},
			wantErr: false,
		},
		{
			name: "[成功]SagamiharaキャンパスのFaculty取得",
			args: args{
				campusID: 2,
			},
			mock: func(rep *mock_repository.MockFacultyGetRepository) *mock_repository.MockFacultyGetRepository {
				rep.EXPECT().FindFacultiesByCampusID(2).Return([]*entity.Faculty{
					{
						ID:     3,
						Name:   "理工学部",
						Type:   model.FacultyTypeFaculty,
						Campus: model.CampusSagamihara,
					},
					{
						ID:     4,
						Name:   "社会情報学部",
						Type:   model.FacultyTypeFaculty,
						Campus: model.CampusSagamihara,
					},
				}, nil)
				return rep
			},
			want: []*entity.Faculty{
				{
					ID:     3,
					Name:   "理工学部",
					Type:   model.FacultyTypeFaculty,
					Campus: model.CampusSagamihara,
				},
				{
					ID:     4,
					Name:   "社会情報学部",
					Type:   model.FacultyTypeFaculty,
					Campus: model.CampusSagamihara,
				},
			},
			wantErr: false,
		},
		{
			name: "[失敗]AoyamaキャンパスのFaculty取得",
			args: args{
				campusID: 3,
			},
			mock: func(rep *mock_repository.MockFacultyGetRepository) *mock_repository.MockFacultyGetRepository {
				rep.EXPECT().FindFacultiesByCampusID(3).Return(nil, fmt.Errorf("bad request"))
				return rep
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			rep := mock_repository.NewMockFacultyGetRepository(ctrl)
			rep = tt.mock(rep)

			got, err := rep.FindFacultiesByCampusID(tt.args.campusID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindFacultiesByCampusID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindFacultiesByCampusID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
