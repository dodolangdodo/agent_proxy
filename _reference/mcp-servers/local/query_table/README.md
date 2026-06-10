# mcp_query_table

1. A financial web table scraper based on `playwright` that supports `Model Context Protocol (MCP)`. Currently supports the following sources:

    - [Tonghuashun iWencai](http://iwencai.com/)
    - [Tongdaxin Ask Xiaoda](https://wenda.tdx.com.cn/)
    - [East Money Stock Screener](https://xuangu.eastmoney.com/)

   In live trading, if a website goes down or undergoes redesign, you can immediately switch to other websites. (Note: Different websites have different table structures and require adaptation in advance)

2. A large language model calling scraper based on `playwright`. Currently supports the following sources:
    - [Nano Search](https://www.n.cn/)
    - [Tencent Yuanbao](https://yuanbao.tencent.com/)
    - [Baidu AI Search](https://chat.baidu.com/)

   `RooCode` provides `Human Reply` functionality. However, we found that the web version of `Nano Search` breaks formatting when copying, so we developed this feature.

## Installation

```commandline
pip install -i https://pypi.org/simple --upgrade mcp_query_table
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple --upgrade mcp_query_table
```

## Usage

```python
import asyncio

from mcp_query_table import *


async def main() -> None:
    async with BrowserManager(endpoint="http://127.0.0.1:9222", executable_path=None, devtools=True) as bm:
        # iWencai requires browser width > 768 to prevent mobile interface adaptation
        page = await bm.get_page()
        df = await query(page, 'Top 200 ETFs with best returns', query_type=QueryType.ETF, max_page=1, site=Site.THS)
        print(df.to_markdown())
        df = await query(page, 'Top 50 funds by year-to-date returns', query_type=QueryType.Fund, max_page=1, site=Site.TDX)
        print(df.to_csv())
        df = await query(page, 'Top 10 industry sectors by market cap', query_type=QueryType.Index, max_page=1, site=Site.TDX)
        print(df.to_csv())
        # TODO East Money pagination requires login in advance
        df = await query(page, 'Top 5 concept sectors by today\'s gains;', query_type=QueryType.Board, max_page=3, site=Site.EastMoney)
        print(df)

        output = await chat(page, "What does 1+2 equal?", provider=Provider.YuanBao)
        print(output)
        output = await chat(page, "What does 3+4 equal?", provider=Provider.YuanBao, create=True)
        print(output)

        print('done')
        bm.release_page(page)
        await page.wait_for_timeout(2000)


if __name__ == '__main__':
    asyncio.run(main())

```

## Important Notes

1. The browser should preferably be `Chrome`. If you must use `Edge`, besides closing all `Edge` windows, you also need to terminate all `Microsoft Edge` processes in Task Manager, i.e., `taskkill /f /im msedge.exe`
2. The browser should maintain sufficient window width to prevent some websites from automatically adapting to mobile version, which could cause table query failures
3. If you have website accounts, please log in in advance. This tool does not have automatic login functionality
4. Different websites have different table structures, and the number of stocks returned under the same conditions also differs. Adaptation is needed after querying

## How It Works

Unlike `requests`, `playwright` is browser-based and simulates user operations in the browser.

1. No need to solve login issues
2. No need to solve request construction and response parsing
3. Can directly obtain table data - what you see is what you get
4. Runs slower than `requests`, but has higher development efficiency

Data acquisition methods include:

1. Direct HTML table parsing
    1. Numbers are converted to text, not conducive to later research
    2. Strongest applicability
2. Intercepting requests to get returned `json` data
    1. Similar to `requests`, requires response parsing
    2. Less flexible, needs re-adaptation after website redesigns

This project uses simulated browser clicks to send requests and intercepts responses for data parsing.

Future adaptations will use more suitable methods based on different website redesign situations.

## Headless Mode

Headless mode runs faster, but some websites require login in advance, so headless mode must specify `user_data_dir`, otherwise login issues may occur.

- When `endpoint=None`, `headless=True` can start a new browser instance headlessly. Specify `executable_path` and `user_data_dir` to ensure normal operation in headless mode.
- When `endpoint` starts with `http://`, it connects to a headed browser started in `CDP` mode, with required parameter `--remote-debugging-port`. `executable_path` is the local browser path.
- When `endpoint` starts with `ws://`, it connects to a remote `Playwright Server`. This is also headless mode, but cannot specify `user_data_dir`, so usage is limited
  - Reference: <https://playwright.dev/python/docs/docker#running-the-playwright-server>

## MCP Support

Ensure you can execute `python -m mcp_query_table -h` in the console. If not, you may need to `pip install mcp_query_table` first.

In `Cline`, you can configure as follows. Where `command` is the absolute path to `python`, and `timeout` is the timeout in seconds. Since AI platforms often require over 1 minute for responses, a large timeout value needs to be set.

### STDIO Method

```json
{
  "mcpServers": {
    "mcp_query_table": {
      "timeout": 300,
      "command": "D:\\Users\\Kan\\miniconda3\\envs\\py312\\python.exe",
      "args": [
        "-m",
        "mcp_query_table",
        "--format",
        "markdown",
        "--endpoint",
        "http://127.0.0.1:9222",
        "--executable_path",
        "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
      ]
    }
  }
}
```

### SSE Method

First execute the following command in the console to start the `MCP` service:

```commandline
python -m mcp_query_table --format markdown --transport sse --port 8000 --endpoint http://127.0.0.1:9222
```

Then you can connect to the `MCP` service:

```json
{
  "mcpServers": {
    "mcp_query_table": {
      "timeout": 300,
      "url": "http://127.0.0.1:8000/sse"
    }
  }
}
```

## Debugging with `MCP Inspector`

```commandline
npx @modelcontextprotocol/inspector python -m mcp_query_table --format markdown --endpoint http://127.0.0.1:9222
```

Opening browsers and pagination is a time-consuming operation that may cause `MCP Inspector` page timeouts. You can use `http://localhost:5173/?timeout=300000` to set a timeout of 300 seconds.

This is my first attempt at writing an `MCP` project, so there may be various issues. Welcome everyone to communicate and exchange ideas.

## `MCP` Usage Tips

1. Top 100 stocks with highest gains in 2024 ranked by total market cap on December 31, 2024. Results differ across the three websites:
    - Tonghuashun: Shows 2201 stocks. Top 5 are ICBC, Agricultural Bank, China Mobile, PetroChina, CCB
    - Tongdaxin: Shows 100 stocks. Top 5 are Cambricon, Zhengdan, Huijin Tech, Wanfeng Auto, Airong Software
    - East Money: Shows 100 stocks. Top 5 are Hygon, Cambricon, Kuang-Chi, Runze Tech, Innolight

2. Large language models have weak question decomposition abilities, so questions should be asked reasonably to ensure query conditions aren't modified. Methods 2 and 3 below are recommended:
    - Top 100 stocks with highest gains in 2024 ranked by total market cap on December 31, 2024
      > LLMs may likely split this sentence, causing a single query to become multiple queries
    - Query East Money for "Top 100 stocks with highest gains in 2024 ranked by total market cap on December 31, 2024"
      > Use quotes to avoid splitting
    - Query East Money sectors for "worst performing industry sectors last year", then query the top 5 best performing stocks in that sector last year
      > Split into two queries: first query sectors, then stocks. But it's best not to be fully automatic, as it doesn't understand the difference between "today's gains" and "period gains" from the first step results, requiring interactive correction

## Streamlit Support

Implements querying financial data on the same page and manually inputting into AI for deep analysis. Refer to the `README.md` file in the `streamlit` directory.

## References

- [Selenium webdriver cannot attach to edge instance, edge's --remote-debugging-port option is ineffective](https://blog.csdn.net/qq_30576521/article/details/142370538)
- <https://github.com/AtuboDad/playwright_stealth/issues/31>
