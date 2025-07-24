package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/corpix/uarand"
)

var companyList = []string{
	"https://example-company1.com",
	"https://example-company2.co.jp",
	"https://example-company3.jp",
	"https://www.win-consul.com/",
	"https://www.ulsgroup.co.jp/",
	"https://www.tcs.com/jp",
	"https://www.nri.com/jp/index.html",
	"https://www.mri.co.jp/",
	"https://www.mizuho-rt.co.jp/",
	"https://www.mamezo-dhd.com/",
	"https://www.jri.co.jp/",
	"https://www.drinet.co.jp/",
	"https://sun-asterisk.com/",
	"https://powersolutions.co.jp/",
	"https://magazine.exmotion.co.jp/all",
	"https://itbook.co.jp/",
	"https://hashport.io/",
	"https://gtl-q.c",
	"https://design.jvckenwood.com/",
	"https://atlstech.com/company",
	"https://www.spintechnology.co.jp/",
	"https://www.panasonic.com/jp/company/is.html",
	"https://www.nykbs.co.jp/",
	"https://www.marubeni-itsol.com/",
	"https://www.agrex.co.jp/",
	"https://www.ymsl.co.jp/",
	"https://www.ye-digital.com/jp/",
	"https://www.xnet.co.jp/",
	"https://www.witz-inc.co.jp/",
	"https://www.wipro.com/ja-JP/",
	"https://www.uniadex.co.jp/",
	"https://www.ubicom-hd.com/ja/index.html",
	"https://www.tsuzuki-techno.com/",
	"https://www.tsg.co.jp/",
	"https://www.triumph.jp/",
	"https://www.toyotasystems.com/",
	"https://www.towa-system.co.jp/",
	"https://www.toukei.co.jp/",
	"https://www.toll-tech.co.jp/",
	"https://www.tokai-com.co.jp/",
	"https://www.tjsys.co.jp/",
	"https://www.tiis.co.jp/",
	"https://www.tecsvc.co.jp/",
	"https://www.tecnos.co.jp/",
	"https://www.tecmira.com/company/",
	"https://www.techno-creative.co.jp/",
	"https://www.tdi.co.jp/",
	"https://www.tcs-net.co.jp/",
	"https://www.tcs-ip.net/",
	"https://www.tandi.co.jp/",
	"https://www.takuma-sc.co.jp/",
	"https://www.t-axis.co.jp/",
	"https://www.syshd.co.jp/",
	"https://www.sword.co.jp/",
	"https://www.swc.co.jp/",
	"https://www.suntory-st.co.jp/",
	"https://www.sts-hd.co.jp/",
	"https://www.sra.co.jp/",
	"https://www.sonyglobalsolutions.jp/",
	"https://www.sojitz-tic.com/",
	"https://www.sms-datatech.co.jp/",
	"https://www.slcs.co.jp/",
	"https://www.sjk.co.jp/",
	"https://www.sig-group.co.jp/",
	"https://www.shinwart.co.jp/",
	"https://www.sed.co.jp/",
	"https://www.sec.co.jp/ja/index.html",
	"https://www.sanwa-comp.co.jp/",
	"https://www.ryucom.co.jp/",
	"https://www.ryobi.co.jp/",
	"https://www.pci-h.co.jp/",
	"https://www.pal-net.co.jp/",
	"https://www.opentone.co.jp/",
	"https://www.one-eighty.co.jp/",
	"https://www.ok-toyota.jp/company/system_service",
	"https://www.odk.co.jp/",
	"https://www.nvc.co.jp/",
	"https://www.nttw-bf.com/",
	"https://www.nttdata.com/global/ja/",
	"https://www.nttdata-newson.co.jp/",
	"https://www.nttdata-kansai.co.jp/",
	"https://www.nttdata-ft.co.jp/",
	"https://www.nttdata-ccs.co.jp/",
	"https://www.nttd-mse.com/",
	"https://www.nttd-i.co.jp/",
	"https://www.nttcom.co.jp/",
	"https://www.ntt-me.co.jp/",
	"https://www.ntn.co.jp/",
	"https://www.nti.co.jp/",
	"https://www.nsware.co.jp/",
	"https://www.nssol.nipponsteel.com/kyushu/",
	"https://www.ns-cs.co.jp/",
	"https://www.nova-system.com/",
	"https://www.nos.co.jp/",
	"https://www.nomura-system.co.jp/",
	"https://www.noar.co.jp/",
	"https://www.njk.co.jp/",
	"https://www.njc.co.jp/",
	"https://www.nittotec.co.jp/",
	"https://www.nitto-cs.co.jp/",
	"https://www.nisseicom.co.jp/",
	"https://www.nissay-it.co.jp/",
	"https://www.nid.co.jp/",
	"https://www.ngs.teldevice.co.jp/",
	"https://www.net-a.co.jp/",
	"https://www.nesty.co.jp/",
	"https://www.nec-nexs.com/",
	"https://www.nds-tyo.co.jp/",
	"https://www.ndr.co.jp/",
	"https://www.ncos.co.jp/",
	"https://www.ncdsol.co.jp/",
	"https://www.nac-corp.jp/",
	"https://www.n-offs.co.jp/",
	"https://www.mytecno.com/",
	"https://www.muto.co.jp/",
	"https://www.msr.co.jp/",
	"https://www.module.co.jp/",
	"https://www.mn-sol.co.jp/",
	"https://www.mki.co.jp/",
	"https://www.mitsuiwa.co.jp/",
	"https://www.miraxia.com/",
	"https://www.mirai-sozo.co.jp/",
	"https://www.minori-sol.jp/",
	"https://www.midigitalservice.co.jp/",
	"https://www.metro.co.jp/",
	"https://www.meitetsucom.co.jp/",
	"https://www.mdi.co.jp/",
	"https://www.maxell-frontier.com/",
	"https://www.kyndryl.com/jp/ja",
	"https://www.kps.co.jp/",
	"https://www.kicnet.co.jp/",
	"https://www.jtp.co.jp/",
	"https://www.jss-net.com/",
	"https://www.jsol.co.jp/index.html",
	"https://www.jris.co.jp/",
	"https://www.jpd.co.jp/",
	"https://www.jmt.co.jp/",
	"https://www.jip.co.jp/",
	"https://www.jbs.co.jp/",
	"https://www.jbcchd.co.jp/",
	"https://www.jastec.co.jp/",
	"https://www.jasmine-sys.co.jp/",
	"https://www.iwatani-info.co.jp/",
	"https://www.is-info.co.jp/",
	"https://www.interactive-beauty.shiseido.com/",
	"https://www.intec.co.jp/",
	"https://www.infotec-s.co.jp/",
	"https://www.image-info.co.jp/",
	"https://www.ibm.com/jp-ja",
	"https://www.i-fourcom.co.jp/",
	"https://www.how.co.jp/",
	"https://www.hitec.co.jp/",
	"https://www.hitachi-sis.co.jp/",
	"https://www.hitachi-ics.co.jp/",
	"https://www.heartware.co.jp/",
	"https://www.hcs.co.jp/",
	"https://www.hcs-hd.co.jp/",
	"https://www.hcnet.co.jp/",
	"https://www.hal-eng.co.jp/",
	"https://www.gtl.co.jp/",
	"https://www.global.toshiba/jp/company/digitalsolution.html",
	"https://www.fusodentsu.co.jp/",
	"https://www.furukawa-ns.co.jp/",
	"https://www.fujitec-sol.co.jp/",
	"https://www.fsastech.com/ja-jp/",
	"https://www.frentec.co.jp/",
	"https://www.forbs.co.jp/",
	"https://www.fit-works.co.jp/",
	"https://www.fis-net.co.jp/",
	"https://www.ess.co.jp/",
	"https://www.esco.co.jp/",
	"https://www.e-wave.co.jp/",
	"https://www.e-sol.co.jp/",
	"https://www.doraku-holdings.co.jp/",
	"https://www.docomobs.com/",
	"https://www.dnp-digitalsolutions.co.jp/",
	"https://www.dg-net.co.jp/",
	"https://www.dentsusoken.com/",
	"https://www.densan-s.co.jp/",
	"https://www.daiwa-computer.co.jp/",
	"https://www.cubesystem.co.jp/",
	"https://www.create-inc.co.jp/",
	"https://www.computron.co.jp/",
	"https://www.cnw.co.jp/",
	"https://www.cns.co.jp/",
	"https://www.cch.co.jp/",
	"https://www.catena.co.jp/",
	"https://www.cap-know.com/",
	"https://www.caica.jp/",
	"https://www.cac-holdings.com/index.php",
	"https://www.busy-bee.co.jp/",
	"https://www.bsnnet.co.jp/",
	"https://www.bft-corp.com/",
	"https://www.beex-inc.com/",
	"https://www.basenet.co.jp/",
	"https://www.astic.co.jp/",
	"https://www.ascom.co.jp/",
	"https://www.argo-tec.co.jp/",
	"https://www.aps-s.co.jp/",
	"https://www.alphatec-sol.co.jp/",
	"https://www.alpha.co.jp/",
	"https://www.ale-systems.co.jp/",
	"https://www.aiskk.co.jp/",
	"https://www.aeonibs.co.jp/",
	"https://www.ado-business.com/",
	"https://www.actis.co.jp/",
	"https://www.acmos.co.jp/",
	"https://the-ings.co.jp/",
	"https://robust-lab.com/",
	"https://ncsa.jp/",
	"https://gophertec.com/",
	"https://giveandtake.co.jp/",
	"https://ebrain.co.jp/",
	"https://digitalmoon.co.jp/",
	"https://back-yard.co.jp/",
	"https://adniss.jp/",
}

func fetchContactURL(site string) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", site, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", uarand.GetRandom())

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	var contactURL string
	doc.Find("a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		href, exists := s.Attr("href")
		text := s.Text()
		if exists && strings.Contains(href, "contact") || strings.Contains(text, "お問い合わせ") {
			parsed, err := url.Parse(href)
			if err != nil {
				return true // skip
			}
			base, _ := url.Parse(site)
			contactURL = base.ResolveReference(parsed).String()
			return false // break
		}
		return true
	})

	if contactURL == "" {
		return "", fmt.Errorf("contact page not found")
	}

	return contactURL, nil
}

func main() {
	outputFile, err := os.Create("results.csv")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	writer.Write([]string{"Company URL", "Contact Page"})

	for _, company := range companyList {
		fmt.Println("🔍 Searching:", company)
		contact, err := fetchContactURL(company)
		if err != nil {
			fmt.Printf("❌ %s: %v\n", company, err)
			writer.Write([]string{company, ""})
		} else {
			fmt.Printf("✅ %s → %s\n", company, contact)
			writer.Write([]string{company, contact})
		}
	}
}
