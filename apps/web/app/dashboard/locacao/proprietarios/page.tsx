import { listOwners } from "@/actions/queries/owner/list-owners"
import { listProperties } from "@/actions/queries/property/list-properties"
import { Metadata } from "next"
import { OwnersHeader } from "./_components/owners-header"
import { OwnersSection } from "./_components/owners-section"

export const metadata: Metadata = {
  title: "Proprietários",
  description: "Gerencie os proprietários dos imóveis",
}

export default async function OwnersPage() {
  const { status: owner_status, owners: found } = await listOwners()

  if (owner_status != 200) {
    return (
      <div className="w-full flex justify-center">
        <p className="font-medium text-muted">Erro ao carregar proprietários</p>
      </div>
    )
  }

  const owners = !found ? [] : found

  const { status: properties_status, properties } = await listProperties()

  const propertiesCount = Math.floor(properties_status !== 200 ? 0 : properties!.length)

  return (
    <div className=" flex flex-col space-y-10">
      <OwnersHeader owners={owners} propertiesCount={propertiesCount} />

      <OwnersSection owners={owners} />
    </div>
  )
}
