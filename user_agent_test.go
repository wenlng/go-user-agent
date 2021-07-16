/**
 * @Author Awen
 * @Description Get "OSName" or "BrowserName" to USER_AGENT for Testing
 * @Date 2021/7/16
 * @Email wengaolng@gmail.com
 **/

package useragent

import (
	"fmt"
	"testing"
)

func getTestUserAgentStr() string {
	var userAgent string
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
	//userAgent = "Mozilla/5.0 {Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
	//userAgent = "Mozilla/5.0 {compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Titan) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
	//userAgent = "Mozilla/5.0 {Macintosh; Intel Mac OS X 10.6.8; U; en) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
	//userAgent = "Mozilla/5.0 {X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
	//userAgent = "Mozilla/5.0 {iPhone; U; CPU iPhone OS 4_3_3 like Mac OS X; en-us) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
	//userAgent = "Mozilla/5.0 {iPhone 6s; CPU iPhone OS 11_4_1 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
	//userAgent = "Mozilla/5.0 {iPad; U; CPU OS 4_3_3 like Mac OS X; en-us) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
	//userAgent = "Mozilla/5.0 {Linux; U; Android 2.3.7; zh-cn; MB200) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
	//userAgent = "Mozilla/5.0 {BlackBerry; U; BlackBerry 9800; en) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
	//userAgent = "Mozilla/5.0 {SymbianOS/9.4; Series60/5.0 NokiaN97-1/20.0.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
	//userAgent = "Mozilla/5.0 {hp-tablet; Linux; hpwOS/3.0.0; U; en-US) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
	//userAgent = "Mozilla/5.0 {X11; CrOS i686 2268.111.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"

	return userAgent
}

func TestGetOsName(t *testing.T) {
	userAgent := getTestUserAgentStr()
	fmt.Println(">>>>>> ", GetOsName(userAgent))
}

func TestGetBrowserName(t *testing.T) {
	userAgent := getTestUserAgentStr()
	fmt.Println(">>>>>> ", GetBrowserName(userAgent))
}