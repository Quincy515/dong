#[derive(Debug)]
struct Res {
    area: String,
    title: String,
    url: String,
    time: String,
}
fn main() {
    let response = reqwest::blocking::get("http://fwpt.csggzy.cn/jyxxfjsz/index.jhtml")
        .unwrap()
        .text()
        .unwrap();
    let document = scraper::Html::parse_document(&response);
    let main_list_selector = scraper::Selector::parse("body > div.bg-main > div.container-div.jyxx > div > div.right-nr > div.main-list > ul > li").unwrap();
    let mut res: Vec<Res> = vec![];
    for element in document.select(&main_list_selector) {
        let title_selector = scraper::Selector::parse("p.list-leftp > a").unwrap();
        let title = element
            .select(&title_selector)
            .next()
            .unwrap()
            .value()
            .attr("title")
            .unwrap()
            .trim()
            .to_owned();
        let url = element
            .select(&title_selector)
            .next()
            .unwrap()
            .value()
            .attr("href")
            .unwrap()
            .trim()
            .to_owned();
        let area_selector = scraper::Selector::parse("p.list-leftp > a > span").unwrap();
        let area = element
            .select(&area_selector)
            .next()
            .unwrap()
            .inner_html()
            .trim()
            .to_owned();
        let time_selector = scraper::Selector::parse("p.list-rightp").unwrap();
        let time = element
            .select(&time_selector)
            .next()
            .unwrap()
            .inner_html()
            .trim()
            .to_owned();
        res.push(Res {
            area,
            title,
            url,
            time,
        })
    }
    println!("{:#?}", res);
}
