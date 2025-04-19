import { getProperties } from "@/actions/properties"
import { Metadata } from "next"
import { PropertiesHeader } from "./_components/properties-header"
import { PropertiesTable } from "./_components/properties-table"

export const metadata: Metadata = {
  title: "Imóveis",
  description: "Gerencie os imóveis disponíveis para aluguel",
}

export default async function PropertiesPage() {
  const found = await getProperties()

  const properties = found.length > 0 ? found : []

  return (
    <div className="container mx-auto flex flex-col space-y-4">
      <PropertiesHeader properties={properties} />

      <PropertiesTable properties={properties} />
    </div >
  )
}
