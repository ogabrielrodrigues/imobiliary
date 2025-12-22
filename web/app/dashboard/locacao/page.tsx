import { listProperties } from "@/actions/queries/property/list-properties"
import { Metadata } from "next"
import { RentalsHeader } from "./_components/rentals-header"

export const metadata: Metadata = {
  title: "Alugueres",
  description: "Visão geral dos imóveis alugados",
}

export default async function RentalsPage() {
  const { properties: foundProperties, status } = await listProperties()

  if (status !== 200) {
    return <div className="w-full flex justify-center">
      <p className="font-medium text-muted">Erro ao carregar os imóveis</p>
    </div>
  }

  const properties = !foundProperties ? [] : foundProperties

  return (
    <div className="container mx-auto flex flex-col space-y-10">
      <RentalsHeader properties={properties} />
    </div>
  )
}
