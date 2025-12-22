import { HeaderCard, HeaderCardContent, HeaderCardHead } from "@/components/header-card"
import { Owner } from "@/types/owner"
import { Property } from "@/types/property"

interface OwnersHeaderProps {
  owners: Owner[]
  properties: Property[]
}

export function OwnersHeader({ owners, properties }: OwnersHeaderProps) {
  const ownersCount = owners.length
  const propertiesCount = properties.length
  const propertiesPerOwner = ownersCount > 0 ? (propertiesCount / ownersCount) : propertiesCount

  return (
    <div className="grid gap-6 md:grid-cols-2">
      <HeaderCard>
        <HeaderCardHead
          title="Total de Proprietários"
          description="Proprietários cadastrados"
        />
        <HeaderCardContent
          count={ownersCount.toString().padStart(2, '0')}
          className="text-muted"
        />
      </HeaderCard>

      <HeaderCard>
        <HeaderCardHead
          title="Imóveis por Proprietário"
          description="Média de imóveis (Aproximada)"
        />
        <HeaderCardContent
          count={propertiesPerOwner.toString().padStart(2, '0')}
          className="text-muted"
        />
      </HeaderCard>
    </div>
  )
}