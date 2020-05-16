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
		EaseOfUnderstanding int `query:"easeOfUnderstanding"`
		Motivation          int `query:"motivation"`
		TestsDifficulty     int `query:"testDifficulty"`
		Easiness            int `query:"easiness"`
	}
)
