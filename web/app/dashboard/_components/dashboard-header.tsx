import { HeaderCard, HeaderCardContent, HeaderCardHead } from "@/components/HeaderCard"

type DashboardHeaderProps = {}

export function DashboardHeader({ }: DashboardHeaderProps) {
  return (
    <div className="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
      <HeaderCard className="from-cyan-500 to-cyan-700">
        <HeaderCardHead
          title="Imóveis Ativos"
          description="Total de imóveis em locação"
        />
        <HeaderCardContent
          count="253"
          label="+8.3% desde o mês passado"
          className="text-cyan-100"
        />
      </HeaderCard>
      <HeaderCard className="from-purple-500 to-purple-700">
        <HeaderCardHead
          title="Contratos Ativos"
          description="Total de contratos em vigor"
        />
        <HeaderCardContent
          count="87"
          label="+5.2% desde o mês passado"
          className="text-purple-100"
        />
      </HeaderCard>
      <HeaderCard className="from-amber-500 to-amber-700" containerClassName="md:col-span-2 xl:col-span-1">
        <HeaderCardHead
          title="Taxa de Ocupação"
          description="Percentual de imóveis ocupados"
        />
        <HeaderCardContent
          count="85.5%"
          label="+2.3% desde o mês passado"
          className="text-amber-100"
        />
      </HeaderCard>
    </div>
  )
}