# Crypto Trends


### Disclaimer
> This project aims to explore AI-driven crypto analysis but does not offer financial advice. Remember:
> * Cryptocurrencies are highly volatile, and past performance does not guarantee future results.
> * Always conduct your own research and due diligence before investing.
> * The developers of this tool are not liable for any losses incurred from its use.


## Tools
- [Coinmarketcap](https://coinmarketcap.com)
- [Google AI Studio](https://aistudio.google.com)

## Endpoints

| Endpoint          | Description                                                                                               |
|-------------------|-----------------------------------------------------------------------------------------------------------|
| `/api/historical` | Returns historical prices of cryptocurrencies. Useful for price trend analysis over time.                |
| `/api/crypto`     | Returns information about cryptocurrencies. This may include data such as name, symbol, market cap, current price, trading volume, etc. |
| `/api/map`        | Lists available cryptocurrencies. Can be used to get an overview of cryptocurrencies supported by the platform. |

## Install

1. Clone the project repository:
   ```bash
   git clone https://github.com/your_username/crypto-trends.git
   ```
2. Navigate to the project directory:
   ```bash
   cd crypto-trends
   ```
3. Install project dependencies:
    ```bash
    go mod tidy
    ```
4. Create a .env file in the root directory of the project and add your API keys like the .env.example:
    ```plaintext
    COINMARKETCAP_API_KEY=YOUR_MARKETCAP_API_KEY
    GEMINI_API_KEY=YOUR_GEMINI_API_KEY
    ```
## Execution

To run the project, execute the command:

```bash
go run main.go
```


<!-- ## Features
- Chatbot to integrate on Telegram/Discord.
-  -->
