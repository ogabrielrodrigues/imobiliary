import { Metadata } from "next"
import { RentalsHeader } from "./_components/rentals-header"

export const metadata: Metadata = {
  title: "Alugueres",
  description: "Visão geral dos imóveis alugados",
}

export default function RentalsPage() {
  return (
    <div className="container mx-auto flex flex-col space-y-10">
      <RentalsHeader />
    </div>
  )
}
