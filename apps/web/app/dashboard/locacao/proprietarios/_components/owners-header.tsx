import { HeaderCard, HeaderCardContent, HeaderCardHead } from "@/components/header-card"
import { Owner } from "@/types/owner"

interface OwnersHeaderProps {
  owners: Owner[]
  propertiesCount: number
}

export function OwnersHeader({ owners, propertiesCount }: OwnersHeaderProps) {
  const ownersCount = owners.length
  const propertiesPerOwner = ownersCount > 0 ? (propertiesCount / ownersCount) : propertiesCount

  return (
    <div className="grid gap-6 md:grid-cols-2">
      <HeaderCard>
        <HeaderCardHead
          title="Total de Proprietários"
          description="Proprietários cadastrados"
        />
        <HeaderCardContent
          count={ownersCount.toString()}
          className="text-muted"
        />
      </HeaderCard>

      <HeaderCard>
        <HeaderCardHead
          title="Imóveis por Proprietário"
          description="Média de imóveis (Aproximada)"
        />
        <HeaderCardContent
          count={propertiesPerOwner.toString()}
          className="text-muted"
        />
      </HeaderCard>
    </div>
  )
}