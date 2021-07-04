package validate

import (
	"fmt"
	"strings"
)

func getChMessage(data map[string][]string,) map[string][]string {
	messages := make(map[string][]string)
	fmt.Println("data>>>",data)
	for field, strSlice := range data {
		var item []string
		for _, rule := range strSlice {
			switch rule {
			case "alpha":
				item = append(item, fmt.Sprintf("%s:%s 只能是字母", field))
			case "alpha_dash":
				item = append(item, fmt.Sprintf("%s:%s 只能包含字母数字字符、破折号和下划线。", field))
			case "alpha_space":
				item = append(item, fmt.Sprintf("%s:%s 只能包含字母数字字符、破折号、下划线和空格。", field))
			case "alpha_num":
				item = append(item, fmt.Sprintf("%s:%s 必须包含字母和数字。", field))
			case "numeric":
				item = append(item, fmt.Sprintf("%s:%s 必须为数字。", field))
			case "bool":
				item = append(item, fmt.Sprintf("%s:%s 必须能够转换为布尔值。", field))
			case "coordinate":
				item = append(item, fmt.Sprintf("%s:%s 必须是有效坐标的值。", field))
			case "css_color":
				item = append(item, fmt.Sprintf("%s:%s 必须是一个有效的CSS颜色值。", field))
			case "date":
				item = append(item, fmt.Sprintf("%s:%s 必须是有效日期。", field))
			case "email":
				item = append(item, fmt.Sprintf("%s:%s 必须是有效的电子邮件。", rule,field))
			case "float":
				item = append(item, fmt.Sprintf("%s:%s 必须是一个有效的浮点数。", field))
			case "mac_address":
				item = append(item, fmt.Sprintf("%s:%s 必须是一个有效的Mac地址。", field))
			case "ip":
				item = append(item, fmt.Sprintf("%s:%s 必须是一个有效的ip地址。", field))
			case "ip_v4":
				item = append(item, fmt.Sprintf("%s:%s 必须是一个有效的IP V4地址。", field))
			case "ip_v6":
				item = append(item, fmt.Sprintf("%s:%s 必须是有效的IP V6地址。", field))
			case "json":
				item = append(item, fmt.Sprintf("%s:%s 必须是一个有效的json字符串。", field))
			case "lat":
				item = append(item, fmt.Sprintf("%s:%s 必须是有效的纬度。", field))
			case "lon":
				item = append(item, fmt.Sprintf("%s:%s 必须是有效的经度。", field))
			case "required":
				item = append(item, fmt.Sprintf("%s:%s 不能为空",rule,field))
			case "url":
				item = append(item, fmt.Sprintf("%s:%s 必须是有效的url。", field))
			case "uuid":
				item = append(item, fmt.Sprintf("%s:%s 必须是有效的uuid。", field))
			case "uuid_v3":
				item = append(item, fmt.Sprintf("%s:%s 必须是有效的UUID V3。", field))
			case "uuid_v4":
				item = append(item, fmt.Sprintf("%s:%s 必须是有效的UUID V4。", field))
			case "uuid_v5":
				item = append(item, fmt.Sprintf("%s:%s 必须是有效的UUID V5。", field))
			case "regex":
				item = append(item, fmt.Sprintf("%s:%s 不是正确的数据格式", field))
			default:
				n := strings.Index(rule, ":")
				if n > 0 {
					rlt := strings.Split(rule, ":")
					switch rlt[0] {
					case "size":
						item = append(item, fmt.Sprintf("%s:%s 文件大小不能超过 %s。", field, rlt[1]))
					case "ext":
						item = append(item, fmt.Sprintf("%s:%s 文件扩展名只能是：%s", field, rlt[1]))
					case "digits:int":
						item = append(item, fmt.Sprintf("%s:%s 必须是数字，并且长度为：%s。", field, rlt[1]))
					case "in":
						item = append(item, fmt.Sprintf("%s:%s 只能是（%s）中的一个。", field, rlt[1]))
					case "not_in":
						item = append(item, fmt.Sprintf("%s:%s 不能能是（%s）中的一个。", field, rlt[1]))
					case "min":
						item = append(item, fmt.Sprintf("%s:%s 的长度不能小于：%s", field, rlt[1]))
					case "max":
						item = append(item, fmt.Sprintf("%s:%s 的长度不能大于：%s", field, rlt[1]))
					case "len":
						item = append(item, fmt.Sprintf("%s:%s 的长度必须是：%s", field, rlt[1]))
					case "between":
						num := strings.Split(rlt[1], ",")
						item = append(item, fmt.Sprintf("%s:%s 长度必须在 %s 和 %s 之间。", field, num[0], num[1]))
					case "numeric_between":
						num := strings.Split(rlt[1], ",")
						item = append(item, fmt.Sprintf("%s:%s 必须是数字，且值范围只能在 %s 和 %s 之间。", field, num[0], num[1]))
					case "digits_between":
						num := strings.Split(rlt[1], ",")
						item = append(item, fmt.Sprintf("%s:%s 必须是数字，且值长度只能在 %s 和 %s 之间。", field, num[0], num[1]))
					}
				}
			}
		}
		messages[field] = item
	}
	return messages
}
