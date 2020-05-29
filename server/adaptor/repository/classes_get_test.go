package repository

import (
	"reflect"
	"testing"

	"github.com/k-ueki/AGPlus/server/domain/entity"

	"github.com/golang/mock/gomock"

	mock_repository "github.com/k-ueki/AGPlus/server/domain/mock"

	"github.com/k-ueki/AGPlus/server/domain/query"
)

func TestClassGetRepositoryImpl_FindAll(t *testing.T) {
	type args struct {
		query *query.ListPaginationQuery
	}
	tests := []struct {
		name    string
		args    args
		mock    func(*mock_repository.MockClassGetRepository) *mock_repository.MockClassGetRepository
		want    []*entity.Class
		wantErr bool
	}{
		{
			name: "[成功]全授業取得",
			args: args{
				&query.ListPaginationQuery{
					Limit:  5,
					Offset: 0,
				},
			},
			mock: func(rep *mock_repository.MockClassGetRepository) *mock_repository.MockClassGetRepository {
				rep.EXPECT().FindAll(&query.ListPaginationQuery{
					Limit:  5,
					Offset: 0,
				}).Return([]*entity.Class{
					{
						ID:          1,
						Name:        "フレッシャーズ・セミナー",
						Semester:    "前期",
						Credits:     2,
						Teacher:     "楠 由記子",
						Description: "「会計」とはどのようなものか、本講義では会計の基礎的な知識を学びます。　本講義では、テキストを用いて輪読を行います。具体的には、毎回、担当箇所のテキスト内容を発表していただきます。そして、その発表内容に関して補足で講義を行ったり、関連する情報を追加説明します。テーマや必要に応じて新聞や雑誌の記事を用いるなど、最新の知識にも触れていきたいと思います。したがって、本講義では、プレゼンテーションやディスカッションを通じて、基本的な会計知識の修得とコミュニケーション能力の向上を目的とします。",
						Year:        2020,
						DayAt:       "月１",
						Campus:      "[青山]",
					},
					{
						ID:          2,
						Name:        "キリスト教概論Ⅰ",
						Semester:    "前期",
						Credits:     2,
						Teacher:     "伊藤 悟",
						Description: "青山学院の建学の精神であるキリスト教とキリスト教信仰について学びます。聖書をもとに、キリスト教の教え、信仰、歴史、生き方、現代の諸問題などを幅広く考察していきます。",
						Year:        2020,
						DayAt:       "月１",
						Campus:      "[青山]",
					},
					{
						ID:          3,
						Name:        "キリスト教概論Ⅰ",
						Semester:    "前期",
						Credits:     2,
						Teacher:     "塩谷 直也",
						Description: "イエスの生涯、聖書に関する基礎知識、およびキリスト教の基本的思想の紹介",
						Year:        2020,
						DayAt:       "月１",
						Campus:      "[青山]",
					},
				}, nil)
				return rep
			},
			want: []*entity.Class{
				{
					ID:          1,
					Name:        "フレッシャーズ・セミナー",
					Semester:    "前期",
					Credits:     2,
					Teacher:     "楠 由記子",
					Description: "「会計」とはどのようなものか、本講義では会計の基礎的な知識を学びます。　本講義では、テキストを用いて輪読を行います。具体的には、毎回、担当箇所のテキスト内容を発表していただきます。そして、その発表内容に関して補足で講義を行ったり、関連する情報を追加説明します。テーマや必要に応じて新聞や雑誌の記事を用いるなど、最新の知識にも触れていきたいと思います。したがって、本講義では、プレゼンテーションやディスカッションを通じて、基本的な会計知識の修得とコミュニケーション能力の向上を目的とします。",
					Year:        2020,
					DayAt:       "月１",
					Campus:      "[青山]",
				},
				{
					ID:          2,
					Name:        "キリスト教概論Ⅰ",
					Semester:    "前期",
					Credits:     2,
					Teacher:     "伊藤 悟",
					Description: "青山学院の建学の精神であるキリスト教とキリスト教信仰について学びます。聖書をもとに、キリスト教の教え、信仰、歴史、生き方、現代の諸問題などを幅広く考察していきます。",
					Year:        2020,
					DayAt:       "月１",
					Campus:      "[青山]",
				},
				{
					ID:          3,
					Name:        "キリスト教概論Ⅰ",
					Semester:    "前期",
					Credits:     2,
					Teacher:     "塩谷 直也",
					Description: "イエスの生涯、聖書に関する基礎知識、およびキリスト教の基本的思想の紹介",
					Year:        2020,
					DayAt:       "月１",
					Campus:      "[青山]",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			r := mock_repository.NewMockClassGetRepository(ctrl)
			r = tt.mock(r)

			got, err := r.FindAll(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassGetRepositoryImpl_FindByID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		mock    func(*mock_repository.MockClassGetRepository) *mock_repository.MockClassGetRepository
		want    *entity.Class
		wantErr bool
	}{
		{
			name: "[成功]授業取得",
			args: args{
				1,
			},
			mock: func(rep *mock_repository.MockClassGetRepository) *mock_repository.MockClassGetRepository {
				rep.EXPECT().FindByID(1).Return(&entity.Class{
					ID:          1,
					Name:        "フレッシャーズ・セミナー",
					Semester:    "前期",
					Credits:     2,
					Teacher:     "楠 由記子",
					Description: "「会計」とはどのようなものか、本講義では会計の基礎的な知識を学びます。　本講義では、テキストを用いて輪読を行います。具体的には、毎回、担当箇所のテキスト内容を発表していただきます。そして、その発表内容に関して補足で講義を行ったり、関連する情報を追加説明します。テーマや必要に応じて新聞や雑誌の記事を用いるなど、最新の知識にも触れていきたいと思います。したがって、本講義では、プレゼンテーションやディスカッションを通じて、基本的な会計知識の修得とコミュニケーション能力の向上を目的とします。",
					Year:        2020,
					DayAt:       "月１",
					Campus:      "[青山]",
				}, nil)
				return rep
			},
			want: &entity.Class{
				ID:          1,
				Name:        "フレッシャーズ・セミナー",
				Semester:    "前期",
				Credits:     2,
				Teacher:     "楠 由記子",
				Description: "「会計」とはどのようなものか、本講義では会計の基礎的な知識を学びます。　本講義では、テキストを用いて輪読を行います。具体的には、毎回、担当箇所のテキスト内容を発表していただきます。そして、その発表内容に関して補足で講義を行ったり、関連する情報を追加説明します。テーマや必要に応じて新聞や雑誌の記事を用いるなど、最新の知識にも触れていきたいと思います。したがって、本講義では、プレゼンテーションやディスカッションを通じて、基本的な会計知識の修得とコミュニケーション能力の向上を目的とします。",
				Year:        2020,
				DayAt:       "月１",
				Campus:      "[青山]",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			r := mock_repository.NewMockClassGetRepository(ctrl)
			r = tt.mock(r)

			got, err := r.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}
