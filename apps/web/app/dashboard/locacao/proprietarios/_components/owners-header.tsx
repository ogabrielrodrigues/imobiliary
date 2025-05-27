import { HeaderCard, HeaderCardContent, HeaderCardHead } from "@/components/HeaderCard"
import { Owner } from "@/types/owner"

interface OwnersHeaderProps {
  owners: Owner[]
  propertiesCount: number
}

export function OwnersHeader({ owners, propertiesCount }: OwnersHeaderProps) {
  const ownersCount = owners.length
  const propertiesPerOwner = owners.length > 0 ? (propertiesCount / owners.length) : propertiesCount
  const propertiesPerOwnerLabel = propertiesCount === 0 ? "Nenhum imóvel cadastrado" : `${propertiesCount} imóveis para ${ownersCount} proprietários`

  return (
    <div className="grid gap-6 md:grid-cols-2">
      <HeaderCard className="from-purple-500 to-purple-700">
        <HeaderCardHead
          title="Total de Proprietários"
          description="Proprietários cadastrados"
        />
        <HeaderCardContent
          count={ownersCount.toString()}
          label="Total de proprietários ativos"
          className="text-purple-100"
        />
      </HeaderCard>

      <HeaderCard className="from-red-500 to-red-700">
        <HeaderCardHead
          title="Imóveis por Proprietário"
          description="Média de imóveis (Aproximada)"
        />
        <HeaderCardContent
          count={propertiesPerOwner.toString()}
          label={propertiesPerOwnerLabel}
          className="text-red-100"
        />
      </HeaderCard>
    </div>
  )
}