import { Trade } from "../models/Trade"
import { numberToUSD } from "../utils/formatting"

export const TradeCard: React.FC<{ trade: Trade | undefined }> = ({
  trade,
}) => (
  <div>
    <div className="max-w-md w-full-md mx-auto border p-8 border-gray-300 mt-4 text-md  bg-gray-700 text-gray-200 rounded-md">
      <p>Company Name: {trade?.companyName}</p>
      <p>Symbol: {trade?.symbol}</p>
      <p>Quantity: {trade?.qty}</p>
      <p>Stock Price: {numberToUSD(trade?.price ? trade?.price : 0)}</p>
      <p>
        Total Price: {numberToUSD(trade?.price ? trade?.price * trade.qty : 0)}
      </p>
    </div>
  </div>
)
