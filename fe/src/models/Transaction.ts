import { Trade } from "./Trade";

export interface Transaction {
    id: number,
    createdAt: string,
    amount: number,
    trade: Trade | null

}
