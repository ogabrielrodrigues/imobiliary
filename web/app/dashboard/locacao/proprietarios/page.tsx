import { getOwners } from "@/actions/owner"
import { getProperties } from "@/actions/property"
import { Separator } from "@/components/ui/separator"
import { Metadata } from "next"
import { OwnersHeader } from "./_components/owners-header"
import { OwnersSection } from "./_components/owners-section"

export const metadata: Metadata = {
  title: "Proprietários",
  description: "Gerencie os proprietários dos imóveis",
}

export default async function OwnersPage() {
  const { status: owners_status, owners } = await getOwners()
  const { status: properties_status, properties } = await getProperties()

  if (owners_status !== 200 || properties_status !== 200) {
    return (
      <div className="container mx-auto flex flex-col space-y-10">
        <h1 className="text-2xl font-bold">Erro ao carregar proprietários</h1>
      </div>
    )
  }

  return (
    <div className="container mx-auto flex flex-col space-y-10">
      <OwnersHeader owners={owners} propertiesCount={properties.length} />

      <Separator />

      <OwnersSection owners={owners} />
    </div>
  )
}
