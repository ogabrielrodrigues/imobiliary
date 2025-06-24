import { HeaderCard, HeaderCardContent, HeaderCardHead } from "@/components/header-card"

type TenantsHeaderProps = {}

export function TenantsHeader({ }: TenantsHeaderProps) {
  return (
    <div className="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
      <HeaderCard>
        <HeaderCardHead
          title="Total de Inquilinos"
          description="Inquilinos ativos no sistema"
        />
        <HeaderCardContent
          count="78"
          className="text-muted"
        />
      </HeaderCard>
      <HeaderCard>
        <HeaderCardHead
          title="Novos Inquilinos"
          description="Adicionados este mês"
        />
        <HeaderCardContent
          count="9"
          className="text-muted"
        />
      </HeaderCard>
      <HeaderCard className="md:col-span-2 xl:col-span-1">
        <HeaderCardHead
          title="Inquilinos Pontuais"
          description="Pagamentos em dia"
        />
        <HeaderCardContent
          count="92%"
          className="text-muted"
        />
      </HeaderCard>
    </div>
  )
}