package main

// The program is used to fetch news from HFUT_XC website
// and store into the database.
import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/dremygit/xwindy-lite/models"
)

const (
	BASE_URL  = "http://xc.hfut.edu.cn"
	LIST_NEWS = "/120/list%d.htm"
	LIST_INFO = "/121/list%d.htm"
	LIST_REPT = "/xsdt/list%d.htm"
)

func main() {

	tatalPagePtf := flag.Int("page", 1, "Number of page list")
	useArticleTime := flag.Bool("t", false, "Use article date as datetime")
	flag.Parse()

	count := 0
	for _, listUrl := range []string{LIST_NEWS, LIST_INFO, LIST_REPT} {
		for i := *tatalPagePtf; i > 0; i-- {
			urlList := getUrlLisyByType(listUrl, i)
			for index, _ := range urlList {
				article, err := getArticleFromUrl(urlList[len(urlList)-1-index])
				if err != nil {
					fmt.Println(err)
					continue
				} else {
					existed, err := article.IsNewsExisted()
					if err != nil {
						fmt.Println(err)
						continue
					}
					if existed {
						continue
					}

					if *useArticleTime {
						article.Time = article.Date.Add(4 * time.Hour)
					}

					if err := article.Create(); err != nil {
						fmt.Println(err)
						continue
					}

					count++
					fmt.Printf("New: title: %s\tdate: %s\turl:%s\n", article.Title, article.Date, article.SourceURL)
				}
			}
		}
	}
	fmt.Printf("Count: %d\n", count)
}

func fetch(url string) (string, error) {
	res, err := http.Get(BASE_URL + url)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	htmlByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	html := string(htmlByte)
	return html, nil
}

func getUrlLisyByType(newsType string, page int) (urlList []string) {

	listUrl := fmt.Sprintf(newsType, page)
	html, err := fetch(listUrl)
	if err != nil {
		log.Fatal(err)
		return []string{}
	}

	re := regexp.MustCompile(`<a class=" articlelist1_a_title ".*?href="(/[^"]*?)"`)
	result := re.FindAllStringSubmatch(html, -1)
	for _, value := range result {
		urlList = append(urlList, value[1])
	}
	return urlList
}

func getArticleFromUrl(newsUrl string) (*models.News, error) {
	var article models.News
	html, err := fetch(newsUrl)
	if err != nil {
		return nil, err
	}
	resTitle := regexp.MustCompile(`<h1 class="atitle">([^<]*?)\s*</h1>`).FindStringSubmatch(html)
	resDate := regexp.MustCompile(`<span class="posttime">发布时间:([^<]*?)</span>`).FindStringSubmatch(html)
	resContent := regexp.MustCompile(`<div class="entry" id="infobox">\s*((?:.*?\s*?)*?)</div>`).FindStringSubmatch(html)

	if len(resTitle) == 0 {
		fmt.Println(resTitle)
		return nil, errors.New(newsUrl + ": Regex title error")
	}
	if len(resDate) == 0 {
		return nil, errors.New(newsUrl + ": Regex date error")
	}
	if len(resContent) == 0 {
		return nil, errors.New(newsUrl + ": Regex content error")
	}

	article.Title = cleanBlank(resTitle[1])
	article.Date, err = time.Parse("2006-01-02", resDate[1])
	if err != nil {
		panic(err)
	}
	article.Summary = cleanHtmlTags(resContent[1])
	summaryRune := []rune(article.Summary)
	if len(summaryRune) > 200 {
		article.Summary = string(summaryRune[:200])
	}
	article.Content = cleanHtmlStyle(resContent[1])
	article.SourceURL = BASE_URL + newsUrl
	return &article, nil
}

func cleanBlank(str string) string {
	return regexp.MustCompile(`\s+`).ReplaceAllString(str, "")
}

func cleanHtmlStyle(html string) (res string) {
	res = html
	res = regexp.MustCompile(`<(\w+?)([^>]*?text-align:right;)`).ReplaceAllString(res, "<$1 TAR $2")
	res = regexp.MustCompile(` (?:style|id|class)="[^"]*"`).ReplaceAllString(res, "")
	res = regexp.MustCompile(`TAR`).ReplaceAllString(res, `style="text-align:right;"`)
	res = regexp.MustCompile(`(href|src)="(/[^"]*?)"`).ReplaceAllString(res, `$1="`+BASE_URL+`$2"`)
	return res
}
func cleanHtmlTags(html string) (res string) {
	res = regexp.MustCompile(`<[^>]*?>`).ReplaceAllString(html, "")
	return res
}

type Connect struct {
	db *sql.DB
}

func (c *Connect) close() {
	c.db.Close()
}
