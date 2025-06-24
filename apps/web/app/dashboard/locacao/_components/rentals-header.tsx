import { HeaderCard, HeaderCardContent, HeaderCardHead } from "@/components/header-card"

type RentalsHeaderProps = {}

export function RentalsHeader({ }: RentalsHeaderProps) {
  return (
    <div className="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
      <HeaderCard>
        <HeaderCardHead
          title="Contratos Ativos"
          description="Total de contratos em vigor"
        />
        <HeaderCardContent
          count="42"
          className="text-muted"
        />
      </HeaderCard>
      <HeaderCard >
        <HeaderCardHead
          title="Receita Mensal"
          description="Valor total de aluguéis"
        />
        <HeaderCardContent
          count="R$ 68.450"
          className="text-muted"
        />
      </HeaderCard>
      <HeaderCard className="md:col-span-2 xl:col-span-1">
        <HeaderCardHead
          title="Imóveis Disponíveis"
          description="Prontos para locação"
        />
        <HeaderCardContent
          count="15"
          className="text-muted"
        />
      </HeaderCard>
    </div>
  )
}