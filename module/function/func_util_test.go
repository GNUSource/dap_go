package function

import (
	"testing"
)

func TestDeleteSubStr(t *testing.T) {
	type args struct {
		str       string
		deleteStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: struct {
				str       string
				deleteStr string
			}{
				str:       "0.1g异烟肼片说明书",
				deleteStr: "说明书|标盒|大箱",
			},

			want: "0.1g异烟肼片",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteSubStr(tt.args.str, tt.args.deleteStr); got != tt.want {
				t.Errorf("DeleteSubStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	// 100mg维生素E软胶囊标盒(15粒*2板)
	// 20ml/500ml碘酊说明书
	// 100/500ml乳酸依沙吖啶溶液说明书
	// 0.3g利福平胶囊说明书（西班牙文）
	DrawStr("乙胺利福异烟片标盒（政府）（10片*5板*3袋）",
		`\(\S+\)|\d+(.\d+)?(/)?(mg|g|ml|mm|粒|盒|袋|片|板)(/|\*)?\d+(.\d+)?(mg|g|ml|mm|粒|盒|袋|片|板)|\d+(.\d+)?(mg|g|ml|mm|粒|盒|袋|片|板)`)
}

// deleteSubStr
// 说明书|政府|标盒|中盒|标签|铝箔|大箱|(西班牙文)|(不干胶)|(质量和疗效一致性评价)|(自销)|(膜厚)|(通用)

// DrawStr
// \(\S+\)|(\d+(.\d+)?\S+(\*|/)+(\d+(.\d+)?)?(mg|g|ml|mm|粒|盒|袋|片|板)?)|(\d+(.\d+)?(mg|g|ml|mm|粒|盒|袋|片|板))
