import Image from "next/image";
export const Coins = ({ }) => {
  return (
      <div className={"coin_page"}>
        <div className={"coin_container"}>
          <Image
            src={"teste.png"}
            alt={"coin.name"}
            width={50}
            height={50}
            className="coin_image"
          />
          <h1 className="coin_name">{"coin.name"}</h1>
          <p className={"styles.coin_ticker"}>{"coin.symbol"}</p>
          <p className={"coin_current"}>
            {"coin.market_data.current_price.usd"}
          </p>
        </div>
      </div>
  );
};
