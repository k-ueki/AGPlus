package input

type (
	/*
		授業に対する評価の、評価項目
		1. 分かりやすさ
		2. 教員の意欲
		3. 出欠確認の有無
		4. テストの難易度
		5. 楽単度
	*/
	ReviewClassRequest struct {
		Understanding   int `query:"understanding" binding:"required"`
		Motivation      int `query:"motivation" binding:"required"`
		Attendance      int `query:"attendance" binding:"required"`
		TestsDifficulty int `query:"testDifficulty" binding:"required"`
		Easiness        int `query:"easiness" binding:"required"`
	}
)
