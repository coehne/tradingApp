import { NewTsxForm } from "../components/molecules/NewTsxForm"

export type TransactionType = "withdraw" | "deposit"

const NewTransaction: React.FC<{ type: TransactionType }> = ({ type }) => (
  <div>
    <NewTsxForm type={type} />
  </div>
)

export default NewTransaction
