import { getProperties } from "@/actions/properties"
import { Separator } from "@/components/ui/separator"
import { Metadata } from "next"
import { PropertiesHeader } from "./_components/properties-header"
import { PropertiesSection } from "./_components/properties-section"


export const metadata: Metadata = {
  title: "Imóveis",
  description: "Gerencie os imóveis disponíveis para aluguel",
}

export default async function PropertiesPage() {
  const found = await getProperties()

  const properties = found.length > 0 ? found : []

  return (
    <div className="container mx-auto flex flex-col space-y-10">
      <PropertiesHeader properties={properties} />

      <Separator />

      <PropertiesSection properties={properties} />
    </div>
  )
}
