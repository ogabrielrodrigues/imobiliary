import { listProperties } from "@/actions/queries/property/list-properties"
import { Metadata } from "next"
import { PropertiesHeader } from "./_components/properties-header"
import { PropertiesSection } from "./_components/properties-section"

export const metadata: Metadata = {
  title: "Imóveis",
  description: "Gerencie os imóveis disponíveis para aluguel",
}

export default async function PropertiesPage() {
  const { properties: foundProperties, status } = await listProperties()

  if (status !== 200) {
    return <div className="w-full flex justify-center">
      <p className="font-medium text-muted">Erro ao carregar os imóveis</p>
    </div>
  }

  const properties = !foundProperties ? [] : foundProperties

  return (
    <div className="container mx-auto flex flex-col space-y-10">
      <PropertiesHeader properties={properties} />

      <PropertiesSection properties={properties} />
    </div>
  )
}
