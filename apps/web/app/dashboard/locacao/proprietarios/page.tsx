import { listOwners } from "@/actions/queries/owner/list-owners"
import { listProperties } from "@/actions/queries/property/list-properties"
import { Separator } from "@/components/ui/separator"
import { Metadata } from "next"
import { OwnersHeader } from "./_components/owners-header"
import { OwnersSection } from "./_components/owners-section"

export const metadata: Metadata = {
  title: "Proprietários",
  description: "Gerencie os proprietários dos imóveis",
}

export default async function OwnersPage() {
  const { status: owner_status, owners: found } = await listOwners()

  if (!found || owner_status != 200) {
    return (
      <div className="container mx-auto flex flex-col space-y-10">
        <h1 className="text-2xl font-bold">Erro ao carregar proprietários</h1>
      </div>
    )
  }

  const { status: properties_status, properties } = await listProperties()

  const owners = found.length > 0 ? found : []
  const propertiesCount = properties_status !== 200 ? 0 : properties!.length

  return (
    <div className="container mx-auto flex flex-col space-y-10">
      <OwnersHeader owners={found} propertiesCount={propertiesCount} />

      <Separator />

      <OwnersSection owners={owners} />
    </div>
  )
}
