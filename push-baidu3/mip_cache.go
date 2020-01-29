package main

import (
	"bufio"
	"bytes"
	"encoding/base32"
	"flag"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

type mip_data struct {
	key map[string]string
	urls sync.Map
}

func new_mip_data() *mip_data {
	return &mip_data{
		key: make(map[string]string),
		urls: sync.Map{},
	}
}


// 获取key
func (m mip_data) get_key(r_url string, cookie string) string {

	client := &http.Client{}

	get_ip, found := c.Get("ip")
	fmt.Println("获取ip：", get_ip)
	if found {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(get_ip.(string))
		}
		transport := &http.Transport{Proxy: proxy}
		client = &http.Client{Transport: transport}
	} else {
		cache_ip()
	}




	req, err := http.NewRequest("POST", "https://ziyuan.baidu.com/mip/GetAuthkey", strings.NewReader("site="+ r_url +"/"))
	if err != nil {
		c.Delete("ip")
		fmt.Println("百度禁用ip：", r_url)
		return ""
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	//req.Header.Set("Cookie", "PSTM=1572583620; BIDUPSID=D15592A0EC34A756B3171129C6C33B68; BDUSS=QwZk9FZXp3NGtaZ1R3UjItMFhLWktHOFF2M3JySG9zcDV0RzdLbmlrVEoxd1ZlRUFBQUFBJCQAAAAAAQAAAAEAAACFeeECMTMyMjc5NDMyMTRhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMlK3l3JSt5da3; Hm_lvt_f799cabcf564e60ed4559f6458346ee9=1574844012,1575367492; BAIDUID=4AD937209B4EA76EE70082B948E6496E:SL=0:NR=10:FG=1; H_WISE_SIDS=141183_100807_135171_139405_139388_114745_139811_135846_141000_139148_120196_138470_140853_141034_133995_138878_137985_140173_131247_132552_137743_118884_118859_118853_118829_118796_138165_138883_140260_141030_140631_139047_139297_138585_138778_139177_139625_140077_140113_136196_131862_140591_139693_133847_140793_134256_131423_140311_138662_136537_110085_139539_127969_141110_140593_137911_139886_137252_139408_128201_138312_138425_141193_138944_140684_141190_140597_140962; cflag=13%3A3; H_PS_PSSID=1436_21125_26350_30481; delPer=0; PSINO=7; __cas__st__=NLI; __cas__id__=0; Hm_lvt_6f6d5bc386878a651cb8c9e1b4a3379a=1578898661,1579164511; BDRCVFR[feWj1Vr5u3D]=I67x6TjHwwYf0; SITEMAPSESSID=anplif5fn5eujaa1fu98osk8c5; lastIdentity=PassUserIdentity; Hm_lpvt_6f6d5bc386878a651cb8c9e1b4a3379a=1579483430; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; SITEMAPZHITONG={\"data\":\"6df2f67bd87ef23b4e6577e344abddf11232b66cb604a965ebfd88b524cd2a9fd2da27b42ee1343a2dcc9d532a2a6199cb97a7b13446abedf063944c5d3c7dd34f9b4ae8580c8745aa404c8bf63add3184b8f2b095d20fb34cd60be2e24f865e61424c338aa5abfb688b655e2054c5701b082f63c773b45de1e9ac56f08f34380885d21e63b63c571af9ac003126a2197f909e08297cebd8c2557ba49a59411a43fa628b88cbbd8956c46dfcb99d136d5743077fca514fae5410b27c8d76ed1c4163c2086b501d96fb33f77e5bb1461b7453ce7bc35704af2cfaaf93e636ba9bc586329e7d4b6b084c551c2f2074c98214f95f329d9a42255948429ee122a5b0d359d53b643e5b036e7b1e27045e362efc5da47073862e0347ecdbb534e5231a5d8667d46350390fe40996ec1b6befcf883116f733568273a09ddb2e7f42479d401b3bad96177dfea0544e3a0324562505c83393c3369b1e8151f7fc5a2b5ac2\",\"key_id\":\"32\",\"sign\":\"a67c2b42\"}; SITEMAPZHITONGEXPIRE=1")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Host", "ziyuan.baidu.com")
	req.Header.Set("Origin", "https://ziyuan.baidu.com")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://ziyuan.baidu.com/mip/index?site="+ url.QueryEscape(r_url) +"/")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	req.Header.Set("X-Request-By", "baidu.ajax")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	resp, err := client.Do(req)
	if err != nil {
		c.Delete("ip")
		fmt.Println("百度key失败1：", r_url)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Delete("ip")
		fmt.Println("百度key失败2：", r_url)
		return ""
	}
	fmt.Println(string(body))
	value := gjson.Get(string(body), "data")
	fmt.Println(r_url, value.String())
	start.urls.Store(r_url, 1)
	return value.String()
	//m.key[url] = value.String()
}

// 发送请求
func mip_http_push_baidu(r_url string, key string) {
	cacle_url := fmt.Sprintf(`%s`, "http://c.mipcdn.com/update-ping/c/")
	//fmt.Println("http://" + r_url + "/"+ getToken(10) + "/" + getToken(10) + suffix(rand.Intn(10)))
	//code_url := url.QueryEscape("http://" + r_url + "/"+ getToken(10) + "/" + getToken(10) + suffix(rand.Intn(10)))

	fmt.Println(r_url + "/"+ time.Now().Format("20060102") + "/" + getToken(10) + suffix(rand.Intn(10)))
	code_url := url.QueryEscape(r_url + "/"+ time.Now().Format("20060102") + "/" + getToken(10) + suffix(rand.Intn(10)))


	push_url := cacle_url + code_url
	var postData = []byte("key=" + key)

	req, err := http.NewRequest("POST", push_url, bytes.NewBuffer(postData))
	req.Header.Set("Content-Type", "text/plain")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		num,_ :=start.urls.Load(r_url)
		start.urls.Store(r_url, num.(int) + 1)
		fmt.Println(push_url, "key=" + key, "response Status:", resp.Status, "response Body:", string(body))
	}
}

// 获取随机数
func suffix(s int) string {
	switch {
	case s == 1:
		return ".html"
	case s == 2:
		return ".xml"
	case s == 3:
		return ".doc"
	case s == 4:
		return ".ai"
	case s == 5:
		return ".rtf"
	case s == 6:
		return ".xls"
	case s == 7:
		return ".txt"
	case s == 8:
		return ".xml"
	default:
		return ".html"
	}
}

// 生成随机字符串
func getToken(length int) string {
	randomBytes := make([]byte, 60)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	s:= base32.StdEncoding.EncodeToString(randomBytes)[:length]
	return strings.ToLower(s)
}

var start = new_mip_data()

func foreach(url string, key string)  {
	for i := 0; i < 10; i++ {
		mip_http_push_baidu(url, key)
	}

}

func file_read(filename string) string {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return ""
	}
	defer fi.Close()
	br := bufio.NewReader(fi)

	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		return string(a)
	}
	return ""
}


func get_daili_ip() string {
	//resp, err := http.Get("http://http.tiqu.alicdns.com/getip3?num=1&type=2&pro=&city=0&yys=0&port=1&time=1&ts=1&ys=0&cs=0&lb=1&sb=0&pb=4&mr=1&regions=&gm=4")
	resp, err := http.Get("http://d.jghttp.golangapi.com/getip?num=1&type=2&pro=&city=0&yys=0&port=1&pack=17242&ts=0&ys=0&cs=0&lb=1&sb=0&pb=4&mr=1&regions=")
	if err != nil {
		fmt.Println("链接代理ip出问题：", err)
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("获取代理ip出问题：", err)
		return ""
	}

	fmt.Println(string(body))

	return string(body)
}

var c = cache.New(300*time.Second, 100*time.Second)

func cache_ip() bool {
	ip:=get_daili_ip()
	if ip == ""{
		return false
	}
	code := gjson.Get(ip, "code")
	if (fmt.Sprintf("%s", code) != "0"){
		return false
	}

	ips := gjson.Get(ip, "data.#.ip").Array()[0]
	port := gjson.Get(ip, "data.#.port").Array()[0]
	get_ip := fmt.Sprintf("http://%s:%s", ips, port)

	//c := cache.New(300*time.Second, 100*time.Second)
	c.Set("ip", get_ip, cache.DefaultExpiration)

	return true
}

var filename = flag.String("filename", "mip_cache.log", "配置文件")
var cookie = flag.String("cookie", "cookie.log", "cookie文件")

func main() {

	//ip:=get_daili_ip()
	//
	//ips := gjson.Get(ip, "data.#.ip").Array()[0]
	//
	//port := gjson.Get(ip, "data.#.port").Array()[0]
	////
	//expire_time := gjson.Get(ip, "data.#.ip").Array()[0]
	////
	//
	//get_ipss := fmt.Sprintf("http://%s:%s", ips, port)
	//
	//
	//fmt.Println(get_ipss)
	//
	//os.Exit(0)


	//rand.Seed(int64(time.Now().UnixNano()))
	//index := rand.Intn(5)

	cache_ip()

	flag.Parse() //flag解析
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	//fmt.Println(*cookie)
	//os.Exit(0)
	// 去读文件内容
	fi, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	br := bufio.NewReader(fi)

	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		start.urls.Store(string(a), 0)
		// 将获取的url在获取key，并且保存在map的key中
		//start.get_key(string(a))
	}

	cookie := file_read(*cookie)

	//fmt.Println((start.key))

	for {
		start.urls.Range(func(k, v interface{}) bool {
			//fmt.Println("iterate:", k, v)
			value,_ := start.urls.Load(k)
			if (value.(int) == 0 || value.(int) == 11) {
				get_key := start.get_key(k.(string), cookie)
				if get_key != "" {
					go foreach(k.(string), get_key)
				}
			}
			return true
		})

	}
}
