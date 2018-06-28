package parser

import (
	"crawler/engine"
	"regexp"
	"strconv"
	"crawler/model"
)

var ageRe = regexp.MustCompile(`<td><span class="label"> 年龄:</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">([\d]+)CM</span></td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label"> 婚况:</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hukouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	//int
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		profile.Weight = weight
	}
	gender := extractString(contents, genderRe)
	if err != nil {
		profile.Gender = gender
	}

	income := extractString(contents, incomeRe)
	if err != nil {
		profile.Income = income
	}

	marriage := extractString(contents, marriageRe)
	if err != nil {
		profile.Marriage = marriage
	}

	education := extractString(contents, educationRe)
	if err != nil {
		profile.Education = education
	}

	occupation := extractString(contents, occupationRe)
	if err != nil {
		profile.Occupation = occupation
	}

	hukou := extractString(contents, hukouRe)
	if err != nil {
		profile.Hukou = hukou
	}

	house := extractString(contents, houseRe)
	if err != nil {
		profile.Hourse = house
	}

	car := extractString(contents, carRe)
	if err != nil {
		profile.Car = car
	}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
