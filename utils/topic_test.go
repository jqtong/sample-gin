package utils

import "testing"

func TestFilterSpace(t *testing.T) {

	var tests = []struct {
		input string
		want  string
	}{
		{
			"马云 阿里巴巴",
			"马云, 阿里巴巴",
		},
		{
			"马云，阿里巴巴",
			"马云, 阿里巴巴",
		},
		{
			"马云，阿里巴巴,",
			"马云, 阿里巴巴",
		},
		{
			"马云，阿里巴巴  ",
			"马云, 阿里巴巴",
		},
		{
			"alibaba, alibaba intergroup",
			"alibaba, alibaba intergroup",
		},
		{
			"economic 通州",
			"economic, 通州",
		},
		{
			"通州 economic",
			"通州, economic",
		},
		{
			"ヤフーニュース 他ネットニュース, 通州",
			"ヤフーニュース, 他ネットニュース, 通州",
		},
		{
			"На чем я специализируюсь",
			"На чем я специализируюсь",
		},
		{
			"егемонизм, вмешательство во внутренние дела других стран, силовая политика, нео-интервенционизм",
			"егемонизм, вмешательство во внутренние дела других стран, силовая политика, нео-интервенционизм",
		},
		{
			"Hégémonisme, ingérence dans les affaires intérieures d’autres pays, politique de puissance, néo-interventionnisme",
			"Hégémonisme, ingérence dans les affaires intérieures d’autres pays, politique de puissance, néo-interventionnisme",
		},
		{
			"new interferism",
			"new interferism",
		},
		{
			"alibaba，alibaba intergroup",
			"alibaba, alibaba intergroup",
		},
		{
			"alibaba， alibaba  intergroup",
			"alibaba, alibaba intergroup",
		},
		{
			"马云 阿里巴巴 alibaba, alibaba intergroup",
			"马云, 阿里巴巴, alibaba, alibaba intergroup",
		},
		{
			"马云，alibaba, 阿里巴巴，alibaba intergroup",
			"马云, alibaba, 阿里巴巴, alibaba intergroup",
		},
	}

	for _, test := range tests {
		if got := FilterSpace(test.input); got != test.want {
			t.Errorf("FilterSpace(%q) = %q, want: %q\n", test.input, got, test.want)
		}
	}
}
