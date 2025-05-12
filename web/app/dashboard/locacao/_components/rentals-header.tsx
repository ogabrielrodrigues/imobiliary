import { HeaderCard, HeaderCardContent, HeaderCardHead } from "@/components/HeaderCard"

type RentalsHeaderProps = {}

export function RentalsHeader({ }: RentalsHeaderProps) {
  return (
    <div className="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
      <HeaderCard className="from-blue-500 to-blue-700">
        <HeaderCardHead
          title="Contratos Ativos"
          description="Total de contratos em vigor"
        />
        <HeaderCardContent
          count="42"
          label="Aumento de 8% em relação ao mês anterior"
          className="text-blue-100"
        />
      </HeaderCard>
      <HeaderCard className="from-green-500 to-green-700">
        <HeaderCardHead
          title="Receita Mensal"
          description="Valor total de aluguéis"
        />
        <HeaderCardContent
          count="R$ 68.450"
          label="Aumento de 5% em relação ao mês anterior"
          className="text-green-100"
        />
      </HeaderCard>
      <HeaderCard className="from-rose-500 to-rose-700" containerClassName="md:col-span-2 xl:col-span-1">
        <HeaderCardHead
          title="Imóveis Disponíveis"
          description="Prontos para locação"
        />
        <HeaderCardContent
          count="15"
          label="3 novos imóveis este mês"
          className="text-rose-100"
        />
      </HeaderCard>
    </div>
  )
}