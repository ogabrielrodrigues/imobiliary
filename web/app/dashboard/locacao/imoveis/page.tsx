import { getProperties } from "@/actions/properties"
import { Metadata } from "next"
import { PropertiesHeader } from "./_components/properties-header"
import { PropertiesTable } from "./_components/properties-table"

export const metadata: Metadata = {
  title: "Imóveis",
  description: "Gerencie os imóveis disponíveis para aluguel",
}

export default async function PropertiesPage() {
  const property = await getProperties()

  const properties = Array.isArray(property) ? property : [property]

  return (
    <div className="container mx-auto">
      <PropertiesHeader />

      <PropertiesTable properties={properties} />
    </div >
  )
}
