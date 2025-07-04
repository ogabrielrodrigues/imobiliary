import { HeaderCard, HeaderCardContent, HeaderCardHead } from "@/components/header-card"
import { Property } from "@/types/property"

interface PropertiesHeaderProps {
  properties: Property[]
}

export function PropertiesHeader({ properties }: PropertiesHeaderProps) {
  const occupied = properties.filter(property => property.status === 'Ocupado')
  const occupied_percent = ((occupied.length / properties.length) * 100) || 0

  const available = properties.filter(property => property.status === 'Disponível')
  const available_percent = ((available.length / properties.length) * 100) || 0

  return (
    <div className="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
      <HeaderCard>
        <HeaderCardHead
          title="Imóveis Cadastrados"
          description="Total de imóveis no sistema"
        />
        <HeaderCardContent
          count={properties.length.toString()}
          className="text-muted"
        />
      </HeaderCard>
      <HeaderCard>
        <HeaderCardHead
          title="Imóveis Alugados"
          description="Atualmente ocupados"
        />
        <HeaderCardContent
          count={occupied.length.toString()}
          className="text-muted"
        />
      </HeaderCard>
      <HeaderCard className="md:col-span-2 xl:col-span-1">
        <HeaderCardHead
          title="Imóveis Disponíveis"
          description="Prontos para locação"
        />
        <HeaderCardContent
          count={available.length.toString()}
          className="text-muted"
        />
      </HeaderCard>
    </div>
  )
}