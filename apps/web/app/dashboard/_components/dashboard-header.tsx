import { HeaderCard, HeaderCardContent, HeaderCardHead } from "@/components/header-card"

type DashboardHeaderProps = {}

export function DashboardHeader({ }: DashboardHeaderProps) {
  return (
    <div className="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
      <HeaderCard>
        <HeaderCardHead
          title="Imóveis Ativos"
          description="Total de imóveis em locação"
        />
        <HeaderCardContent
          count="253"
          className="text-muted"
        />
      </HeaderCard>
      <HeaderCard>
        <HeaderCardHead
          title="Contratos Ativos"
          description="Total de contratos em vigor"
        />
        <HeaderCardContent
          count="87"
          className="text-muted"
        />
      </HeaderCard>
      <HeaderCard className="md:col-span-2 xl:col-span-1">
        <HeaderCardHead
          title="Taxa de Ocupação"
          description="Percentual de imóveis ocupados"
        />
        <HeaderCardContent
          count="85.5%"
          className="text-muted"
        />
      </HeaderCard>
    </div>
  )
}