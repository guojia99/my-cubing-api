package gaoxiao

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const filePath = "/etc/mycube/gaoxiao.txt"

func removeRepeatedElement[S ~[]E, E comparable](s S) S {
	result := make([]E, 0)
	m := make(map[E]bool)
	for _, v := range s {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}

func getWcaIDs() []string {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return []string{}
	}
	data := string(f)

	line := strings.Split(data, "\n")
	for i := range line {
		newLine := strings.Split(line[i], ".")[1]
		line[i] = strings.ReplaceAll(newLine, " ", "")
	}
	line = removeRepeatedElement(line)
	return line
}

type EventData struct {
	Event          string `json:"event"`
	CountryRank    string `json:"country_rank"`
	ContinentRank  string `json:"continent_rank"`
	WorldRank      string `json:"world_rank"`
	Single         string `json:"single"`
	Average        string `json:"average"`
	WorldRank2     string `json:"world_rank_2"`
	ContinentRank2 string `json:"continent_rank_2"`
	CountryRank2   string `json:"country_rank_2"`
}

type PersonResult struct {
	WCA    string
	Name   string
	Events []EventData
}

var eventsList = []string{
	"3x3x3Cube",
	"2x2x2Cube",
	"4x4x4Cube",
	"5x5x5Cube",
	"6x6x6Cube",
	"7x7x7Cube",
	"3x3x3Blindfolded",
	"3x3x3FewestMoves",
	"3x3x3One-Handed",
	"Clock",
	"Megaminx",
	"Pyraminx",
	"Skewb",
	"Square-1",
	"4x4x4Blindfolded",
	"5x5x5Blindfolded",
	//3x3x3Multi-Blind
}

const urlFormat = "https://www.worldcubeassociation.org/persons/%s" // 2017XUYO01

func getWCAResults(wcaID string) (*PersonResult, error) {
	// 请求网页
	resp, err := http.Get(fmt.Sprintf(urlFormat, wcaID))
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	// 解析网页
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	var events []EventData
	// 查找并提取所需的数据 (例如：提取所有链接的文本和地址)
	doc.Find(".personal-records").Each(
		func(_ int, item *goquery.Selection) {
			item.Find(".table-responsive").Each(
				func(_ int, table *goquery.Selection) {
					table.Find("table.table-striped tbody tr").Each(
						func(index int, row *goquery.Selection) {
							// 提取每一列的数据
							event := row.Find("td.event").Text()
							countryRank := row.Find("td.country-rank").First().Text()
							continentRank := row.Find("td.continent-rank").First().Text()
							worldRank := row.Find("td.world-rank").First().Text()
							single := row.Find("td.single").Text()
							average := row.Find("td.average").Text()
							worldRank2 := row.Find("td.world-rank").Last().Text()
							continentRank2 := row.Find("td.continent-rank").Last().Text()
							countryRank2 := row.Find("td.country-rank").Last().Text()

							// 将数据添加到 events 列表中
							events = append(
								events, EventData{
									Event:          strings.ReplaceAll(strings.ReplaceAll(event, "\n", ""), " ", ""),
									CountryRank:    strings.ReplaceAll(strings.ReplaceAll(countryRank, "\n", ""), " ", ""),
									ContinentRank:  strings.ReplaceAll(strings.ReplaceAll(continentRank, "\n", ""), " ", ""),
									WorldRank:      strings.ReplaceAll(strings.ReplaceAll(worldRank, "\n", ""), " ", ""),
									Single:         strings.ReplaceAll(strings.ReplaceAll(single, "\n", ""), " ", ""),
									Average:        strings.ReplaceAll(strings.ReplaceAll(average, "\n", ""), " ", ""),
									WorldRank2:     strings.ReplaceAll(strings.ReplaceAll(worldRank2, "\n", ""), " ", ""),
									ContinentRank2: strings.ReplaceAll(strings.ReplaceAll(continentRank2, "\n", ""), " ", ""),
									CountryRank2:   strings.ReplaceAll(strings.ReplaceAll(countryRank2, "\n", ""), " ", ""),
								},
							)
						},
					)
				},
			)
		},
	)

	var name string
	doc.Find("#person").Each(
		func(i int, selection *goquery.Selection) {
			selection.Find(".text-center").Each(
				func(i int, selection *goquery.Selection) {
					selection.Find("h2").Each(
						func(i int, selection *goquery.Selection) {
							name = selection.Text()
						},
					)
				},
			)
		},
	)

	name = strings.ReplaceAll(strings.ReplaceAll(name, "\n", ""), " ", "")
	//fmt.Printf("%+v\n", events)
	return &PersonResult{
		WCA:    wcaID,
		Name:   name,
		Events: events,
	}, nil
}

type Persons struct {
	WcaId string `gorm:"column:wca_id"`
	Name  string `gorm:"column:name"`
}

type Results struct {
	EventId    string  `json:"eventId"`
	Best       float64 `json:"best"`
	BestStr    string  `json:"bestStr"`
	Average    float64 `json:"average"`
	AverageStr string  `json:"averageStr"`
	PersonName string  `json:"personName"`
	PersonId   string  `json:"personId"`
}

type PersonBestResults struct {
	PersonName string             `json:"PersonName"`
	Best       map[string]Results `json:"Best"`
	Avg        map[string]Results `json:"Avg"`
}

func getCnName(input string) string {
	re := regexp.MustCompile(`\((.*?)\)`)
	match := re.FindStringSubmatch(input)

	if len(match) > 1 {
		return match[1]
	}
	return input
}

func parserTimeToSeconds(t string) float64 {
	// 解析纯秒数格式
	if regexp.MustCompile(`^\d+(\.\d+)?$`).MatchString(t) {
		seconds, _ := strconv.ParseFloat(t, 64)
		return seconds
	}

	// 解析分+秒格式
	if regexp.MustCompile(`^\d{1,3}:\d{1,3}(\.\d+)?$`).MatchString(t) {
		parts := strings.Split(t, ":")
		minutes, _ := strconv.ParseFloat(parts[0], 64)
		seconds, _ := strconv.ParseFloat(parts[1], 64)
		return minutes*60 + seconds
	}

	// 解析时+分+秒格式
	if regexp.MustCompile(`^\d{1,3}:\d{1,3}:\d{1,3}(\.\d+)?$`).MatchString(t) {
		parts := strings.Split(t, ":")
		hours, _ := strconv.ParseFloat(parts[0], 64)
		minutes, _ := strconv.ParseFloat(parts[1], 64)
		seconds, _ := strconv.ParseFloat(parts[2], 64)
		return hours*3600 + minutes*60 + seconds
	}

	return -1
}

func getAllPersonBestResultsMap(allP []*PersonResult) map[string]PersonBestResults {
	var out = make(map[string]PersonBestResults)

	for _, p := range allP {
		pbr := PersonBestResults{
			PersonName: p.Name,
			Best:       make(map[string]Results),
			Avg:        make(map[string]Results),
		}

		for _, val := range p.Events {
			if !slices.Contains(eventsList, val.Event) {
				continue
			}
			pbr.Best[val.Event] = Results{
				EventId:    val.Event,
				Best:       parserTimeToSeconds(val.Single),
				BestStr:    val.Single,
				PersonName: p.Name,
				PersonId:   p.WCA,
			}
			if val.Average != "" {
				pbr.Avg[val.Event] = Results{
					EventId:    val.Event,
					Average:    parserTimeToSeconds(val.Average),
					AverageStr: val.Average,
					PersonName: p.Name,
					PersonId:   p.WCA,
				}
			}
		}
		out[p.Name] = pbr
	}
	return out
}

func getAllResult() map[string]PersonBestResults {
	var allP []*PersonResult
	for _, wcaId := range getWcaIDs() {
		res, err := getWCAResults(wcaId)
		if err != nil {
			continue
		}
		allP = append(allP, res)
	}
	datas := getAllPersonBestResultsMap(allP)
	return datas
}

type WcaResult struct {
	BestRank        int    `json:"BestRank"`
	BestStr         string `json:"BestStr"`
	BestPersonName  string `json:"BestPersonName"`
	BestPersonWCAID string `json:"BestPersonWCAID"`
	AvgRank         int    `json:"AvgRank"`
	AvgStr          string `json:"AvgStr"`
	AvgPersonName   string `json:"AvgPersonName"`
	AvgPersonWCAID  string `json:"AvgPersonWCAID"`
}

func GetSorAllResults() map[string][]WcaResult {
	var out = make(map[string][]WcaResult)
	data := getAllResult()

	for _, eid := range eventsList {
		var bests []Results
		var avgs []Results

		for _, r := range data {
			if b, ok := r.Best[eid]; ok {
				bests = append(bests, b)
			}
			if a, ok := r.Avg[eid]; ok {
				avgs = append(avgs, a)
			}
		}
		sort.Slice(bests, func(i, j int) bool { return bests[i].Best < bests[j].Best })
		sort.Slice(avgs, func(i, j int) bool { return avgs[i].Average < avgs[j].Average })

		var wrs []WcaResult
		for idx, b := range bests {
			var index = idx + 1
			if idx >= 1 && wrs[idx-1].BestStr == b.BestStr {
				index = wrs[idx-1].BestRank
			}
			wrs = append(
				wrs, WcaResult{
					BestRank:        index,
					BestStr:         b.BestStr,
					BestPersonName:  b.PersonName,
					BestPersonWCAID: b.PersonId,
				},
			)
		}

		for idx, a := range avgs {
			var index = idx + 1

			if idx >= 1 && wrs[idx-1].AvgStr == a.AverageStr {
				index = wrs[idx-1].AvgRank
			}
			wrs[idx].AvgRank = index
			wrs[idx].AvgStr = a.AverageStr
			wrs[idx].AvgPersonName = a.PersonName
			wrs[idx].AvgPersonWCAID = a.PersonId
		}
		out[eid] = wrs
	}
	return out
}
