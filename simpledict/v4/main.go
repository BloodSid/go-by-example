package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type CaiyunDictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

type CaiyunDictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

type VolcRequest struct {
	Source         string   `json:"source"`
	Words          []string `json:"words"`
	SourceLanguage string   `json:"source_language"`
	TargetLanguage string   `json:"target_language"`
}

type VolcResponse struct {
	Details []struct {
		Detail string `json:"detail"`
		Extra  string `json:"extra"`
	} `json:"details"`
	BaseResp struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}

type VolcDetail struct {
	ErrorCode string `json:"errorCode"`
	RequestID string `json:"requestId"`
	Msg       string `json:"msg"`
	Result    []struct {
		Ec struct {
			ReturnPhrase []string `json:"returnPhrase"`
			Synonyms     []struct {
				Pos   string   `json:"pos"`
				Words []string `json:"words"`
				Trans string   `json:"trans"`
			} `json:"synonyms"`
			Etymology struct {
				ZhCHS []struct {
					Description string `json:"description"`
					Detail      string `json:"detail"`
					Source      string `json:"source"`
				} `json:"zh-CHS"`
			} `json:"etymology"`
			SentenceSample []struct {
				Sentence     string `json:"sentence"`
				SentenceBold string `json:"sentenceBold"`
				Translation  string `json:"translation"`
			} `json:"sentenceSample"`
			WebDict string `json:"webDict"`
			Web     []struct {
				Phrase   string   `json:"phrase"`
				Meanings []string `json:"meanings"`
			} `json:"web"`
			MTerminalDict string `json:"mTerminalDict"`
			RelWord       struct {
				Word string `json:"word"`
				Stem string `json:"stem"`
				Rels []struct {
					Rel struct {
						Pos   string `json:"pos"`
						Words []struct {
							Word string `json:"word"`
							Tran string `json:"tran"`
						} `json:"words"`
					} `json:"rel"`
				} `json:"rels"`
			} `json:"relWord"`
			Dict  string `json:"dict"`
			Basic struct {
				UsPhonetic string   `json:"usPhonetic"`
				UsSpeech   string   `json:"usSpeech"`
				Phonetic   string   `json:"phonetic"`
				UkSpeech   string   `json:"ukSpeech"`
				ExamType   []string `json:"examType"`
				Explains   []struct {
					Pos   string `json:"pos"`
					Trans string `json:"trans"`
				} `json:"explains"`
				UkPhonetic  string `json:"ukPhonetic"`
				WordFormats []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"wordFormats"`
			} `json:"basic"`
			Phrases []struct {
				Phrase   string   `json:"phrase"`
				Meanings []string `json:"meanings"`
			} `json:"phrases"`
			Lang   string `json:"lang"`
			IsWord bool   `json:"isWord"`
		} `json:"ec"`
	} `json:"result"`
}

func queryCaiyun(word string) {
	client := &http.Client{}
	request := CaiyunDictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("app-name", "xy")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("device-id", "")
	req.Header.Set("os-type", "web")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponse CaiyunDictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("----------彩云----------")
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}

func queryVolc(word string) {
	client := &http.Client{}
	request := VolcRequest{Source: "youdao", Words: []string{word}, SourceLanguage: "en", TargetLanguage: "zh"}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://translate.volcengine.com/web/dict/detail/v1/?msToken=&X-Bogus=DFSzswVLQDcEkLCWSZZYMxewyPWe&_signature=_02B4Z6wo00001T6tJMQAAIDDN5XM6gE7cLE-rSBAACxxNbak0q2E4bm98d1HXINzJ6Ks0-XEgal0ju8Zkp8Ou-tnjOTrivttijXqU6D5KmnL0hb5hEJpC8VbFQsJXwQgHyFizRvrO.-0ny-83a", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "translate.volcengine.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "x-jupiter-uuid=1673864353202977; i18next=zh-CN; s_v_web_id=verify_lcynm9z4_xvZngb9k_7GHf_45uI_9F6C_VsoVpYYgAFMq; ttcid=881dc62803db4171b7f8b3636ac18c6412; tt_scid=8laa0vXJeBnyRDvapswHJ0oXLU.J-UPN439olGWjtdjalOLXJRvSX.8KEgcjAimyc203; csrfToken=8b13c7ddf17a645320fff792b453c01c; isIntranet=-1; ve_doc_history=4640; referrer_title=%E6%96%87%E6%9C%AC%E7%BF%BB%E8%AF%91API%20%E6%9C%BA%E5%99%A8%E7%BF%BB%E8%AF%91-%E7%81%AB%E5%B1%B1%E5%BC%95%E6%93%8E; __tea_cache_tokens_3569={%22web_id%22:%227189196833822180919%22%2C%22user_unique_id%22:%227189196833822180919%22%2C%22timestamp%22:1673865353752%2C%22_type_%22:%22default%22}")
	req.Header.Set("origin", "https://translate.volcengine.com")
	req.Header.Set("referer", "https://translate.volcengine.com/?category=&home_language=zh&source_language=detect&target_language=zh&text=open")
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="99", "Microsoft Edge";v="109", "Chromium";v="109"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.52")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var volcResponse VolcResponse
	err = json.Unmarshal(bodyText, &volcResponse)
	if err != nil {
		log.Fatal(err)
	}
	var detail VolcDetail
	err = json.Unmarshal([]byte(volcResponse.Details[0].Detail), &detail)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("----------火山----------")
	fmt.Printf("%s UK: [%s] US: [%s]\n", word, detail.Result[0].Ec.Basic.UkPhonetic, detail.Result[0].Ec.Basic.UsPhonetic)
	for _, explain := range detail.Result[0].Ec.Basic.Explains {
		fmt.Println(explain.Pos, explain.Trans)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
		`)
		os.Exit(1)
	}
	word := os.Args[1]
	queryCaiyun(word)
	queryVolc(word)
}
